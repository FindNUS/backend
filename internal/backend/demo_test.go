package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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

	// Setup request for VALID item on database
	ginContext, _ := gin.CreateTestContext(httpWriter)
	req := &http.Request{
		URL: &url.URL{},
	}
	q := req.URL.Query()
	q.Add("name", "Laptop")
	ginContext.Request, _ = http.NewRequest("GET", "/debug/getDemoItem", nil)

	// Set this since under the hood c.BindQuery calls
	// req.URL.Query(), which calls ParseQuery(u.RawQuery)
	req.URL.RawQuery = q.Encode()

	// Test if we get intended outcomes
	ginContext.Request = req
	debugGetDemoItem(ginContext)
	if httpWriter.Code != 200 {
		t.Errorf("Could not find a valid item")
	}

	// Setup request for NON-EXISTENT item on database
	httpWriter = httptest.NewRecorder()
	ginContext, _ = gin.CreateTestContext(httpWriter)
	// Setup request
	req = &http.Request{
		URL: &url.URL{},
	}
	q = req.URL.Query()
	q.Add("name", "Foobar")
	ginContext.Request, _ = http.NewRequest("GET", "/debug/getDemoItem", nil)
	req.URL.RawQuery = q.Encode()

	// Test if we get intended outcomes
	ginContext.Request = req
	debugGetDemoItem(ginContext)
	if httpWriter.Code != 404 {
		t.Errorf("Expected 404 but got something else")
	}
}
