package main

import (
	"github.com/gin-gonic/gin"
)

// Pings back the requester. Used to show that backend container is alive.
func debugPingHandler(c *gin.Context) {
	// Format response
	c.JSON(200, gin.H{
		"message": "Hello! You have reached FindNUS.",
	})
}

// Completes the authcheck process
func debugCheckAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Your id is ok!",
	})
}
