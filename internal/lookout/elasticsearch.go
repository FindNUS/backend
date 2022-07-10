package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NOTE: EsClient is thread safe -- driver implements muSync
var EsClient *elastic.Client
var IndexName string

func debugTestQuery(query elastic.Query) {
	src, err := query.Source()
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

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
	// Init index name
	IndexName = "found_items_uat"
	if prodstr, ok := os.LookupEnv("PRODUCTION"); ok {
		if prodstr == "true" {
			IndexName = "found_items"
		}
	}
	log.Println("IndexName for EsClient is:", IndexName)
	fmt.Println("ElasticSearch setup done")
}

// Startup the ElasticSearch index if needed.
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
				},
				"Image_url": {
					"type":"text"
				}
			}
		}
	}`
	ctx := context.Background()
	// Check if the index exists
	exists, err := EsClient.IndexExists(IndexName).Do(context.Background())
	if err != nil {
		// Handle error
		log.Fatal(err.Error())
	}
	if !exists {
		// Index does not exist yet.
		createIndex, err := EsClient.CreateIndex(IndexName).BodyString(esMap).Do(ctx)
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

func ElasticLookoutSearch(qry string, cat string) []ElasticItem {
	query := elastic.NewBoolQuery()
	mmq := elastic.NewMultiMatchQuery(
		qry,
		"Name", "Location", "Item_details",
	)
	mmq.Type("most_fields")
	mmq.Fuzziness("2")
	mmq.MinimumShouldMatch("5") // 5 clauses
	// mmq.Analyzer("standard")
	mmq.FieldWithBoost("Name", 2)
	mmq.FieldWithBoost("Item_details", 2)
	query.Must(mmq)
	query.Filter(elastic.NewTermQuery("Category", []string{"Etc", cat}))

	// Execute the search
	ctx := context.TODO()
	EsClient.Refresh().Do(context.Background())
	res, err := EsClient.Search().Index(IndexName).Query(mmq).Do(ctx)
	if err != nil {
		panic(err.Error())
	}
	esItemList := []ElasticItem{}
	esItem := ElasticItem{}
	for _, item := range res.Each(reflect.TypeOf(esItem)) {
		log.Println("Search found: ", item)
		esItem = item.(ElasticItem)
		esItemList = append(esItemList, esItem)
	}
	return esItemList
}

// Takes in raw bytes and Id argument to create an elastic item
func ElasticParseBody(item Item, id primitive.ObjectID) ElasticItem {
	esItem := ElasticItem{}
	raw, _ := json.Marshal(item)
	json.Unmarshal(raw, &esItem)
	esItem.Id = id.Hex()
	return esItem
}
