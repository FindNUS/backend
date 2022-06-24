package main

import (
	"encoding/json"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestElasticCRUD(t *testing.T) {
	SetupElasticClient()
	ElasticInitIndex()
	testItems := loadTestItems("debug_es_item.json")
	bytes, _ := json.Marshal(testItems[0])
	// ADD TEST
	var item ElasticItem
	json.Unmarshal(bytes, &item)
	ElasticAddItem(item)
	log.Println("Elastic Add OK")
	// time.Sleep(500 * time.Millisecond)
	// GET TEST
	item = ElasticGetItem(item.Id)
	if item == (ElasticItem{}) {
		t.Fatalf("Elastic GET after ADD failed. Empty item returned")
	}
	PrettyPrintStruct(item)
	log.Println("Elastic GET OK")
	// PATCH TEST
	prevName := item.Name
	newName := "FOOBAR"
	item.Name = newName
	ElasticUpdateItem(item)
	item = ElasticGetItem(item.Id)
	if item.Name == prevName || item.Name != newName {
		t.Fatalf("Elastic UPDATE failed. Wrong name returned. \n Expected %s, got %s", newName, item.Name)
	}
	log.Println("Elastic UPDATE OK")
	// DELETE TEST
	del := ElasticDeleteItem(item.Id)
	if del != 1 {
		t.Fatalf("Elastic DELETE failed. Wrong name returned. \n Expected %d, got %d", 1, del)
	}
	if ElasticGetItem(item.Id) != (ElasticItem{}) {
		t.Fatalf("Elastic DELETE failed. Item still exists on the database!")
	}
	log.Println("Elastic DELETE OK")
}

// func TestElasticDeleteItem(t *testing.T) {
// 	SetupElasticClient()
// 	ElasticInitIndex()
// 	testItems := loadTestItems("debug_es_item.json")
// 	bytes, _ := json.Marshal(testItems[0])
// 	var item ElasticItem
// 	json.Unmarshal(bytes, &item)
// 	ElasticDeleteItem(item.Id)
// }

// func TestElasticGetItem(t *testing.T) {
// 	SetupElasticClient()
// 	ElasticInitIndex()
// 	testItems := loadTestItems("debug_es_item.json")
// 	bytes, _ := json.Marshal(testItems[0])
// 	var item ElasticItem
// 	json.Unmarshal(bytes, &item)
// 	//ADD OK
// 	ElasticGetItem(item.Id)
// }

func TestElasticSearchGeneral(t *testing.T) {
	// This test is run after CRUD is validated
	SetupElasticClient()
	ElasticInitIndex()
	testItems := loadTestItems("debug_es_item.json")
	bytes, _ := json.Marshal(testItems[0])
	var item ElasticItem
	json.Unmarshal(bytes, &item)
	ElasticAddItem(item)
	q := "ble item utwn"
	log.Println("Executing search...")
	items := ElasticSearchGeneral(q)
	PrettyPrintStruct(items)
	// Cleanup
	ElasticDeleteItem(item.Id)
}

func TestElasticParseBody(t *testing.T) {
	item := loadTestItems("debug_add_item.json")
	bytes, _ := json.Marshal(item[0])
	// msg := buildItemMsgJson(nil, bytes)
	randId := primitive.NewObjectID()
	var itemStruct Item
	json.Unmarshal(bytes, &itemStruct)
	esItem := ElasticParseBody(itemStruct, randId)
	if esItem == (ElasticItem{}) {
		t.Fatalf("Parse item body failed, parsed esItem returned empty!")
	}
}
