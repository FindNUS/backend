package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/* ---- MILESTONE 1 DEMO ----*/
// This script is to demonstrate mongoDB connection capabilities
// Temporary function for milestone 1

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

func setupMongo(dbName string) {
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

// DEMO HANDLER FOR ITEMS
type DemoForm struct {
	Name string `json:"name" binding:"required"`
}

func debugGetDemoItem(c *gin.Context) {
	coll := mongoDb.Collection(string(COLL_DEBUG))
	name := c.Query("name")
	res, err := coll.Find(
		context.TODO(),
		bson.D{{"Name", name}},
	)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Server Side Error",
		})
	}
	var items []Item
	for res.Next(context.TODO()) {
		var item Item
		res.Decode(&item)
		items = append(items, item)
	}
	if len(items) == 0 {
		c.JSON(404, gin.H{
			"message": "Nothing found!",
		})
		return
	}
	println("GET OK")
	c.JSON(200, items)
}

// Demo parser
func debugDemoPostItems() {
	f, err := os.Open("demo.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	var DemoItems []map[string]interface{}
	err = json.Unmarshal(bytes, &DemoItems)
	if err != nil {
		log.Fatal(err.Error())
	}
	PrettyPrintStruct(DemoItems[0])
	var newitems []NewItem
	for _, item := range DemoItems {
		var tmp NewItem
		var ok bool
		tmp.Name, ok = item["name"].(string)
		tmp.Date = ParseDateString(item["date"].(string))
		if !ok {
			log.Fatal("date unmarshall fail")
		}
		tmp.Location = item["location"].(string)
		tmp.Category = GetCategoryType(item["category"].(string))
		tmp.Image_url = item["image_url"].(string)
		newitems = append(newitems, tmp)
	}
	//post
	for _, item := range newitems {
		debugPostItem(COLL_DEBUG, item)
	}
}
