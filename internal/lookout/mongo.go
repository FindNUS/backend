package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

// Processes MongoDB item and remaps certain fields
func ParseGetItemBody(tmp map[string]interface{}) Item {
	if tmp == nil {
		return Item{}
	}
	var item Item
	// Transform category
	tmp["Category"] = GetCategoryString(tmp["Category"].(int32))
	// Transform contact method
	if val, ok := tmp["Contact_method"]; ok {
		tmp["Contact_method"] = GetContactString(val.(int32))
	}
	tmp["Id"] = tmp["_id"]
	b, err := json.Marshal(tmp)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(b, &item)
	return item
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

// Lookup Lost collection and get all items that have set Lookout=true
// Pass COLL_LOST for live usage, COLL_DEBUG for debug usage
func MongoGetAllLookoutRequests(coll_name ItemCollections) []Item {
	coll := mongoDb.Collection(string(coll_name))
	// Parse pagination variables
	// Parse category filters
	// { $or : [ { "Category": {"$eq", "foo"}, {...} } ]
	filter := bson.M{}
	filter["Lookout"] = true

	log.Println("Searching MongoDB with filter:", filter)
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
