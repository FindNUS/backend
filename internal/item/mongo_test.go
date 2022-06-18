package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func loadTestItems(filename string) []map[string]interface{} {
	var f *os.File
	var err error
	f, err = os.Open("./test/" + filename)
	if err != nil {
		log.Fatalf(err.Error())
	}
	data, err := ioutil.ReadAll(f)
	var res []map[string]interface{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res
}

func buildItemMsgJson(params map[string][]string, body []byte) ItemMsgJSON {
	var msg ItemMsgJSON
	if params != nil {
		msg.Params = params
	}
	if body != nil {
		msg.Body = body
	}
	return msg
}

func TestMongoGetManyItems(t *testing.T) {
	SetupMongo()
	args := make(map[string][]string)
	// Test limit
	args["limit"] = []string{"10"}
	items := MongoGetManyItems(COLL_FOUND, args)
	if len(items) != 10 {
		t.FailNow()
	}
	for _, item := range items {
		PrettyPrintStruct(item)
	}
	// Test offset
	args["limit"] = []string{"5"}
	args["offset"] = []string{"5"}
	items = MongoGetManyItems(COLL_FOUND, args)
	if len(items) != 5 {
		t.FailNow()
	}
	for _, item := range items {
		PrettyPrintStruct(item)
	}
	// Test category filterimg
	args["category"] = []string{"Electronics", "Notes"}
	items = MongoGetManyItems(COLL_DEBUG, args)
	for _, item := range items {
		if !(item.Category == 3 || item.Category == 2) {
			t.Fatalf("Category filter query returned wrong item category")
		}
		PrettyPrintStruct(item)
	}
}
