package main

import (
	"log"
	"time"
)

// Called periodically
func PeriodicCheck() {
	for true {
		lookoutRequests := MongoGetAllLookoutRequests(COLL_LOST)
		log.Println("Lookout requests:", lookoutRequests)
		for _, request := range lookoutRequests {
			query := NlpGetQuery(request)
			elasticItems := ElasticLookoutSearch(query, request.Category)
			userEmail := FirebaseGetEmailFromUser(request.User_id)
			if userEmail == "" {
				// skip invalid email
				continue
			}
			PrettyPrintStruct(elasticItems)
			// if !MailSendMessage(elasticItems, request, userEmail) {
			// 	log.Println("Error sending email.")
			// }
		}
		log.Println("Periodic Lookout Sleeping...")
		time.Sleep(time.Hour * 24) // daily reset, or whenever the container gets woken up
	}
}

// Receives a message to perform a smart lookout search on-demand for a given item
// Returns a list of found items that are good matches for the item
func LookoutDirect(msg ItemMsgJSON) []ElasticItem {
	items := []ElasticItem{}
	// Item should have the parameters, but double check for safety
	var user_id string
	var id string
	if tmp, ok := msg.Params["User_id"]; ok {
		user_id = tmp[0]
	} else {
		// No user_id, return zero-initialised items
		return items
	}
	if tmp, ok := msg.Params["Id"]; ok {
		id = tmp[0]
	} else {
		// No id, return zero-initialised items
		return items
	}
	// Item is valid after parameter check.
	// Get the item from the database
	item := MongoGetItem(COLL_LOST, id, user_id)
	if item == (Item{}) {
		// Item could not be found from MongoDB
		return items
	}
	// Start the smart search process - preprocess the item's text
	queryString := NlpGetQuery(item)
	esItems := ElasticLookoutSearch(queryString, item.Category)
	return esItems
}
