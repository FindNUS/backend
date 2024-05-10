package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func loadTestItems(filename string) []map[string]interface{} {
	var f *os.File
	var err error
	f, err = os.Open("./test/" + filename)
	if err != nil {
		log.Fatalf(err.Error())
	}
	data, err := ioutil.ReadAll(f)
	var res []map[string]interface{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return res
}

func buildItemMsgJson(params map[string][]string, body []byte) ItemMsgJSON {
	var msg ItemMsgJSON
	if params != nil {
		msg.Params = params
	}
	if body != nil {
		msg.Body = body
	}
	return msg
}
