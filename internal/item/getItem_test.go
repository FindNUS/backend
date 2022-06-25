package main

import (
	"testing"
)

// Warning: This test will break if we remove the debugging items
func TestDoGetItem(t *testing.T) {
	SetupMongo()
	// Test FOUND collection
	params := make(map[string][]string, 2)
	params["Id"] = []string{"62b47abcf9904afd25588691"}
	msg := ItemMsgJSON{
		OPERATION_GET_ITEM,
		params,
		nil,
	}
	item := DoGetItem(msg)
	PrettyPrintStruct(item)
	if item == (Item{}) {
		t.Fatal("Item returned is empty")
	}
	// Test LOST collection
	params["User_id"] = []string{"123a"}
	params["Id"] = []string{"62b47a2f6511b87bb640f118"}
	msg = ItemMsgJSON{
		OPERATION_GET_ITEM,
		params,
		nil,
	}
	item = DoGetItem(msg)
	PrettyPrintStruct(item)
	if item == (Item{}) {
		t.Fatal("Item returned is empty")
	}
	// Test for invalid Id with User_id
	params["Id"] = []string{"foo"}
	msg = ItemMsgJSON{
		OPERATION_GET_ITEM,
		params,
		nil,
	}
	item = DoGetItem(msg)
	if item != (Item{}) {
		t.Fatalf("Expected empty item")
	}
	// Test for invalid Id without User_id
	delete(params, "User_id")
	msg = ItemMsgJSON{
		OPERATION_GET_ITEM,
		params,
		nil,
	}
	item = DoGetItem(msg)
	if item != (Item{}) {
		t.Fatalf("Expected empty item")
	}
}
