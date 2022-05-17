package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	// Safety check to prevent overzealous ping-ers
	if c.GetHeader("Cookie") != "do-not-ddos" {
		c.JSON(http.StatusForbidden, "")
		return
	}
	// Format response
	c.JSON(200, gin.H{
		"message": "Hello! You have reached FindNUS.",
	})
}

func main() {
	// Env Variables for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		// App running locally
		port = "8080"
	}

	// For now, we will create a dummy application to test docker integration
	router := gin.Default()
	router.GET("/ping", pingHandler)
	router.Run(":" + port)
}
