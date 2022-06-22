package main

import (
	"encoding/json"
	"testing"
)

func TestElasticCRUD(t *testing.T) {
	SetupElasticClient()
	ElasticInitIndex()
	testItems := loadTestItems("debug_es_item.json")
	bytes, _ := json.Marshal(testItems[0])
	var item ElasticItem
	json.Unmarshal(bytes, &item)
	ElasticAddItem(item)
	//ADD OK
	ElasticGetItem(testItems[0]["Id"].(string))
}

func TestElasticUpdateItem(t *testing.T) {

}

func TestElasticDeleteItem(t *testing.T) {
	// ctx := context.Background()
	// res, err := EsClient.Delete().
	// 	Index("twitter").
	// 	Type("tweet").
	// 	Id("1").
	// 	Do(ctx)
	// if err != nil {
	// 	// Handle error
	// 	panic(err)
	// }
	// if res.{
	// 	fmt.Print("Document deleted from from index\n")
	// }
}

func TestElasticGetItem(t *testing.T) {
	SetupElasticClient()
	ElasticInitIndex()
	testItems := loadTestItems("debug_es_item.json")
	bytes, _ := json.Marshal(testItems[0])
	var item ElasticItem
	json.Unmarshal(bytes, &item)
	//ADD OK
	ElasticGetItem(testItems[0]["Id"].(string))
}
