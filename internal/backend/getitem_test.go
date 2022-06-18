package main

import (
	"encoding/json"
	"testing"
)

func TestParseManyItemsRPC(t *testing.T) {
	// Mock list of items
	itemsInterf := loadTestItems("mock_items_rpc.json")
	bytes, _ := json.Marshal(itemsInterf)

	items := ParseGetManyItemsRPC(bytes)
	for _, item := range items {
		if item == (Item{}) {
			t.Fail()
		}
		PrettyPrintStruct(item)
	}
}

func TestParseSingleItemRPC(t *testing.T) {
	// Mock list of items
	itemsInterf := loadTestItems("mock_items_rpc.json")
	for _, item := range itemsInterf {
		res := ParseGetOneItemRPC(item)
		if res == (Item{}) {
			t.FailNow()
		}
		PrettyPrintStruct(res)
	}
}
