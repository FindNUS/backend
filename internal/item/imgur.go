package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// If an image changes, delete and add again.

// Handle CRUD of imgur images
var imgurClientID string = ""

// Interface of Imgur Collection item schema
type ImgurRef struct {
	ImageLink    string `bson:"ImageLink"`
	ImageDelHash string `bson:"ImageDelHash"`
}

// Setup clientID details
func SetupImgur() {
	imgurClientID = os.Getenv("IMGUR_CLIENT_ID")
	if imgurClientID == "" {
		// Read from secrets file
		f, err := os.Open("../../secrets/imgurId.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(f)
		defer f.Close()
		for scanner.Scan() {
			imgurClientID = scanner.Text()
		}
	}
	if imgurClientID == "" {
		log.Fatalf("imgur client id still empty")
	}
	log.Println("Imgur setup OK, id: " + imgurClientID)
}

// Store the deleteHash of the image link
// Returns the Mongo _id of the document entry
func MongoStoreImgurRef(link, delhash string) interface{} {
	coll := mongoDb.Collection(string(COLL_IMGUR))
	imgurRef := ImgurRef{
		ImageLink:    link,
		ImageDelHash: delhash,
	}
	res, err := coll.InsertOne(context.TODO(), imgurRef)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res.InsertedID
}

func MongoGetImgurRef(id string) ImgurRef {
	coll := mongoDb.Collection(string(COLL_IMGUR))
	objId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{"_id", objId}}
	res, err := coll.Find(
		context.TODO(),
		query,
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// var generalItem map[string]interface{}
	imgurRef := ImgurRef{}
	for res.Next(context.TODO()) {
		// This should only run once
		// If there are more than one case, the lastmost item will be returned
		res.Decode(&imgurRef)
	}
	return imgurRef
}

func MongoDeleteImgurRef(link string) int64 {
	coll := mongoDb.Collection(string(COLL_IMGUR))
	filter := bson.D{{"ImageLink", link}}
	res, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res.DeletedCount
}

func ImgurDeleteImageRef(delHash string) bool {
	url := "https://api.imgur.com/3/image" + "/" + delHash

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.Close()

	// http.NewRequest("POST", url, )
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, payload)
	if err != nil {
		panic(err.Error())
	}
	clientId := "Client-ID " + imgurClientID
	log.Println("-" + clientId + "-")
	req.Header.Add("Authorization", "Client-ID "+imgurClientID)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Execute and validate the HTTP request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	return res.StatusCode == 200
}

// Takes in a base64 value and uploads it to Imgur.
// Returns the publically available url and its delete hash
func ImgurAddNewImage(base64str string) (string, string) {
	imageLink := ""
	deleteHash := ""

	url := "https://api.imgur.com/3/image"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.WriteField("image", base64str)
	if err != nil {
		panic(err.Error())
	}

	// http.NewRequest("POST", url, )
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		panic(err.Error())
	}
	clientId := "Client-ID " + imgurClientID
	log.Println("-" + clientId + "-")
	req.Header.Add("Authorization", "Client-ID "+imgurClientID)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Execute and validate the HTTP request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Unmarshal Imgur's response body for processing
	b, err := ioutil.ReadAll(res.Body)
	var tmp map[string]interface{}
	json.Unmarshal(b, &tmp)

	// Sieve out imagelink and deletehash from the response body
	PrettyPrintStruct(tmp)
	if data, ok := tmp["data"]; ok {
		payload, exist := data.(map[string]interface{})
		if exist {
			imageLink, _ = payload["link"].(string)
			deleteHash, _ = payload["deletehash"].(string)
		}
	}

	return imageLink, deleteHash
}

func ImgurDeleteImageFromId(mongoId string) {
	ref := MongoGetImgurRef(mongoId)
	if ref != (ImgurRef{}) {
		numDel := MongoDeleteImgurRef(ref.ImageLink)
		delOK := ImgurDeleteImageRef(ref.ImageDelHash)
		if numDel != 1 {
			log.Println("Error updating Id=", mongoId, "MongoDelete failed")
		}
		if !delOK {
			log.Println("Error updating Id=", mongoId, "ImgurDelete failed")
		}
	}
}
