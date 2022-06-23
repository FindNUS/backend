package main

import "github.com/gin-gonic/gin"

// Handler for /search endpoint
func HandleElasticSearch(c *gin.Context) {
	params := GetParams(c)
	if _, ok := params["query"]; !ok {
		c.JSON(400, gin.H{
			"message": "'query' missing from URL",
		})
		return
	}
	msg := PrepareMessage(params, nil, OPERATION_SEARCH)
	id := GetJobId()
	PublishGetItemMessage(ItemChannel, msg, id)
	res := PollResponse(id)
	if res == nil {
		c.JSON(500, gin.H{
			"message": "No response internally after 10s",
		})
		return
	}
	item := ParseGetManyItemsRPC(res)
	c.JSON(200, item)
}
