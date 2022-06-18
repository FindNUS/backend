package main

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ----- MODIFY ITEM -----
* This file contains code to handle the Creation, Update and Deletion of items on the database
 */

// Handle creation of new item
func DoAddItem(msg ItemMsgJSON) interface{} {
	// Unmarshall body
	var item NewItem
	var res interface{}
	json.Unmarshal(msg.Body, &item)
	if item.User_id == "" {
		// Assert that user_id only exists for found items
		res = MongoAddItem(COLL_FOUND, item)
	} else {
		res = MongoAddItem(COLL_LOST, item)
	}
	return res
}

func DoUpdateItem(msg ItemMsgJSON) int64 {
	var item PatchItem
	json.Unmarshal(msg.Body, &item)
	var id string
	var err error
	// Safety check, should not trigger
	if _, ok := msg.Params["Id"]; !ok {
		log.Println("Update failed, item does not exist")
		return -1
	}
	id = msg.Params["Id"][0]
	item.Id, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("ERROR WHILE PATCHING:", err.Error())
		return -1
	}
	id = msg.Params["Id"][0]
	if _, ok := msg.Params["User_id"]; ok {
		item.User_id = msg.Params["User_id"][0]
	}
	// Check which collection the request belongs to
	if item.User_id == "" {
		// Item belongs to FOUND collection
		return MongoPatchItem(COLL_FOUND, item)
	} else {
		// User_id presence implies the msg belongs to LOST collection
		return MongoPatchItem(COLL_LOST, item)
	}
}

func DoDeleteItem(msg ItemMsgJSON) int64 {
	// Assert that msg contains enough parameters
	var item DeletedItem
	var id string
	var err error
	// Safety check, should not trigger
	if _, ok := msg.Params["Id"]; !ok {
		log.Println("Delete failed, item does not exist")
		return -1
	}
	id = msg.Params["Id"][0]
	item.Id, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("ERROR WHILE DELETING:", err.Error())
		return -1
	}
	id = msg.Params["Id"][0]
	if _, ok := msg.Params["User_id"]; ok {
		item.User_id = msg.Params["User_id"][0]
	}
	// Check which collection the deleted item belongs to
	if item.User_id == "" {
		return MongoDeleteItem(COLL_FOUND, item)
	} else {
		return MongoDeleteItem(COLL_LOST, item)
	}
}
