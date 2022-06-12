package main

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleNewFoundItem(c *gin.Context) {
	// No params check needed
	// TODO safety check for required fields
	params := GetParams(c)
	log.Println("Params: ")
	PrettyPrintStruct(params)
	rawBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Could not read body: " + err.Error(),
		})
		return
	}
	body, err := ParseFoundItemBody(rawBody)
	if body == nil {
		c.JSON(400, gin.H{
			"message": "Form body has issues: " + err.Error(),
		})
		return
	}
	msg := PrepareMessage(params, body, OPERATION_NEW_ITEM)
	PublishMessage(ItemChannel, msg)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func HandleNewLostItem(c *gin.Context) {
	// No params check needed
	params := GetParams(c)
	log.Println("Params: ")
	PrettyPrintStruct(params)
	rawBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Could not read body: " + err.Error(),
		})
		return
	}
	body, err := ParseLostItemBody(rawBody)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Form body has issues: " + err.Error(),
		})
		return
	}
	msg := PrepareMessage(params, body, OPERATION_NEW_ITEM)
	PublishMessage(ItemChannel, msg)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func HandleDeleteItem(c *gin.Context) {
	// Process & validate parameters
	params := GetParams(c)
	keys := []string{
		"Id",
		"User_id",
	}
	paramlen := len(params)
	if paramlen < 1 || params["Id"] == nil {
		c.JSON(400, gin.H{
			"message": "Missing parameters! Ensure that there is at least an Id parameter",
		})
		return
	}
	// Check if at least 1 valid key exists
	for _, key := range keys {
		if _, ok := params[key]; ok {
			break
		}
	}
	// Create dummy body and wrap message
	var body []byte
	msg := PrepareMessage(params, body, OPERATION_DEL_ITEM)
	PublishMessage(ItemChannel, msg)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func HandleUpdateItem(c *gin.Context) {
	params := GetParams(c)
	keys := []string{
		"Id",
		"User_id",
	}
	paramlen := len(params)
	if paramlen < 1 || params["Id"] == nil {
		c.JSON(400, gin.H{
			"message": "Missing parameters! Ensure that there is at least an Id parameter",
		})
		return
	}
	// Check if at least 1 valid key exists
	for _, key := range keys {
		if _, ok := params[key]; ok {
			break
		}
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	msg := PrepareMessage(params, body, OPERATION_PATCH_ITEM)
	PublishMessage(ItemChannel, msg)
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
