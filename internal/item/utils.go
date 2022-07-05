package main

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// Category Handler
	if !BodyHandleCategory(&generalItem) {
		return nil
	}
	// Other special field handlers
	BodyHandleContactMethod(&generalItem)
	BodyHandleImage_Base64(&generalItem)
	bytes, err := json.Marshal(generalItem)
	if err != nil {
		log.Println("ParseNewItem failed, returning nil due to:", err.Error())
		return nil
	}
	return bytes
}

// Processes the raw message body, remapping certain fields.
// Functionally identical to ParseNewItemBody, but does not enforce the existence of certain parameters
func ParseUpdateItemBody(bytes []byte) []byte {
	var generalItem map[string]interface{}
	json.Unmarshal(bytes, &generalItem)
	// Category Handler
	BodyHandleCategory(&generalItem)
	// Other special field handlers
	BodyHandleContactMethod(&generalItem)
	BodyHandleImage_Base64(&generalItem)
	bytes, err := json.Marshal(generalItem)
	if err != nil {
		log.Println("ParseUpdateItemBody failed, returning nil due to:", err.Error())
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

// Parses an ElasticSearch URL query for use in the actual ElasticSearch
func GetElasticQuery(params map[string][]string) string {
	query := ""
	if param, ok := params["query"]; ok {
		query = param[0]
	}
	return query
}

// Special handler for Image_base64 parameter.
// Updates an item's Imgur reference
func BodyHandleImage_Base64(body *map[string]interface{}) {
	base64str, ok := (*body)["Image_base64"].(string)
	if !ok {
		return
	}
	// Check if the Image is an update to an existing item
	// If yes, delete the image reference
	objId, ok := (*body)["Id"].(primitive.ObjectID)
	if ok {
		ref := MongoGetImgurRef(objId.Hex())
		numDel := MongoDeleteImgurRef(ref.ImageLink)
		delOK := ImgurDeleteImageRef(ref.ImageDelHash)
		if numDel != 1 {
			log.Println("Error updating Id=", objId, "MongoDelete failed")
		}
		if !delOK {
			log.Println("Error updating Id=", objId, "ImgurDelete failed")
		}
	}
	// Remove the large base64 parameter
	delete((*body), "Image_base64")
	link, hash := ImgurAddNewImage(base64str)
	if link == "" {
		return
	}
	log.Println("Imgur link:", link, "Imgur hash:", hash)
	// Set the image url
	(*body)["Image_url"] = link
	newId := MongoStoreImgurRef(link, hash).(primitive.ObjectID)
	if newId == primitive.NilObjectID {
		log.Println("WARNING: Possible error storing imgurRef for", link)
	}
	log.Println("Final body:", *body)
}
