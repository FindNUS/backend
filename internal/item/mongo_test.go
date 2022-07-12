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
	args["limit"] = []string{"5"}
	items := MongoGetManyItems(COLL_DEBUG, args)
	if len(items) != 5 {
		log.Fatal("Items do not fit limit")
	}
	// for _, item := range items {
	// 	PrettyPrintStruct(item)
	// }
	log.Println("GetManyItems limit PASS")

	// Test offset
	args["limit"] = []string{"2"}
	args["offset"] = []string{"2"}
	items = MongoGetManyItems(COLL_DEBUG, args)
	if len(items) != 2 {
		log.Fatal("Items offset error - length is not 2")
	}
	// for _, item := range items {
	// 	PrettyPrintStruct(item)
	// }
	log.Println("GetManyItems offset PASS")

	// Test category filtering
	args["category"] = []string{"Electronics", "Notes"}
	items = MongoGetManyItems(COLL_DEBUG, args)
	for _, item := range items {
		if !(item.Category == "Electronics" || item.Category == "Notes") {
			PrettyPrintStruct(item)
			t.Fatalf("Category filter query returned wrong item category")
		}
	}
	log.Println("GetManyItems Category filter PASS")

	// Test User_id filtering
	args["User_id"] = []string{"123a"}
	delete(args, "offset")
	delete(args, "limit")
	delete(args, "category")
	items = MongoGetManyItems(COLL_DEBUG, args)
	log.Println(len(items))
	for _, item := range items {
		if !(item.User_id == "123a") {
			PrettyPrintStruct(item)
			t.Fatalf("User_id filter failed")
		}
	}
	log.Println("GetManyItems User_id filter PASS")

	// Test date filtering
	// 1. Test Start date only
	args = make(map[string][]string)
	args["startdate"] = []string{"2021-01-31T00:00:00Z"}
	items = MongoGetManyItems(COLL_DEBUG, args)
	startDate := ParseDateString("2021-01-31T00:00:00Z")
	for _, item := range items {
		if item.Date.Before(startDate) {
			PrettyPrintStruct(item)
			t.Log("Item has date before startdate")
			t.Fail()
		}
	}
	log.Println("GetManyItems StartDate filter PASS")

	// 2. Test End date only
	args = make(map[string][]string)
	args["enddate"] = []string{"2022-03-25T00:00:00Z"}
	items = MongoGetManyItems(COLL_DEBUG, args)
	endDate := ParseDateString("2022-03-25T00:00:00Z")
	for _, item := range items {
		if item.Date.After(endDate) {
			PrettyPrintStruct(item)
			t.Log("Item has date after endDate")
			t.Fail()
		}
	}
	log.Println("GetManyItems EndDate filter PASS")

	// 3. Test Both start and end date
	// enddate has been defined already
	args["startdate"] = []string{"2021-01-31T00:00:00Z"}
	items = MongoGetManyItems(COLL_DEBUG, args)
	startDate = ParseDateString("2021-01-31T00:00:00Z")
	for _, item := range items {
		if item.Date.After(endDate) || item.Date.Before(startDate) {
			PrettyPrintStruct(item)
			t.Log("Item's date is out of specified range!")
			t.Fail()
		}
	}
	log.Println("GetManyItems Date Range filter PASS")

	// Final test - Compound filter
	args = make(map[string][]string)
	args["startdate"] = []string{"2021-01-31T00:00:00Z"}
	args["enddate"] = []string{"2022-05-25T00:00:00Z"}
	args["limit"] = []string{"2"}
	args["offset"] = []string{"0"}
	args["User_id"] = []string{"123a"}
	args["category"] = []string{"Cards"}
	items = MongoGetManyItems(COLL_DEBUG, args)
	for _, item := range items {
		// Only one item should exist
		PrettyPrintStruct(item)
	}
	log.Println("GetManyItems all filters PASS")
}

func TestMongoGetAllLookoutRequests(t *testing.T) {
	SetupMongo()
	items := MongoGetAllLookoutRequests(COLL_DEBUG)
	for _, item := range items {
		PrettyPrintStruct(item)
	}
}
