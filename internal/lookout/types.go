package main

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ---- COLLECTIONS EXPLANATIONS ----
// All collections are stored under the ITEMS database
// 1. LOST items (lookout requests) are stored in the LOST collection
// 2. FOUND items are stored in the FOUND collection
// 3. DEBUG items are meant for testings and debugs

// Type definitions for the marshalling of data
type Item struct {
	Id              primitive.ObjectID `bson:"_id"`
	Name            string
	Date            time.Time
	Location        string
	Category        string
	Contact_method  string
	Contact_details string
	Item_details    string
	Image_url       string
	User_id         string `bson:"User_id,omitempty"`
	Lookout         bool   `bson:"Lookout,omitempty"`
}

// NOTE: New Item will require some preprocessing, namely the storage of imgr
type NewItem struct {
	Name            string    `bson:"Name"`
	Date            time.Time `bson:"Date"`
	Location        string    `bson:"Location"`
	Category        int       `bson:"Category"`
	Contact_method  int       `bson:"Contact_method,omitempty"`
	Contact_details string    `bson:"Contact_details,omitempty"`
	Item_details    string    `bson:"Item_details,omitempty"`
	Image_url       string    `bson:"Image_url,omitempty"`
	Image_base64    byte      `bson:"Image_base64,omitempty"`
	User_id         string    `bson:"User_id,omitempty"`
	Lookout         bool      `bson:"Lookout, omitempty" json:"Lookout,omitempty"`
}

type PatchItem struct {
	Id              primitive.ObjectID `bson:"_id" json:"Id"`
	Name            string             `bson:"Name,omitempty" json:"Name,omitempty"`
	Date            time.Time          `bson:"Date,omitempty" json:"Date,omitempty"`
	Location        string             `bson:"Location,omitempty" json:"Location,omitempty"`
	Category        int                `bson:"Category,omitempty" json:"Category,omitempty"`
	Contact_method  int                `bson:"Contact_method,omitempty" json:"Contact_method,omitempty"`
	Contact_details string             `bson:"Contact_details,omitempty" json:"Contact_details,omitempty"`
	Item_details    string             `bson:"Item_details,omitempty" json:"Item_details,omitempty"`
	Image_url       string             `bson:"Image_url,omitempty" json:"Image_url,omitempty"`
	Image_base64    string             `bson:"-" json:"Image_base64,omitempty"`
	User_id         string             `bson:"User_id,omitempty" json:"User_id,omitempty"`
	Lookout         bool               `bson:"Lookout,omitempty" json:"Lookout,omitempty"`
}

type DeletedItem struct {
	Id      primitive.ObjectID `bson:"_id"`
	User_id string             `bson:"User_id,omitempty"`
}

type SingleItem struct {
	Id      primitive.ObjectID `bson:"_id"`
	User_id string             `bson:"User_id,omitempty"`
}

type ElasticItem struct {
	Id           string    `json:"Id"`
	Name         string    `json:"Name"`
	Location     string    `json:"Location"`
	Category     string    `json:"Category"`
	Item_details string    `json:"Item_details"`
	Image_url    string    `json:"Image_url"`
	Date         time.Time `json:"Date"`
}

// JSON Message Wrapper
type ItemMsgJSON struct {
	Operation_type int
	Params         map[string][]string
	Body           []byte
}

// Operation enum to help with routing messages to correct service
const (
	OPERATION_DEBUG    int = 0 // /debug/pingItem
	OPERATION_NEW_ITEM int = 1 // /item/new*
	// OPERATION_NEW_LOST_ITEM  int = 1
	// OPERATION_NEW_FOUND_ITEM int = 2
	OPERATION_GET_ITEM         int = 3 // /item
	OPERATION_GET_ITEM_LIST    int = 4 // /item/peek
	OPERATION_PATCH_ITEM       int = 5 // /item/update
	OPERATION_DEL_ITEM         int = 6 // /item/delete
	OPERATION_SEARCH           int = 7 // /search
	OPERATION_LOOKOUT_EXPLICIT int = 8 // /lookout
	OPERATION_LOOKOUT_CRON     int = 9 // cron scheduler microservice
)

// CATEGORY MAPPING str -> int
func GetCategoryType(cat string) int {
	cat = strings.ToLower(cat)
	switch cat {
	case "cards":
		return 1
	case "notes":
		return 2
	case "electronics":
		return 3
	case "bottles":
		return 4
	case "etc":
		return 5
	default:
		return -1
	}
}

// Category mapping int -> str
func GetCategoryString(cat int32) string {
	// WARNING: Floating point errors probable
	switch cat {
	case 0:
		return "Etc" //legacy issue
	case 1:
		return "Cards"
	case 2:
		return "Notes"
	case 3:
		return "Electronics"
	case 4:
		return "Bottles"
	case 5:
		return "Etc"
	default:
		return "Unknown"
	}
}

// CONTACT_METHOD MAPPING
func GetContactMethod(method string) int {
	method = strings.ToLower(method)
	switch method {
	case "telegram":
		return 1
	case "whatsapp":
		return 2
	case "wechat":
		return 3
	case "line":
		return 4
	case "phone_number":
		return 5
	case "nus_security":
		return 6
	default:
		return -1
	}
}

func GetContactString(cat int32) string {
	// WARNING: Floating point errors probable
	switch cat {
	case 0:
		return "Nus_security" //legacy issue
	case 1:
		return "Telegram"
	case 2:
		return "Whatsapp"
	case 3:
		return "Wechat"
	case 4:
		return "Line"
	case 5:
		return "Phone_number"
	case 6:
		return "Nus_security"
	default:
		return "Unknown"
	}
}
func ParseDateString(datestring string) time.Time {
	layout := "2006-01-02T15:04:05Z"
	if res, err := time.Parse(layout, datestring); err == nil {
		return res
	}
	return time.Now()
}
