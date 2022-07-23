package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Get Heroku's PORT env variable to listen for HTTP requests on
	port := os.Getenv("PORT")
	if port == "" {
		// App is running locally
		port = "8080"
	}

	router := gin.Default()
	router.Use(
		cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Headers"},
			AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
			AllowCredentials: true,
		}),
	)

	// Auth Handler
	firebaseApp := InitFirebase()

	// DEBUG ENDPOINTS
	grpDebug := router.Group("/debug")
	{
		grpDebug.GET("/ping", debugPingHandler)
		grpDebug.GET("/checkAuth", CheckAuthMiddleware(&firebaseApp), debugCheckAuth)
		grpDebug.GET("/getDemoItem", debugGetDemoItem)
	}

	// CRUD HANDLERS (Excluding search)
	grpItem := router.Group("/item")
	{
		// Creation of new items
		grpItem.POST("", HandleNewItem)
		// Update of items
		grpItem.PATCH("", HandleUpdateItem) //TODO: Add auth middleware
		// Deletion
		grpItem.DELETE("", HandleDeleteItem) //TODO: Add auth middleware
		// Get specific item
		grpItem.GET("", HandleGetOneItem) //TODO
		// Get range of items
		grpItem.GET("/peek", HandleGetManyItems) //TODO
	}

	// Elastic search for items
	router.GET("/search", HandleElasticSearch)

	// Lookout microservice endpoints
	router.GET("/lookout", HandleLookoutGet)

	// Database setup (for debugging)
	setupMongo("Items")

	// Message Broker setup for inter-microservice communication
	SetupMessageBrokerConnection()
	SetupChannelQueues()

	// Consume RPC return calls for GET messages
	go ConsumeGetItemMessage()
	go ConsumeLookoutResponseMessages()
	router.Run(":" + port)
}
