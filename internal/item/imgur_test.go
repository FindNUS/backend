package main

import (
	"bufio"
	"log"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Loads a base64 image from a text file and returns the base64 string
func debugLoadTestImageB64(filename string) string {
	var f *os.File
	var err error
	f, err = os.Open("./test/" + filename)
	if err != nil {
		log.Fatal("Error opening file:", err.Error())
	}
	var res string
	scanner := bufio.NewScanner(f)
	defer f.Close()
	for scanner.Scan() {
		res = scanner.Text()
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	return res
}

func TestImgurCRUD(t *testing.T) {
	SetupMongo()
	SetupImgur()
	// Test add imgur API
	b64 := debugLoadTestImageB64("tux.txt")
	link, hash := ImgurAddNewImage(b64)
	if link == "" {
		t.Fatal("AddNewImage returned empty link")
	}
	log.Println(link, hash)
	log.Println("ADD IMGUR IMAGE PASS")

	// Test Mongo Store API
	res := MongoStoreImgurRef(link, hash)
	objId := res.(primitive.ObjectID).Hex()
	if res == nil {
		t.Fatal("MongoStoreImgurRef returned empty Mongo Id")
	}
	PrettyPrintStruct(res)
	log.Println("ADD IMGURREF TO MONGO PASS")

	// Test Get ImgurRef API
	imgRef := MongoGetImgurRef(objId)
	if imgRef == (ImgurRef{}) {
		t.Fatal("ImgurGetImageRef returned empty")
	}
	log.Println("GET IMGURREF PASS")

	// Test Mongo Delete ImgurRef API
	numDel := MongoDeleteImgurRef(link)
	if numDel != 1 {
		t.Fail()
		log.Println("MongoDeleteImgurRef failed -- expected 1 deletion, got", numDel)
	}
	log.Println("DELETE MongoDeleteImgurRef PASS")

	// Test Imgur delete image
	ok := ImgurDeleteImageRef(hash)
	if !ok {
		t.Fail()
		log.Println("ImgurDeleteImageRef fail")
	}
	log.Println("DELETE ImgurDeleteImageRef PASS")

}
