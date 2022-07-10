package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// File must be json and located in test folder
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

// Redundant test for Body Parsing Function
func TestParseItemBody(t *testing.T) {
	// FOUND ITEMS
	testdata := loadTestItems("valid_found_items.json")
	for _, item := range testdata {
		bytes, _ := json.Marshal(item)
		if _, err := ParseItemBody(bytes); err != nil {
			t.Log("Found item wrongly flagged as invalid: ", item)
			t.Log("Error: ", err.Error())
			t.Fail()
		}
	}
	testdata = loadTestItems("invalid_found_items.json")
	for _, item := range testdata {
		bytes, _ := json.Marshal(item)
		if _, err := ParseItemBody(bytes); err == nil {
			t.Log("Found item wrongly flagged as valid: ", item)
			t.Log(err.Error())
			t.Fail()
		}
	}
	// LOST ITEMS
	testdata = loadTestItems("valid_lost_items.json")
	for _, item := range testdata {
		PrettyPrintStruct(item)
		bytes, _ := json.Marshal(item)
		if _, err := ParseItemBody(bytes); err != nil {
			t.Log("Lost item wrongly flagged as invalid: ", item)
			t.Log("Error: ", err.Error())
			t.Fail()
		}
	}
	testdata = loadTestItems("invalid_lost_items.json")
	for _, item := range testdata {
		bytes, _ := json.Marshal(item)
		if _, err := ParseItemBody(bytes); err == nil {
			t.Log("Lost item wrongly flagged as valid: ", item)
			t.Log(err.Error())
			t.Fail()
		}
	}
}
func TestValidatePeekParams(t *testing.T) {
	loadRaw := loadTestItems("valid_peek_url.json")
	b, _ := json.Marshal(loadRaw)
	var items []map[string][]string
	json.Unmarshal(b, &items)
	var err error
	for _, item := range items {
		// PrettyPrintStruct(item)
		err = ValidatePeekParams(item)
		if err != nil {
			t.Fail()
			t.Log("ValidatePeekParams failed when it is supposed to pass!\n Failed query:", item)
		}
	}
	loadRaw = loadTestItems("invalid_peek_url.json")
	b, _ = json.Marshal(loadRaw)
	json.Unmarshal(b, &items)
	for _, item := range items {
		// PrettyPrintStruct(item)
		err = ValidatePeekParams(item)
		if err == nil {
			t.Fail()
			t.Log("ValidatePeekParams passed when it is supposed to pass!\n Failed query:", item)
		}
	}
}
