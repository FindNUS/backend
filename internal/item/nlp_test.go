package main

import (
	"encoding/json"
	"testing"
)

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
		// log.Println("----- END TEST ITEM -----")
	}
}
