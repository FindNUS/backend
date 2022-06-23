package main

import (
	"encoding/json"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAddUpdateDelete(t *testing.T) {
	SetupMongo()
	// Test routine for Lost item (with user id)
	item := loadTestItems("debug_add_item.json")
	bytes, _ := json.Marshal(item[0])
	msg := buildItemMsgJson(nil, bytes)
	// Add test
	_, id := DoAddItem(msg)
	if id == primitive.NilObjectID {
		t.Fatal("Add item returned nil ObjectId", item)
	}
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
	if _, err := DoUpdateItem(msg); err != nil {
		t.Fatal("Patch item failed:", err.Error())
	}
	// Delete test
	if _, err := DoDeleteItem(msg); err != nil {
		log.Fatal("Delete fail for ", id, "; Error:", err.Error())
	}
}
