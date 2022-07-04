package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Test that overall HTTP handler and utils body parser work
func TestHandleNewItem(t *testing.T) {
	// SetupDebugQueues()
	// Test invalid items
	httpwriter := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpwriter)
	bodies := loadTestItems("invalid_lost_items.json")
	bodies = append(bodies, loadTestItems("invalid_found_items.json")...)
	for _, body := range bodies {
		log.Println("OK")
		bodybytes, _ := json.Marshal(body)
		ginContext.Request, _ = http.NewRequest("POST", "", bytes.NewBuffer(bodybytes))
		HandleNewItem(ginContext)
		if httpwriter.Code != 400 {
			t.Fail()
			t.Log("HandleNewItem failed -- Expected 400 but got 200.\nFailed item:", body)
		}

	}
	// Test valid items
	httpwriter = httptest.NewRecorder()
	ginContext, _ = gin.CreateTestContext(httpwriter)
	bodies = loadTestItems("valid_lost_items.json")
	bodies = append(bodies, loadTestItems("valid_found_items.json")...)
	for _, body := range bodies {
		bodybytes, _ := json.Marshal(body)
		ginContext.Request, _ = http.NewRequest("POST", "", bytes.NewBuffer(bodybytes))
		HandleNewItem(ginContext)
		if httpwriter.Code != 200 {
			t.Fail()
			t.Log("HandleNewItem failed -- Expected 200 but got !200.\nFailed item:", body)
		}
	}
}
