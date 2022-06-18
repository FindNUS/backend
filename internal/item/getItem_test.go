package main

import (
	"log"
	"testing"
)

func TestDoGetItem(t *testing.T) {
	SetupMongo()
	params := make(map[string][]string, 2)
	params["Id"] = []string{"629cc43263533a84f60c4c66"}
	msg := ItemMsgJSON{
		OPERATION_GET_ITEM,
		params,
		nil,
	}
	item := DoGetItem(msg)
	log.Println(item)
	if item == (Item{}) {
		t.Fatal("Item returned is empty")
	}
	// Test LOST collection
	params["User_id"] = []string{"123a"}
	params["Id"] = []string{"62a3f742a972503bb927997c"}
	msg = ItemMsgJSON{
		OPERATION_GET_ITEM,
		params,
		nil,
	}
	item = DoGetItem(msg)
	log.Println(item)
	if item == (Item{}) {
		t.Fatal("Item returned is empty")
	}

}
