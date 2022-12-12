package main

import (
	"github.com/gin-gonic/gin"
)

// Pings back the requester. Used to show that backend container is alive.
// Also tests MQ and database functions to 'keep-alive' the connections
func debugPingHandler(c *gin.Context) {
	// Format response
	c.JSON(200, gin.H{
		"message": "Hello! You have reached FindNUS.",
	})
	// Additional wake-up functionality
	debugGetDemoItem()
}

// User passed the AuthGuard. User is logged in, hence authenticated to do priviledged operations.
func debugCheckAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Your id is OK!",
	})
}
