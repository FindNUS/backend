package main

import "encoding/json"

// Unmarshall message from Message Broker
func UnmarshallMessage(bytes []byte) ItemMsgJSON {
	var msg ItemMsgJSON
	json.Unmarshal(bytes, &msg)
	return msg
}
