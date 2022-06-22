package main

import (
	"encoding/json"
	"log"
)

// Marshaller-Unmarshaller

// Unmarshall message from Message Broker
func UnmarshallMessage(bytes []byte) ItemMsgJSON {
	var msg ItemMsgJSON
	json.Unmarshal(bytes, &msg)
	return msg
}

// Processes the raw message body, remapping certain fields
func ParseNewItemBody(bytes []byte) []byte {
	var generalItem map[string]interface{}
	json.Unmarshal(bytes, &generalItem)
	if !BodyHandleCategory(&generalItem) {
		return nil
	}
	BodyHandleContactMethod(&generalItem)
	bytes, err := json.Marshal(generalItem)
	if err != nil {
		return nil
	}
	return bytes
}

// Processes MongoDB item and remaps certain fields
func ParseGetItemBody(tmp map[string]interface{}) Item {
	if tmp == nil {
		return Item{}
	}
	var item Item
	// Transform category
	tmp["Category"] = GetCategoryString(tmp["Category"].(int32))
	// Transform contact method
	if val, ok := tmp["Contact_method"]; ok {
		tmp["Contact_method"] = GetContactString(val.(int32))
	}
	tmp["Id"] = tmp["_id"]
	b, err := json.Marshal(tmp)
	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(b, &item)
	return item
}

// Special handlers for Category mapping
// Pass-by pointer to reduce stack memory load
func BodyHandleCategory(body *map[string]interface{}) bool {
	tmp, ok := (*body)["Category"].(string)
	if !ok {
		return false
	}
	if cat := GetCategoryType(tmp); cat < 0 {
		return false
	} else {
		(*body)["Category"] = cat
	}
	return true
}

// Special handler for Contact_method mapping
// Pass-by pointer to reduce stack memory load
func BodyHandleContactMethod(body *map[string]interface{}) {
	tmp, ok := (*body)["Contact_method"].(string)
	if !ok {
		return
	}
	// TODO: Invalid contact method will be processed as "Unspecified"
	cat := GetContactMethod(tmp)
	(*body)["Contact_method"] = cat
}
