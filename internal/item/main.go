package main

// Item microservice entrypoint
func main() {
	SetupMongo()
	SetupMessageBrokerConnection()
	SetupChannelQueues()
	ConsumeMessages()
}
