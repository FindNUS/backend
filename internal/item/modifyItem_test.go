package main

import (
	"encoding/json"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddUpdateDelete(t *testing.T) {
	// Test routine for Lost item (with user id)
	item := loadTestItems("debug_add_item.json")
	SetupMongo()
	bytes, _ := json.Marshal(item[0])
	msg := buildItemMsgJson(nil, bytes)
	// Add test
	res := DoAddItem(msg)
	if res == nil {
		t.Fatal("Add item returned nil", item)
	}
	id, ok := res.(primitive.ObjectID)
	if !ok {
		t.Fatal("Object ID error")
	}
	// log.Println("Object Id for testing: ", id)

	// Create dummy parameters for proper parsing
	dummyparams := make(map[string][]string)
	dummyparams["Id"] = []string{id.Hex()}
	userid, _ := item[0]["User_id"].(string)
	dummyparams["User_id"] = []string{userid}
	// Simulate a change in item details
	item[0]["Id"] = id.Hex()
	item[0]["Location"] = "New Location"
	// Patch test
	bytes, _ = json.Marshal(item[0])
	msg = buildItemMsgJson(dummyparams, bytes)
	if numUpdate := DoUpdateItem(msg); numUpdate != 1 {
		t.Fatal("Patch item failed, number of modified: ", numUpdate)
	}
	// Delete test
	if numDel := DoDeleteItem(msg); numDel != 1 {
		log.Fatal("Delete fail for ", id, ". Expected 1 delete but got ", numDel)
	}
}
