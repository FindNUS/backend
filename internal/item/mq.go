package main

import (
	"bufio"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// Global variables to allow MQ access
var MqConn *amqp.Connection
var ItemChannel *amqp.Channel
var ItemQueueConfig amqp.Queue

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

// Setup Queue reference
func SetupChannelQueues() {
	var err error
	ItemChannel, err = MqConn.Channel()
	if err != nil {
		log.Fatalln(err.Error())
	}
	ItemQueueConfig, err = ItemChannel.QueueDeclare(
		"q_item", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
}

// Subroutine to handle message consumption
func ConsumeMessages() {
	msgs, err := ItemChannel.Consume(
		ItemQueueConfig.Name, // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			HandleRequest(d.Body)
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-forever
}
