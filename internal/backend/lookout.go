package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Validate the URL parameters that request for a On-Demand lookout
func ValidateLookoutParams(params map[string][]string) error {
	ok := true
	// Check that User_id is legit
	if _, ok = params["User_id"]; !ok {
		return errors.New("Missing User_id in URL query")
	}
	if _, ok = params["Id"]; !ok {
		return errors.New("Missing Item Id in URL query")
	}
	return nil
}

// Handler to parse on-demand lookout requests
func HandleLookoutGet(c *gin.Context) {
	err := keepItemAlive()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
	// Parse the query params
	params := GetParams(c)
	err = ValidateLookoutParams(params)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Send the query to Lookout microservice
	msg := PrepareMessage(params, nil, OPERATION_LOOKOUT_EXPLICIT)
	id := GetJobId()
	PublishGetLookoutMessage(ItemChannel, msg, id)
	res := PollResponse(id)
	if res == nil {
		c.JSON(500, gin.H{
			"message": "No response internally after 10s",
		})
	}
	items := ParseGetManyItemsRPC(res)
	c.JSON(200, items)
}
