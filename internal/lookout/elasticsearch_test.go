package main

import "testing"

func TestElasticLookoutSearch(t *testing.T) {
	SetupElasticClient()
	ElasticLookoutSearch("Sengkang", "Electronics")
}
