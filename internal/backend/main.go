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
			AllowHeaders:     []string{"Origin", "Authorization"},
			AllowMethods:     []string{"GET", "POST", "PATCH", "PUT"},
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
		new := grpItem.Group("/new")
		{
			new.POST("/lost", HandleNewLostItem)
			new.POST("/found", HandleNewFoundItem)
		}
		// Update of items
		grpItem.PATCH("/update", HandleUpdateItem) //TODO: Add auth middleware
		// Deletion
		grpItem.DELETE("/delete", HandleDeleteItem) //TODO: Add auth middleware
		// Get specific item
		grpItem.GET("/get") //TODO
		// Get range of items
		grpItem.GET("/peek") //TODO
	}
	setupMongo("Items")
	SetupMessageBrokerConnection()
	SetupChannelQueues()
	router.Run(":" + port)
}
