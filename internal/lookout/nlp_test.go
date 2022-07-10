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

// Test with UAT data that the general function works as expected
func TestNlpGetQuery(t *testing.T) {
	SetupElasticClient()
	items := loadTestItems("item.json")
	for _, tmp := range items {
		var item Item
		bytes, _ := json.Marshal(tmp)
		json.Unmarshal(bytes, &item)
		// NlpGetQuery(item)
		qry := NlpGetQuery(item)
		ElasticLookoutSearch(qry, item.Category)
		log.Println("----- END TEST ITEM -----")
	}
}
