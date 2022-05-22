package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Factored out the main Routing functions to allow for better testing
func main() {
	// Get Heroku's PORT env variable to listen for HTTP requests on
	port := os.Getenv("PORT")
	if port == "" {
		// App running locally
		port = "8080"
	}

	// For now, we will create a dummy application to test docker integration
	router := gin.Default()

	// Auth Handler
	firebaseApp := InitFirebase()

	// DEBUG ENDPOINTS
	grpDebug := router.Group("/debug")
	{
		grpDebug.GET("/ping", debugPingHandler)
		grpDebug.GET("/checkAuth", CheckAuthMiddleware(&firebaseApp), debugCheckAuth)
	}

	router.Run(":" + port)
}
