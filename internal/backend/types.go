package main

import (
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
	Category        int
	Contact_method  int
	Contact_details string
	Item_details    string
	Image_url       string
}

// NOTE: New Item will require some preprocessing, namely the storage of imgr
type NewItem struct {
	Name            string
	Date            time.Time
	Location        string
	Category        int
	Contact_method  int    `bson:"contact_method,omitempty"`
	Contact_details string `bson:"contact_details,omitempty"`
	Item_details    string `bson:"item_details,omitempty"`
	Image_url       string `bson:"image_url,omitempty"`
	Image_base64    byte   `bson:"-"` // Ignore this field
}

// CATEGORY MAPPING
func GetCategoryType(cat string) int {
	switch cat {
	case "Cards":
		return 1
	case "Notes":
		return 2
	case "Electronics":
		return 3
	case "Bottles":
		return 4
	default:
		return 0
	}
}

func ParseDateString(datestring string) time.Time {
	layout := "2006-01-02T15:04:05Z"
	if res, err := time.Parse(layout, datestring); err == nil {
		return res
	}
	return time.Now()
}
