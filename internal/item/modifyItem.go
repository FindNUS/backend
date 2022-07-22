package main

import (
	"encoding/json"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ----- MODIFY ITEM -----
* This file contains code to handle the Creation, Update and Deletion of items on the database
 */

// Handle creation of new item
// Returns a string representing which ItemCollection the item was added to and its Id
func DoAddItem(msg ItemMsgJSON) (ItemCollections, primitive.ObjectID) {
	// Unmarshall body
	var item NewItem
	var res primitive.ObjectID
	var interf interface{} // safety
	collection := COLL_LOST
	body := ParseNewItemBody(msg.Body)
	json.Unmarshal(body, &item)
	if item.User_id == "" {
		// Assert that user_id only exists for found items
		interf = MongoAddItem(COLL_FOUND, item)
		collection = COLL_FOUND
	} else {
		interf = MongoAddItem(COLL_LOST, item)
		collection = COLL_LOST
	}
	var ok bool
	if res, ok = interf.(primitive.ObjectID); !ok {
		log.Println("Error casting object ID in DoAddItem.")
		return collection, primitive.NilObjectID
	}
	log.Println("ID of added item:", res)
	return collection, res
}

// Updates an Item and returns its mongoDB id and an error, if any
// Does not return the string when User_id exists (avoid false deletion in FOUND ES)
func DoUpdateItem(msg ItemMsgJSON) (string, error) {
	var err error
	var item PatchItem
	body := ParseUpdateItemBody(msg.Body)

	err = json.Unmarshal(body, &item)
	if err != nil {
		log.Println(err.Error())
	}

	var id string
	// Safety check, should not trigger
	if _, ok := msg.Params["Id"]; !ok {
		return "", errors.New("Update item failed, item does not exist")
	}
	id = msg.Params["Id"][0]
	if id == "" {
		return "", errors.New("ERROR WHILE PATCHING: NO ID FOUND")
	}
	item.Id, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", errors.New("ERROR WHILE PATCHING: " + err.Error())
	}
	if _, ok := msg.Params["User_id"]; ok {
		item.User_id = msg.Params["User_id"][0]
		id = "" // Prevent ElasticSearch operation
	}
	// Check which collection the request belongs to
	numModify := int64(0)
	if item.User_id == "" {
		// Item belongs to FOUND collection
		numModify = MongoPatchItem(COLL_FOUND, item)
	} else {
		// User_id presence implies the msg belongs to LOST collection
		numModify = MongoPatchItem(COLL_LOST, item)
	}
	if numModify != 1 {
		log.Println("WARNING: Potential error in DoUpdateItem. Expected 1 modified item, got", numModify)
		log.Println("Affected update id:", item.Id)
	}
	return id, nil
}

// Deletes an Item and returns its mongoDB id and an error, if any
// Does not return the string when User_id exists (avoid false deletion in FOUND ES)
func DoDeleteItem(msg ItemMsgJSON) (string, error) {
	// Assert that msg contains enough parameters
	var item DeletedItem
	var id string
	var err error
	// Safety check, should not trigger
	if _, ok := msg.Params["Id"]; !ok {
		return "", errors.New("Delete failed, item does not exist")
	}
	id = msg.Params["Id"][0]
	item.Id, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", errors.New("Error while deleting: " + err.Error())
	}
	// User_id check to determine collection and execute Imgur deletion if necessary
	if _, ok := msg.Params["User_id"]; ok {
		item.User_id = msg.Params["User_id"][0]
		id = "" // Prevent ElasticSearch operation
	} else {
		// Delete image of item, if needed
		ImgurDeleteImageFromId(id)
	}

	// Based which collection the deleted item belongs to
	if item.User_id == "" {
		MongoDeleteItem(COLL_FOUND, item)
	} else {
		MongoDeleteItem(COLL_LOST, item)
	}
	return id, nil
}
