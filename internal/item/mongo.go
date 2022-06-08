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

// Recieves work to be done
func HandleRequest(rawMsg []byte) {
	msg := UnmarshallMessage(rawMsg)
	switch msg.Operation_type {
	case OPERATION_NEW_ITEM:
		DoAddItem(msg)
		break
	case OPERATION_PATCH_ITEM:
		DoUpdateItem(msg)
		break
	case OPERATION_GET_ITEM:
		//foo
		fmt.Println("Get Item Triggered")
		break
	case OPERATION_GET_ITEM_LIST:
		//foo
		fmt.Println("Get Item List Triggered")
		break
	case OPERATION_DEL_ITEM:
		DoDeleteItem(msg)
		break
	default:
		// foo
		break
	}
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

// Handle creation of new item
func DoAddItem(msg ItemMsgJSON) interface{} {
	// Unmarshall body
	var item NewItem
	var res interface{}
	json.Unmarshal(msg.Body, &item)
	if item.User_id == "" {
		// Assert that user_id only exists for found items
		res = MongoAddItem(COLL_FOUND, item)
	} else {
		res = MongoAddItem(COLL_LOST, item)
	}
	return res
}

func DoUpdateItem(msg ItemMsgJSON) int64 {
	var item PatchItem
	json.Unmarshal(msg.Body, &item)
	var id string
	var err error
	// Safety check, should not trigger
	if _, ok := msg.Params["Id"]; !ok {
		log.Println("Update failed, item does not exist")
		return -1
	}
	id = msg.Params["Id"][0]
	item.Id, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("ERROR WHILE PATCHING:", err.Error())
		return -1
	}
	id = msg.Params["Id"][0]
	if _, ok := msg.Params["User_id"]; ok {
		item.User_id = msg.Params["User_id"][0]
	}
	// Check which collection the request belongs to
	if item.User_id == "" {
		// Item belongs to FOUND collection
		return MongoPatchItem(COLL_FOUND, item)
	} else {
		// User_id presence implies the msg belongs to LOST collection
		return MongoPatchItem(COLL_LOST, item)
	}
}

func DoDeleteItem(msg ItemMsgJSON) int64 {
	// Assert that msg contains enough parameters
	var item DeletedItem
	var id string
	var err error
	// Safety check, should not trigger
	if _, ok := msg.Params["Id"]; !ok {
		log.Println("Delete failed, item does not exist")
		return -1
	}
	id = msg.Params["Id"][0]
	item.Id, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("ERROR WHILE DELETING:", err.Error())
		return -1
	}
	id = msg.Params["Id"][0]
	if _, ok := msg.Params["User_id"]; ok {
		item.User_id = msg.Params["User_id"][0]
	}
	// Check which collection the deleted item belongs to
	if item.User_id == "" {
		return MongoDeleteItem(COLL_FOUND, item)
	} else {
		return MongoDeleteItem(COLL_LOST, item)
	}
}
