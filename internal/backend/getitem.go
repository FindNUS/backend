package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

/* ---- ITEM RETRIEVAL ---- */
func HandleGetOneItem(c *gin.Context) {
	params := GetParams(c)
	// Ensure that required parameters exist
	if _, ok := params["Id"]; !ok {
		c.JSON(400, gin.H{
			"message": "Missing Id parameter",
		})
		return
	}
	msg := PrepareMessage(params, nil, OPERATION_GET_ITEM)
	id := GetJobId()
	PublishGetItemMessage(ItemChannel, msg, id)
	res := PollResponse(id)
	if res == nil {
		c.JSON(500, gin.H{
			"message": "No response internally after 10s",
		})
		return
	}
	var generic map[string]interface{}
	json.Unmarshal(res, &generic)
	item := ParseGetOneItemRPC(generic)
	c.JSON(200, item)
}

// Filter peek
func HandleGetManyItems(c *gin.Context) {
	params := GetParams(c)
	msg := PrepareMessage(params, nil, OPERATION_GET_ITEM_LIST)
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

func ParseGetOneItemRPC(tmp map[string]interface{}) Item {
	// var tmp map[string]interface{}
	var res Item
	raw, _ := json.Marshal(tmp)
	json.Unmarshal(raw, &res)
	return res
}

func ParseGetManyItemsRPC(data []byte) []Item {
	var tmp []map[string]interface{} // element of unmarshalled items
	var res []Item
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, elem := range tmp {
		item := ParseGetOneItemRPC(elem)
		res = append(res, item)
	}
	return res
}
