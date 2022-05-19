package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Pings back the requester. Used to show that backend container is alive.
func debugPingHandler(c *gin.Context) {
	// Format response
	c.JSON(200, gin.H{
		"message": "Hello! You have reached FindNUS.",
	})
}

func main() {
	// Get Heroku's PORT env variable to listen for HTTP requests on
	port := os.Getenv("PORT")
	if port == "" {
		// App running locally
		port = "8080"
	}

	// For now, we will create a dummy application to test docker integration
	router := gin.Default()

	// DEBUG ENDPOINTS
	grpDebug := router.Group("/debug")
	{
		grpDebug.GET("/ping", debugPingHandler)
	}

	router.Run(":" + port)
}
