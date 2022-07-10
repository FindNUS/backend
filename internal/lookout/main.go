package main

import "github.com/streadway/amqp"

/*
	---- LOOKOUT MICROSERVICE ----
	This microservice  helps Losters automatically search the database for potential matches and email them.
	1. Queries for lookout=true Lost items
	2. Pre-processes the Lost Items with a NLP library
	3. ElasticSearch the items - Only accept item scores that match well
	4. If good candidates are found, emails the user for potential matches
	5. For every email ping, we decrement the item's ping number on Mongo
*/

func HandleRequest(d amqp.Delivery) {
	msg := UnmarshallMessage(d.Body)
	switch msg.Operation_type {
	case OPERATION_LOOKOUT_CRON:
		PeriodicCheck()
		break
	case OPERATION_LOOKOUT_EXPLICIT:
		items := LookoutDirect(msg)
		PublishResponse(items, d)
		break
	default:
		if d.ReplyTo != "" {
			PublishResponse(nil, d)
		}
	}
}

func main() {
	SetupMongo()
	SetupElasticClient()
	SetupMessageBrokerConnection()
	SetupChannel()
	go ConsumeLookoutMessages()
	forever := make(chan bool)
	<-forever
}
