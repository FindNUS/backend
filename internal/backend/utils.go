package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func PrettyPrintStruct(any interface{}) {
	fmt.Printf("%+v\n", any)
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

// Date validity handler
// Pass-by pointer to reduce stack memory load
func BodyHandleDate(body *map[string]interface{}) bool {
	tmp, ok := (*body)["Date"].(string)
	if !ok {
		return false
	}
	_, err := time.Parse("2006-01-02T15:04:05Z07:00", tmp)
	return (err == nil)
}

// Checks message body for valid New Lost Item structure
func ParseLostItemBody(bytes []byte) []byte {
	var generalItem map[string]interface{}
	// var item NewItem
	json.Unmarshal(bytes, &generalItem)
	// Handle special parameters
	if !BodyHandleDate(&generalItem) {
		return nil
	}
	if !BodyHandleCategory(&generalItem) {
		return nil
	}
	BodyHandleContactMethod(&generalItem)
	// Check for general required fields existence
	var ok bool
	requiredFields := []string{"Name", "Location", "User_id"}
	for _, field := range requiredFields {
		if _, ok = generalItem[field]; !ok {
			return nil
		}
	}
	bytes, _ = json.Marshal(generalItem)
	return bytes
}

// Checks message body for valid New Found Item structure
func ParseFoundItemBody(bytes []byte) []byte {
	var generalItem map[string]interface{}
	// var item NewItem
	json.Unmarshal(bytes, &generalItem)
	// Handle special parameters
	if !BodyHandleDate(&generalItem) {
		return nil
	}
	if !BodyHandleCategory(&generalItem) {
		return nil
	}
	BodyHandleContactMethod(&generalItem)
	// Check for general required fields existence
	var ok bool
	requiredFields := []string{"Name", "Location"}
	for _, field := range requiredFields {
		if _, ok = generalItem[field]; !ok {
			return nil
		}
	}
	bytes, _ = json.Marshal(generalItem)
	return bytes
}

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
