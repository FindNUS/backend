package main

import (
	"net/http"
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

// User passed the AuthGuard. User is logged in, hence authenticated to do priviledged operations.
func debugCheckAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Your id is OK!",
	})
}

// Ensure that the item microservice is alive and well
func keepItemAlive() error {
	prodVar, _ := os.LookupEnv("PRODUCTION")
	var err error
	if prodVar == "true" {
		_, err = http.Get("https://findnus-prod-item.onrender.com/ping")
	} else {
		// prodVar == false
		_, err = http.Get("https://findnus-uat-item.onrender.com/ping")
	}
	return err
}
