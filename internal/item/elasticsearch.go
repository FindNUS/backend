package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/olivere/elastic/v7"
)

// This file aims to handle CRUD of items to the Bonsai Elasticsearch cloud
var EsClient *elastic.Client

func SetupElasticClient() {
	bonsaiURI := os.Getenv("BONSAI_ES_URI")
	var err error
	if bonsaiURI == "" {
		// Read from secrets file
		f, err := os.Open("../../secrets/bonsaiEs.txt")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(f)
		defer f.Close()
		for scanner.Scan() {
			bonsaiURI = scanner.Text()
		}
	}
	EsClient, err = elastic.NewClient(
		elastic.SetURL(bonsaiURI),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Ping to test
	info, code, err := EsClient.Ping("https://findnus-prod-8254101466.eu-west-1.bonsaisearch.net:443").Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(info)
	fmt.Println(code)
	fmt.Println("Done")
	// Create Index

}

//
func ElasticInitIndex() {
	if EsClient == nil {
		log.Println("EsClient not intialised")
		return
	}
	// FindNUS ElasticIndex schema
	esMap := `{
		"settings": {
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings": {
			"properties": {
				"collection": {
					"type":"keyword"
				},
				"Id": {
					"type":"keyword"
				},
				"Name": {
					"type":"text"	
				},
				"Category": {
					"type":"text"	
				},
				"Item_details": {
					"type":"text"	
				},
				"Location": {
					"type":"text"	
				}
			}
		}
	}`
	ctx := context.Background()
	// Check if the index exists
	exists, err := EsClient.IndexExists("found_items").Do(context.Background())
	if err != nil {
		// Handle error
		log.Fatal(err.Error())
	}
	if !exists {
		// Index does not exist yet.
		createIndex, err := EsClient.CreateIndex("found_items").BodyString(esMap).Do(ctx)
		if err != nil {
			// Handle error
			log.Fatal(err.Error())
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			log.Fatal("NACK")
		}
	}
	log.Println("found_items index exists and has been initialised!")
}

func ElasticAddItem(item ElasticItem) {
	res, err := EsClient.Index().Index("found_items").BodyJson(item).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}

func ElasticUpdateItem() {

}

func ElasticDeleteItem(id string) {
	ctx := context.Background()
	qry := elastic.NewTermQuery("Id", id)
	res, err := EsClient.Delete(qry).Index("found_items").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if res.Found {
		fmt.Print("Document deleted from from index\n")
	}
}

// Get item by id (not search)
func ElasticGetItem(id string) {
	qry := elastic.NewTermQuery("Id", id)
	res, err := EsClient.Search().Index("found_items").Query(qry).Pretty(true).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var esItem ElasticItem
	for _, item := range res.Each(reflect.TypeOf(esItem)) {
		PrettyPrintStruct(item.(ElasticItem))
	}
}
