package main

import (
	"log"
	"strings"

	"github.com/jdkato/prose/v2"
)

// NlpGetQuery takes in an item, does preprocessing..
// ..to get an optimiseds search query for Elasticsearching
// Returns a query string to be fed to Elasticsearch
func NlpGetQuery(item Item) string {
	// Uppercase Name to let NLP identify this as a noun
	rawString := strings.ToUpper(item.Name)
	rawString += ". " + item.Category
	rawString += ". " + item.Item_details
	rawString += ". " + strings.ToUpper(item.Location)

	doc, _ := prose.NewDocument(rawString)
	query := ""
	// Entities hold more weight than descriptive tokens, so add them in as duplicates
	ents := doc.Entities()
	for _, ent := range ents {
		query += " " + ent.Text
	}
	toks := doc.Tokens()
	for _, tok := range toks {
		// Take in all the nouns
		if tok.Tag == "NN" || tok.Tag == "NNS" || tok.Tag == "NNPS" || tok.Tag == "NNP" {
			query += " " + tok.Text
		} else if tok.Tag == "JJ" || tok.Tag == "JJR" {
			// As well as the adjectives
			query += " " + tok.Text
		}
	}
	log.Println("NLP Preprocess Result:", query)
	return query
}
