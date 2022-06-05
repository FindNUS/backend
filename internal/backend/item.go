package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HandleNewFoundItem(c *gin.Context) {
	// No params check needed
	params := GetParams(c)
	// TODO safety check for required fields
	rawBody, _ := ioutil.ReadAll(c.Request.Body)
	body := ParseFoundItemBody(rawBody)
	if body == nil {
		c.JSON(400, gin.H{
			"message": "Form body has issues",
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
	rawBody, _ := ioutil.ReadAll(c.Request.Body)
	body := ParseLostItemBody(rawBody)
	if body == nil {
		c.JSON(400, gin.H{
			"message": "Form body has issues. Did you include in a User_id?",
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
			"message": "Missing parameters!",
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
			"message": "Missing parameters!",
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
