package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func PrettyPrintStruct(any interface{}) {
	fmt.Printf("%+v\n", any)
}

// Checks message body for valid New Lost Item structure
func ParseLostItemBody(bytes []byte) []byte {
	var generalItem map[string]interface{}
	// var item NewItem
	json.Unmarshal(bytes, &generalItem)
	var cat int
	var tmp string
	var ok bool
	tmp, ok = generalItem["Category"].(string)
	if !ok {
		return nil
	}
	if cat = GetCategoryType(tmp); cat == 0 {
		return nil
	}
	generalItem["Category"] = cat
	// Check for required fields existence
	requiredFields := []string{"Name", "Date", "Location", "User_id"}
	for _, field := range requiredFields {
		if _, ok = generalItem[field]; !ok {
			return nil
		}
	}
	if bytes, err := json.Marshal(generalItem); err != nil {
		return nil
	} else {
		return bytes
	}
}

// Checks message body for valid New Found Item structure
func ParseFoundItemBody(bytes []byte) []byte {
	var generalItem map[string]interface{}
	// var item NewItem
	json.Unmarshal(bytes, &generalItem)
	var cat int
	var tmp string
	var ok bool
	tmp, ok = generalItem["Category"].(string)
	if !ok {
		return nil
	}
	if cat = GetCategoryType(tmp); cat == 0 {
		return nil
	}
	generalItem["Category"] = cat
	requiredFields := []string{"Name", "Date", "Location"}
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
