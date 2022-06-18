package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Recieves work to be done
func HandleRequest(d amqp.Delivery) {
	msg := UnmarshallMessage(d.Body)
	switch msg.Operation_type {
	case OPERATION_NEW_ITEM:
		DoAddItem(msg)
		break
	case OPERATION_PATCH_ITEM:
		DoUpdateItem(msg)
		break
	case OPERATION_GET_ITEM:
		item := DoGetItem(msg)
		PublishResponse(item, d)
		break
	case OPERATION_GET_ITEM_LIST:
		//foo
		fmt.Println("Get Item List Triggered")
		items := DoGetManyItems(msg)
		PublishResponse(items, d)
		break
	case OPERATION_DEL_ITEM:
		DoDeleteItem(msg)
		break
	default:
		// foo
		break
	}
}

// Item microservice entrypoint
func main() {
	SetupMongo()
	SetupMessageBrokerConnection()
	SetupChannel()
	go ConsumeMessages()
	go ConsumeGetMessages()
	forever := make(chan bool)
	<-forever
}
