package main

import "testing"

func TestMongoGetAllLookoutRequests(t *testing.T) {
	SetupMongo()
	items := MongoGetAllLookoutRequests(COLL_DEBUG)
	for _, item := range items {
		PrettyPrintStruct(item)
	}
}
