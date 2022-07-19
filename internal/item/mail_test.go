package main

import "testing"

func TestSendEmail(t *testing.T) {
	esItems := []ElasticItem{}
	esItems = append(esItems, ElasticItem{
		Name:         "Foobar0",
		Item_details: "A foobar. Baz right.",
		Id:           "032174vof",
	})
	esItems = append(esItems, ElasticItem{
		Name: "Foobar1",
		Id:   "0a9vsf",
	})
	lostItem := Item{
		Name: "FooBar",
	}
	if !MailSendMessage(esItems, lostItem, "findnus@outlook.com") {
		t.Fail()
		t.Log("Email sending failed!")
	}
}

// func TestPeriodicCheck(t *testing.T) {
// 	SetupMongo()
// 	SetupElasticClient()
// 	PeriodicCheck()
// }
