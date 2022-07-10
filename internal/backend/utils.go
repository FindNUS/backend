package main

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrettyPrintStruct(any interface{}) {
	log.Printf("%+v\n", any)
}

// Special handlers for Category mapping
// Pass-by pointer to reduce stack memory load
func BodyValidateCategory(body *map[string]interface{}) bool {
	tmp, ok := (*body)["Category"].(string)
	if !ok {
		return false
	}
	// log.Println("Category string received: " + tmp)
	if cat := GetCategoryType(tmp); cat < 0 {
		return false
	}
	return true
}

// Date validity handler
// Pass-by pointer to reduce stack memory load
func BodyValidateDate(body *map[string]interface{}) bool {
	tmp, ok := (*body)["Date"].(string)
	if !ok {
		return false
	}
	// log.Println("Date string received: " + tmp)
	_, err := time.Parse("2006-01-02T15:04:05Z", tmp)
	return (err == nil)
}

// Validate Lookout field in the body
// Lookout field should only exist on items that have User_id
// If enforceDefault is true, missing Lookout values for Lost items will be defaulted to false
// Returns an error if there is an issue with field
func BodyValidateLookout(body *map[string]interface{}, enforceDefault bool) error {
	_, hasUserId := (*body)["User_id"].(string)
	if hasUserId {
		// If Lookout does not exist, create the default false value
		value, hasLookout := (*body)["Lookout"]
		if hasLookout {
			// Check if value is valid
			if _, ok := value.(bool); ok {
				log.Println("Request has user_id and lookout -- no issue")
				return nil
			} else {
				return errors.New("Lookout value in request payload is not in boolean form.")
			}
		}
		// Lookout expected but not found.
		if enforceDefault {
			// This is likely a new item. Default lookout to false
			log.Println("No Lookout field detected for item with User_id. Adding default Lookout=false.")
			(*body)["Lookout"] = false
			return nil
		}
		// Ignore the request. This is likely a patch item
		return nil
	}
	// No User_id. Lookout should not exist
	_, hasLookout := (*body)["Lookout"].(bool)
	if hasLookout {
		// Found item has a lookout parameter. Reject
		log.Println("Non-Lost item has a Lookout field. This should not exist!")
		return errors.New("Non-Lost item has a Lookout field. This should not exist!")
	}
	return nil
}

// Checks a New Item's message body for valid New Found Item structure
func ParseItemBody(bytes []byte) ([]byte, error) {
	var generalItem map[string]interface{}
	json.Unmarshal(bytes, &generalItem)
	// Handle special parameters
	if !BodyValidateDate(&generalItem) {
		return nil, errors.New("Date is invalid")
	}
	if !BodyValidateCategory(&generalItem) {
		return nil, errors.New("Category is invalid")
	}
	// Check for general required fields existence
	var ok bool
	requiredFields := []string{"Name", "Location"}
	for _, field := range requiredFields {
		if _, ok = generalItem[field]; !ok {
			return nil, errors.New("Missing Name &/or Location")
		}
	}

	// Check for Lookout field
	if err := BodyValidateLookout(&generalItem, true); err != nil {
		return nil, err
	}

	bytes, _ = json.Marshal(generalItem)
	return bytes, nil
}

func ParseUpdateItemBody(bytes []byte) ([]byte, error) {
	var generalItem map[string]interface{}
	// var item NewItem
	json.Unmarshal(bytes, &generalItem)

	// log.Println(generalItem)

	// Handle special parameters
	if _, ok := generalItem["Date"]; ok {
		if !BodyValidateDate(&generalItem) {
			return nil, errors.New("Date is invalid")
		}
	}
	if _, ok := generalItem["Category"]; ok {
		if !BodyValidateCategory(&generalItem) {
			return nil, errors.New("Category is invalid")
		}
	}

	// Check for general required fielsds existence
	if err := BodyValidateLookout(&generalItem, false); err != nil {
		return nil, err
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

// Checks if offset & limit are proper
func ValidatePeekParams(params map[string][]string) error {
	if paramArr, ok := params["offset"]; ok {
		tmp := paramArr[0]
		num, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return err
		}
		if num <= 0 {
			return errors.New("offset cannot be <= 0!")
		}
	}
	if paramArr, ok := params["limit"]; ok {
		tmp := paramArr[0]
		num, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return err
		}
		if num <= 0 {
			return errors.New("limit cannot be <= 0!")
		}
	}
	return nil
}

// Wraps a HTTP request context into a JSON format ready to delivery to RabbitMQ
func PrepareMessage(params map[string][]string, body []byte, operation int) ItemMsgJSON {
	var message ItemMsgJSON
	message.Params = params
	message.Body = body
	message.Operation_type = operation
	return message
}
