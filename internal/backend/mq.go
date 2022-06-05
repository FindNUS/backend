package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

// Connection -> Channel -> Queue

// Queue Name declarations
const (
	QUEUE_ITEM   string = "q_item"
	QUEUE_SEARCH string = "q_search"
)

// Global variables to allow MQ access
var MqConn *amqp.Connection
var ItemChannel *amqp.Channel
var ItemQueueConfig amqp.Queue

// var ItemJobId uint64              // Unique Job ID. Overflows OK!
type ItemJobId struct {
	mu sync.Mutex
	id uint64
}

var itemJobId ItemJobId

// Thread safe channel

// var ItemChannel string

// Thread-safe jobId implementation
func GetItemJobId() uint64 {
	// Force one goroutine to access jobId at a time
	itemJobId.mu.Lock()
	defer itemJobId.mu.Unlock()
	itemJobId.id += 1
	jobId := itemJobId.id
	return jobId
}

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
		QUEUE_ITEM, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
}

func PublishMessage(channel *amqp.Channel, jsonMsg ItemMsgJSON) {
	var err error
	bytes, err := json.Marshal(jsonMsg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	channel.Publish(
		"",                   // exchange
		ItemQueueConfig.Name, // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(bytes),
		})
}
