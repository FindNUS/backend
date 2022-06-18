package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// Global variables to allow MQ access
var MqConn *amqp.Connection
var ItemChannel *amqp.Channel
var ItemQueue amqp.Queue
var GetItemQueue amqp.Queue

// Thread-safe jobId implementation

// Setup RabbitMQ service. Remember to defer connection close.
func SetupMessageBrokerConnection() {
	uri := os.Getenv("RABBITMQ_URI")
	if uri == "" {
		// Try searching for URI locally
		f, err := os.Open("../../secrets/rabbitmqDev.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(f)
		defer f.Close()
		for scanner.Scan() {
			uri = scanner.Text()
		}
		if uri == "" {
			log.Fatalln("RABBITMQ URI NOT FOUND")
		}
	}
	var err error
	MqConn, err = amqp.Dial(uri)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// Setup Channels
func SetupChannel() {
	var err error
	ItemChannel, err = MqConn.Channel()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Subroutine to handle message consumption
func ConsumeMessages() {
	// Consume Item (POST PATCH DELETE) messages
	msgs, err := ItemChannel.Consume(
		"q_item", // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		HandleRequest(d)
	}
	log.Println("Consume Item messages shutting down")
}

// Subroutine to handle GET messages
func ConsumeGetMessages() {
	// Consume Item (POST PATCH DELETE) messages
	msgs, err := ItemChannel.Consume(
		"q_get_req", // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	ItemChannel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	for d := range msgs {
		log.Printf("Received a GET message: %s", d.Body)
		HandleRequest(d)
	}
	log.Println("Consume Get Messages shutting down")
}

func PublishResponse(item interface{}, d amqp.Delivery) {
	// Marshall response
	msg, _ := json.Marshal(item)
	ItemChannel.Publish(
		"",        // exchange
		d.ReplyTo, // routing key (name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: d.CorrelationId,
			Body:          []byte(msg),
		},
	)
}
