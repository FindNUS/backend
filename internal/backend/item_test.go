package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestHandleNewLostItem(t *testing.T) {
	// Test that user_id guard works
	httpwriter := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpwriter)
	body := map[string]interface{}{
		"Name":     "Laptop",
		"Date":     time.Now(),
		"Location": "Unknown",
		"Category": "Cards",
	}
	bodybytes, _ := json.Marshal(body)
	ginContext.Request, _ = http.NewRequest("POST", "", bytes.NewBuffer(bodybytes))

	HandleNewLostItem(ginContext, nil)
	if httpwriter.Code != 400 {
		t.Fatalf("Wrong code")
	}
	// Test type-senstive Category guard
	httpwriter = httptest.NewRecorder()
	ginContext, _ = gin.CreateTestContext(httpwriter)
	body = map[string]interface{}{
		"Name":     "Laptop",
		"Date":     time.Now(),
		"Location": "Unknown",
		"Category": 77,
		"User_id":  "7j0fs",
	}
	bodybytes, _ = json.Marshal(body)
	ginContext.Request, _ = http.NewRequest("POST", "", bytes.NewBuffer(bodybytes))
	HandleNewLostItem(ginContext, nil)
	if httpwriter.Code != 400 {
		t.Fatalf("Wrong code - Assertion type")
	}
}

func TestHandleNewFoundItem(t *testing.T) {
	// Test type-senstive Category guard
	httpwriter := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpwriter)
	body := map[string]interface{}{
		"Name":     "Laptop",
		"Date":     time.Now(),
		"Location": "Unknown",
		"Category": 77,
		"User_id":  "7j0fs",
	}
	bodybytes, _ := json.Marshal(body)
	ginContext.Request, _ = http.NewRequest("POST", "", bytes.NewBuffer(bodybytes))
	HandleNewFoundItem(ginContext, nil)
	if httpwriter.Code != 400 {
		t.Fatalf("Wrong code - Assertion type")
	}
}
