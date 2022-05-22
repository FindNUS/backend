package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get Heroku's PORT env variable to listen for HTTP requests on
	port := os.Getenv("PORT")
	if port == "" {
		// App running locally
		port = "8080"
	}

	router := gin.Default()

	// Auth Handler
	firebaseApp := InitFirebase()

	// DEBUG ENDPOINTS
	grpDebug := router.Group("/debug")
	{
		grpDebug.GET("/ping", debugPingHandler)
		grpDebug.GET("/checkAuth", CheckAuthMiddleware(&firebaseApp), debugCheckAuth)
	}

	// TODO: Group handlers for /item and /search endpoints in future versions

	router.Run(":" + port)
}
