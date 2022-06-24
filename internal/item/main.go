package main

import (
	"log"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Recieves work from MessageQueue to execute
func HandleRequest(d amqp.Delivery) {
	msg := UnmarshallMessage(d.Body)
	switch msg.Operation_type {
	case OPERATION_NEW_ITEM:
		collection, objId := DoAddItem(msg)
		if objId == primitive.NilObjectID {
			log.Println("Error parsing DoAddItem - Aborting Es Add")
			return
		}
		if collection == COLL_FOUND {
			item := MongoGetItem(COLL_FOUND, objId.Hex(), "")
			esItem := ElasticParseBody(item, objId)
			if esItem != (ElasticItem{}) {
				ElasticAddItem(esItem)
			}
		}
		break

	case OPERATION_PATCH_ITEM:
		id, err := DoUpdateItem(msg)
		if err != nil {
			log.Println("Error patching item on MongoDB:", err.Error())
			return
		}
		// Check if patched item should be registered on Es
		if id != "" {
			objId, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				log.Println("Error patching item to EsCloud:", err.Error())
				return
			}
			item := MongoGetItem(COLL_FOUND, objId.Hex(), "")
			// Safety check for FOUND collection
			if item != (Item{}) {
				esItem := ElasticParseBody(item, objId)
				ElasticUpdateItem(esItem)
			}
		}
		break

	case OPERATION_GET_ITEM:
		item := DoGetItem(msg)
		PublishResponse(item, d)
		break

	case OPERATION_GET_ITEM_LIST:
		items := DoGetManyItems(msg)
		PublishResponse(items, d)
		break

	case OPERATION_DEL_ITEM:
		id, err := DoDeleteItem(msg)
		if err != nil {
			log.Println("Error deleting item on MongoDB:", err.Error())
			return
		}
		if id != "" {
			ElasticDeleteItem(id)
		}
		break

	case OPERATION_SEARCH:
		qry := GetElasticQuery(msg.Params)
		// Safety catch to prevent ALL the items from being returned
		if qry == "" {
			foo := []ElasticItem{}
			PublishResponse(foo, d)
			return
		}
		res := ElasticSearchGeneral(qry)
		PublishResponse(res, d)

	default:
		// Should not reach here -- do nothing
		break
	}
}

// Item microservice entrypoint
func main() {
	SetupImgur()
	SetupElasticClient()
	ElasticInitIndex()
	SetupMongo()
	SetupMessageBrokerConnection()
	SetupChannel()
	go ConsumeMessages()
	go ConsumeGetMessages()
	forever := make(chan bool)
	<-forever
}
