package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Utility function
func PrettyPrintStruct(any interface{}) {
	log.Printf("%+v\n", any)
}

// ---- GLOBAL DB VARIABLES ----
// Global variables are generally discouraged
// But I will use them here to greatly simplify things
// NOTE: The following are thread safe, so concurrency is possible.
var mongoClient *mongo.Client
var mongoDb *mongo.Database

type ItemCollections string

const (
	COLL_LOST  ItemCollections = "Lost"
	COLL_FOUND                 = "Found"
	COLL_DEBUG                 = "Debug"
	COLL_IMGUR                 = "Imgur"
)

func debugPostItem(collName ItemCollections, item NewItem) {
	coll := mongoDb.Collection(string(collName))
	_, err := coll.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// println(res.InsertedID)
	println("POST OK")
}

func SetupMongo() {
	mongoURI := os.Getenv("MONGO_URI")
	var err error
	if mongoURI == "" {
		// Read from secrets file
		f, err := os.Open("../../secrets/mongoDev.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(f)
		defer f.Close()
		for scanner.Scan() {
			mongoURI = scanner.Text()
		}
	}

	// README: https://www.mongodb.com/docs/drivers/go/current/fundamentals/connection/
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf(err.Error())
	}

	mongoDb = mongoClient.Database("Items")
	fmt.Println("MONGO SETUP OK")
}

func MongoAddItem(collName ItemCollections, item NewItem) interface{} {
	coll := mongoDb.Collection(string(collName))
	res, err := coll.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res.InsertedID
}

func MongoPatchItem(collname ItemCollections, item PatchItem) int64 {
	coll := mongoDb.Collection(string(collname))
	update := bson.M{"$set": item}
	res, err := coll.UpdateByID(context.TODO(), item.Id, update)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res.ModifiedCount
}

func MongoDeleteItem(collname ItemCollections, item DeletedItem) int64 {
	coll := mongoDb.Collection(string(collname))
	res, err := coll.DeleteOne(context.TODO(), item)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res.DeletedCount
}

// Get one specific item based on its id
func MongoGetItem(collname ItemCollections, id string, userid string) Item {
	coll := mongoDb.Collection(string(collname))
	myid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error getting ObjectID from hex")
		log.Println("Hex:", id)
	}
	query := bson.D{{"_id", myid}}
	if len(userid) > 0 {
		query = append(query, bson.E{"User_id", userid})
	}
	res, err := coll.Find(
		context.TODO(),
		query,
	)
	var generalItem map[string]interface{}
	// This should only run once
	for res.Next(context.TODO()) {
		res.Decode(&generalItem)
	}
	if generalItem == nil {
		return Item{}
	}
	return ParseGetItemBody(generalItem)
}

// Query for a paginated list of truncated items based on a set of filters
func MongoGetManyItems(collname ItemCollections, args map[string][]string) []Item {
	coll := mongoDb.Collection(string(collname))
	limit := int64(0)
	offset := int64(0)
	var ok bool
	var err error
	// Parse pagination variables
	if _, ok = args["limit"]; ok {
		limit, err = strconv.ParseInt(args["limit"][0], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if _, ok = args["offset"]; ok {
		offset, err = strconv.ParseInt(args["offset"][0], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	// Parse category filters
	// { $or : [ { "Category": {"$eq", "foo"}, {...} } ]
	filter := bson.M{}
	if catFilter, ok := args["category"]; ok {
		tmp := []bson.M{}
		for _, catStr := range catFilter {
			catInt := GetCategoryType(catStr)
			if err != nil {
				log.Println("Error parsing category filter")
			}
			tmp = append(tmp, bson.M{"Category": catInt})
		}
		filter = bson.M{"$or": tmp}
	}
	// Parse User_id filter, if exist
	if tmp, ok := args["User_id"]; ok {
		filter["User_id"] = tmp[0]
	}

	// Parse Date filters, if they exist
	startDateArr, startExist := args["startdate"]
	endDateArr, endExist := args["enddate"]
	if startExist || endExist {
		tmp := bson.M{}
		if startExist {
			startTime := primitive.NewDateTimeFromTime(ParseDateString(startDateArr[0]))
			tmp["$gte"] = startTime
		}
		if endExist {
			endTime := primitive.NewDateTimeFromTime(ParseDateString(endDateArr[0]))
			tmp["$lte"] = endTime
		}
		filter["Date"] = tmp
	}

	// log.Println("Searching MongoDB with filter:", filter)
	opts := options.Find()
	// Specify what fields to return. Id is implicitly returned
	opts.SetProjection(
		bson.D{
			{"Name", 1},
			{"Date", 1},
			{"Location", 1},
			{"Category", 1},
			{"User_id", 1},
			{"Image_url", 1},
		},
	)
	opts.SetSort(bson.D{{"Date", -1}, {"_id", 1}})
	opts.SetSkip(offset)
	opts.SetLimit(limit)

	res, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatal(err.Error())
	}
	items := []Item{}
	for res.Next(context.TODO()) {
		var generalItem map[string]interface{}
		var item Item
		res.Decode(&generalItem)
		if generalItem == nil {
			continue
		}
		item = ParseGetItemBody(generalItem)
		items = append(items, item)
	}
	return items
}

// Lookup Lost collection and get all items that have set Lookout=true
// Pass COLL_LOST for live usage, COLL_DEBUG for debug usage
func MongoGetAllLookoutRequests(coll_name ItemCollections) []Item {
	coll := mongoDb.Collection(string(coll_name))
	// Parse pagination variables
	// Parse category filters
	// { $or : [ { "Category": {"$eq", "foo"}, {...} } ]
	filter := bson.M{}
	filter["Lookout"] = true

	opts := options.Find()
	// Specify what fields to return. Id is implicitly returned
	opts.SetProjection(
		bson.D{
			{"Name", 1},
			{"Date", 1},
			{"Location", 1},
			{"Category", 1},
			{"User_id", 1},
			{"Item_details", 1},
		},
	)
	opts.SetSort(bson.D{{"Date", -1}})

	// TODO: Consider parsing all other filters, if they exist
	res, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatal(err.Error())
	}
	items := []Item{}
	for res.Next(context.TODO()) {
		var generalItem map[string]interface{}
		var item Item
		res.Decode(&generalItem)
		if generalItem == nil {
			continue
		}
		item = ParseGetItemBody(generalItem)
		items = append(items, item)
	}
	return items
}
