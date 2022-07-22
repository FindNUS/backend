package main

import (
	"encoding/json"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Checks an `Item` if its expected key-value is met
func checkUpdateParamsEqual(key string, expected string, object Item) bool {
	// log.Println("Checking: ")
	// PrettyPrintStruct(object)
	tmp, _ := json.Marshal(object)
	var item map[string]interface{}
	json.Unmarshal(tmp, &item)
	if val, ok := item[key]; !ok {
		log.Println(key, "not found!")
		return false
	} else {
		return (expected == val)
	}

}

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

	// Patch Test for Category
	var verifyItem Item

	// UPDATE 1 UNMAPPABLE PARAM TEST
	log.Println("Testing PATCH for ONE STRING parameters")
	updateItem := map[string]string{
		"Id":       id.Hex(),
		"User_id":  userid,
		"Location": "New Location",
	}
	bytes, err := json.Marshal(updateItem)
	if err != nil {
		log.Println(err.Error())
	}
	msg = buildItemMsgJson(dummyparams, bytes)
	if _, err := DoUpdateItem(msg); err != nil {
		t.Fatal("Patch item failed:", err.Error())
	}
	// Check that item was correctly updated
	verifyItem = MongoGetItem(COLL_LOST, id.Hex(), userid)
	if !checkUpdateParamsEqual("Location", updateItem["Location"], verifyItem) {
		t.Fail()
		t.Log("Update location only failed!")
	}
	// Additional verification for particular parameters
	if !checkUpdateParamsEqual("Pluscode", item[0]["Pluscode"].(string), verifyItem) {
		t.Fail()
		t.Log("Pluscode missing!")
	}
	log.Println("Testing PATCH for LOOKOUT parameter")
	updateItemMixed := PatchItem{
		Id:      id,
		User_id: userid,
		Lookout: true,
	}
	bytes, err = json.Marshal(updateItemMixed)
	if err != nil {
		log.Println(err.Error())
	}
	msg = buildItemMsgJson(dummyparams, bytes)
	if _, err := DoUpdateItem(msg); err != nil {
		t.Fatal("Patch item failed:", err.Error())
	}
	// Check that item was correctly updated
	verifyItem = MongoGetItem(COLL_LOST, id.Hex(), userid)
	if !verifyItem.Lookout {
		t.Fail()
		t.Log("Update lookout only failed!")
	}
	// UPDATE MAPPABLE ITEMS TEST
	log.Println("Testing PATCH for MAPPABLE parameters")
	updateItem = map[string]string{
		"Id":             id.Hex(),
		"User_id":        userid,
		"Contact_method": "Telegram",
		"Category":       "Notes",
	}
	bytes, err = json.Marshal(updateItem)
	if err != nil {
		log.Println(err.Error())
	}
	msg = buildItemMsgJson(dummyparams, bytes)
	if _, err := DoUpdateItem(msg); err != nil {
		t.Fatal("Patch item failed:", err.Error())
	}
	// Check that item was correctly updated
	verifyItem = MongoGetItem(COLL_LOST, id.Hex(), userid)
	if !checkUpdateParamsEqual("Category", updateItem["Category"], verifyItem) {
		t.Fail()
		t.Log("Update Category failed!")
	}
	if !checkUpdateParamsEqual("Contact_method", updateItem["Contact_method"], verifyItem) {
		t.Fail()
		t.Log("Update Contact_method failed!")
	}

	// UPDATE MIXABLE ITEMS TEST
	log.Println("Testing PATCH for MIXED parameters with ADDITION")
	updateItem = map[string]string{
		"Id":             id.Hex(),
		"User_id":        userid,
		"Contact_method": "Whatsapp",
		"Category":       "Etc",
		"Name":           "Debug Add Item Unit Test Final",
		"Item_details":   "New parameter added",
	}
	bytes, err = json.Marshal(updateItem)
	if err != nil {
		log.Println(err.Error())
	}
	msg = buildItemMsgJson(dummyparams, bytes)
	if _, err := DoUpdateItem(msg); err != nil {
		t.Fatal("Patch item failed:", err.Error())
	}
	// Check that item was correctly updated
	verifyItem = MongoGetItem(COLL_LOST, id.Hex(), userid)
	PrettyPrintStruct(verifyItem)
	if !checkUpdateParamsEqual("Category", updateItem["Category"], verifyItem) {
		t.Fail()
		t.Log("Update Category failed!")
	}
	if !checkUpdateParamsEqual("Contact_method", updateItem["Contact_method"], verifyItem) {
		t.Fail()
		t.Log("Update Contact_method failed!")
	}
	if !checkUpdateParamsEqual("Name", updateItem["Name"], verifyItem) {
		t.Fail()
		t.Log("Update Name failed!")
	}
	if !checkUpdateParamsEqual("Item_details", updateItem["Item_details"], verifyItem) {
		t.Fail()
		t.Log("Update Item_details (new parameter) failed!")
	}

	// Delete test
	if _, err := DoDeleteItem(msg); err != nil {
		log.Fatal("Delete fail for ", id, "; Error:", err.Error())
	}
	log.Println("Testing cleanup done!")
}
