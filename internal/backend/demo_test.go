package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Ensure that demo works
func TestSetupMongo(t *testing.T) {
	setupMongo("Items")
}

func TestDebugGetDemoItem(t *testing.T) {
	// Setup testing environment
	setupMongo("Items")
	httpWriter := httptest.NewRecorder()

	ginContext, _ := gin.CreateTestContext(httpWriter)
	ginContext.Request, _ = http.NewRequest("GET", "/debug/getDemoItem", nil)
	jsonbody := `{ "name":"Laptop" }`
	req := ioutil.NopCloser(bytes.NewReader([]byte(jsonbody)))
	// Test if we get intended outcomes
	ginContext.Request.Body = req
	debugGetDemoItem(ginContext)

	if httpWriter.Code != 200 {
		t.Errorf("Could not find a valid item")
	}

	httpWriter = httptest.NewRecorder()
	ginContext, _ = gin.CreateTestContext(httpWriter)
	ginContext.Request, _ = http.NewRequest("GET", "/debug/getDemoItem", nil)
	jsonbody = `{ "name":"Foobar" }`
	req = ioutil.NopCloser(bytes.NewReader([]byte(jsonbody)))
	// Test if we get intended outcomes
	ginContext.Request.Body = req
	debugGetDemoItem(ginContext)

	if httpWriter.Code != 404 {
		t.Errorf("Expected 404 but got something else")
	}
}
