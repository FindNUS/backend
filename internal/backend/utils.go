package main

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func PrettyPrintStruct(any interface{}) {
	log.Printf("%+v\n", any)
}

// Special handlers for Category mapping
// Pass-by pointer to reduce stack memory load
func BodyHandleCategory(body *map[string]interface{}) bool {
	tmp, ok := (*body)["Category"].(string)
	if !ok {
		return false
	}
	// log.Println("Category string received: " + tmp)
	if cat := GetCategoryType(tmp); cat < 0 {
		return false
	}
	// else {
	// 	(*body)["Category"] = cat
	// }
	return true
}

// Date validity handler
// Pass-by pointer to reduce stack memory load
func BodyHandleDate(body *map[string]interface{}) bool {
	tmp, ok := (*body)["Date"].(string)
	if !ok {
		return false
	}
	// log.Println("Date string received: " + tmp)
	_, err := time.Parse("2006-01-02T15:04:05Z", tmp)
	return (err == nil)
}

// Checks message body for valid New Lost Item structure
func ParseLostItemBody(bytes []byte) ([]byte, error) {
	var generalItem map[string]interface{}
	// var item NewItem
	json.Unmarshal(bytes, &generalItem)

	log.Println(generalItem)

	// Handle special parameters
	if !BodyHandleDate(&generalItem) {
		return nil, errors.New("Date is invalid")
	}
	if !BodyHandleCategory(&generalItem) {
		return nil, errors.New("Category is invalid")
	}
	// Check for general required fields existence
	var ok bool
	requiredFields := []string{"Name", "Location", "User_id"}
	for _, field := range requiredFields {
		if _, ok = generalItem[field]; !ok {
			return nil, errors.New("Missing Name, Location &/or User_id")
		}
	}
	bytes, _ = json.Marshal(generalItem)
	return bytes, nil
}

// Checks message body for valid New Found Item structure
func ParseFoundItemBody(bytes []byte) ([]byte, error) {
	var generalItem map[string]interface{}
	// var item NewItem
	json.Unmarshal(bytes, &generalItem)

	log.Println(generalItem)

	// Handle special parameters
	if !BodyHandleDate(&generalItem) {
		return nil, errors.New("Date is invalid")
	}
	if !BodyHandleCategory(&generalItem) {
		return nil, errors.New("Category is invalid")
	}
	// BodyHandleContactMethod(&generalItem)
	// Check for general required fields existence
	var ok bool
	requiredFields := []string{"Name", "Location"}
	for _, field := range requiredFields {
		if _, ok = generalItem[field]; !ok {
			return nil, errors.New("Missing Name &/or Location")
		}
	}
	bytes, _ = json.Marshal(generalItem)
	return bytes, nil
}

// Unwraps the URL's parameters
func GetParams(c *gin.Context) map[string][]string {
	var params map[string][]string
	params = c.Request.URL.Query()
	return params
}

// Wraps a HTTP request context into a JSON format ready to delivery to RabbitMQ
func PrepareMessage(params map[string][]string, body []byte, operation int) ItemMsgJSON {
	var message ItemMsgJSON
	message.Params = params
	message.Body = body
	message.Operation_type = operation
	return message
}
