package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

// Connection -> Channel -> Queue

// Queue Name declarations
const (
	QUEUE_ITEM     string = "q_item"     // Creation, Update, Deletion
	QUEUE_SEARCH   string = "q_search"   // Elastisearch
	QUEUE_GET_REQ  string = "q_get_req"  // Request queue for Get one, get many
	QUEUE_GET_RESP string = "q_get_resp" // Response queue
)

// Global variables to allow MQ access
// TODO: Consider salting the queue names with random strings to allow unique scaling
var MqConn *amqp.Connection
var ItemChannel *amqp.Channel // Channel to hold multiple queues?
var ItemQueueConfig amqp.Queue
var GetItemQueueConfig amqp.Queue

// Common map to concurrently read RPC return calls
var GetItemReturns sync.Map

// JobID Unique Job ID. Overflows OK!
type JobId struct {
	mu sync.Mutex
	id uint64
}

var rpcJobId JobId

// Thread-safe jobId implementation
func GetJobId() uint64 {
	// Force one goroutine to access jobId at a time
	// To prevent a race condition
	rpcJobId.mu.Lock()
	defer rpcJobId.mu.Unlock()
	rpcJobId.id += 1
	jobId := rpcJobId.id
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
/* --- FINDNUS QUEUES ---
* Three queues are setup as seperate pipelines to handle
*
*
 */
func SetupChannelQueues() {
	var err error
	ItemChannel, err = MqConn.Channel()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Item Creation, Update, Deletion queue
	ItemQueueConfig, err = ItemChannel.QueueDeclare(
		QUEUE_ITEM, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	// Get Item (RPC) Queue
	ItemChannel.QueueDeclare(
		QUEUE_GET_RESP, // name TODO salt this queue and make exclusive true to enable true scalability
		false,          // durable
		false,          // delete when unused
		true,           // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	ItemChannel.QueueDeclare(
		QUEUE_GET_REQ, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func PublishMessage(channel *amqp.Channel, jsonMsg ItemMsgJSON) {
	var err error
	bytes, err := json.Marshal(jsonMsg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// Publish to a queue with a unique ID
	channel.Publish(
		"",         // exchange
		QUEUE_ITEM, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(bytes),
		})
}

// Publish a GET item/items message to item microservice...
// ...to get a single item/a list of truncated items
func PublishGetItemMessage(channel *amqp.Channel, jsonMsg ItemMsgJSON, jobId uint64) {
	var err error
	bytes, err := json.Marshal(jsonMsg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = channel.Publish(
		"",            // exchange
		QUEUE_GET_REQ, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			Body:          []byte(bytes),
			CorrelationId: strconv.FormatInt(int64(jobId), 10),
			ReplyTo:       QUEUE_GET_RESP,
		})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Always running subroutine to listen for RPC returns
func ConsumeGetItemMessage() {
	msgs, err := ItemChannel.Consume(
		QUEUE_GET_RESP, // queue
		"",             // consumer
		true,           // auto-ack
		true,           // exclusive; get resp should be salted
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Populate RPC responses to a sync map
	for msg := range msgs {
		id, _ := strconv.ParseUint(msg.CorrelationId, 10, 64)
		GetItemReturns.Store(id, msg.Body)
	}
	// Should not reach here
	log.Println("Shutting down channel")
}

// Poll response with a timeout of 10s
func PollResponse(jobId uint64) []byte {
	ok := true
	var res []byte
	var interf interface{}
	// TODO add a timeout
	timeout := time.Second * 10
	end := time.Now().Add(timeout)
	for time.Now().Before(end) {
		interf, ok = GetItemReturns.LoadAndDelete(jobId)
		if ok {
			res = interf.([]byte) // type assertion
			return res
		}
	}
	return nil
}
