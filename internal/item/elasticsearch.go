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
				},
				"Pluscode": {
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

// Handler for Adding Item
func ElasticAddItem(item ElasticItem) {
	// Check for item existence first as a safety catch to avoid redundant (the bad kind) copies
	if ElasticGetItem(item.Id) != (ElasticItem{}) {
		// Item already exists! Delete it and re-add in.
		// Deletion is done due to paywalled API and wonky driver implementation...
		// ...as explained in the Update subroutines
		log.Println("Deleting")
		ElasticDeleteItem(item.Id)
	}
	res, err := EsClient.Index().Index(IndexName).BodyJson(item).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Add item response:", res)
	// PrettyPrintStruct(res)
}

// Update an existing item on elasticsearch
func ElasticUpdateItem(item ElasticItem) {
	// The most direct way to update is via the UpdateByQuery API:
	// res, err := EsClient.UpdateByQuery().Query(qry).Index(IndexName).Do(ctx)
	// HOWEVER - this function is paywalled by our ElasticSearch provider. :(

	// Other ways of updating require hacky, non-trivial implementations.
	// For sake of simplicity, naiively delete and re-add the item.
	// This is logically equivalent to what is written in ElasticAddItem().
	// Nevertheless, this handler will stay as good(?) SWE practice.
	// We may change provider or find a better way to implement update.
	// Having a seperate handler for update keeps things decoupled and easier to debug.
	ElasticDeleteItem(item.Id)
	EsClient.Refresh().Do(context.Background())
	ElasticAddItem(item)
}

func ElasticDeleteItem(id string) int64 {
	ctx := context.Background()
	qry := elastic.NewTermQuery("Id", id)
	res, err := EsClient.DeleteByQuery().Query(qry).Index(IndexName).Do(ctx)
	if err != nil {
		// Handle error
		log.Println("Error deleting item in ElasticDeleteItem(),", err.Error())
	}
	return res.Deleted
}

// Get item by id (not search)
func ElasticGetItem(id string) ElasticItem {
	qry := elastic.NewTermQuery("Id", id)
	EsClient.Refresh().Do(context.Background())
	res, err := EsClient.Search().Index(IndexName).Query(qry).Pretty(true).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var esItem ElasticItem
	for _, item := range res.Each(reflect.TypeOf(esItem)) {
		PrettyPrintStruct(item.(ElasticItem))
		esItem = item.(ElasticItem)
	}
	return esItem
}

func ElasticSearchGeneral(qry string) []ElasticItem {
	mmq := elastic.NewMultiMatchQuery(
		qry,
		"Name", "Location", "Item_details", "Category", "Id",
	)
	// Search query tuning
	mmq.Type("most_fields")
	mmq.Operator("or")
	mmq.Fuzziness("2")
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

/* LOOKOUT MICROSERVICE */
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

func ElasticLookoutSearch(qry string, cat string) []ElasticItem {
	// query := elastic.NewBoolQuery()
	mmq := elastic.NewMultiMatchQuery(
		qry,
		"Name", "Location", "Item_details", "Category",
	)
	mmq.Type("combined_fields")
	// mmq.Fuzziness("2")
	// Min match 2 clauses, 1 for category, 2 for others
	mmq.MinimumShouldMatch("3")
	// mmq.Analyzer("standard")
	mmq.FieldWithBoost("Name", 3)
	mmq.FieldWithBoost("Item_details", 2)
	// query.Must(mmq)
	// query.Filter(elastic.NewTermQuery("Category", []string{"Etc", cat}))

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
