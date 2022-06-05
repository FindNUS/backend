package main

import "encoding/json"

// Marshaller-Unmarshaller

// Unmarshall message from Message Broker
func UnmarshallMessage(bytes []byte) ItemMsgJSON {
	var msg ItemMsgJSON
	json.Unmarshal(bytes, &msg)
	return msg
}
