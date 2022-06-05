package main

func main() {
	SetupMongo()
	SetupMessageBrokerConnection()
	SetupChannelQueues()
	ConsumeMessages()
}
