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

func TestParseFoundItemBody(t *testing.T) {
	testdata := loadTestItems("valid_found_items.json")
	for _, item := range testdata {
		bytes, _ := json.Marshal(item)
		if ans := ParseFoundItemBody(bytes); ans == nil {
			t.Log("Found item wrongly flagged as invalid: ", item)
			t.Fail()
		}
	}
	testdata = loadTestItems("invalid_found_items.json")
	for _, item := range testdata {
		bytes, _ := json.Marshal(item)
		if ans := ParseFoundItemBody(bytes); ans != nil {
			t.Log("Found item wrongly flagged as valid: ", item)
			t.Fail()
		}
	}
}

func TestParseLostItemBody(t *testing.T) {
	testdata := loadTestItems("valid_lost_items.json")
	for _, item := range testdata {
		bytes, _ := json.Marshal(item)
		if ans := ParseLostItemBody(bytes); ans == nil {
			t.Log("Lost item wrongly flagged as invalid: ", item)
			t.Fail()
		}
	}
	testdata = loadTestItems("invalid_lost_items.json")
	for _, item := range testdata {
		bytes, _ := json.Marshal(item)
		if ans := ParseLostItemBody(bytes); ans != nil {
			t.Log("Lost item wrongly flagged as valid: ", item)
			t.Fail()
		}
	}
}
