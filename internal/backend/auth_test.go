/*
	---- AUTH TEST ----
	This file contains testing functions + setup logic to verify Firebase Auth setup.
*/

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Returns a userId token for a 'test' user on staging environment for debugging & testing
func getUserIdToken() string {
	// firebase-dev endpoint
	endpt := "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyATt8zAsltdjrjO2Na_IFV58cIYCL646Hs"
	loginBody, _ := json.Marshal(map[string]string{
		"email":             "automated_testuser@foobar.com",
		"password":          "test1234",
		"returnSecureToken": "true",
	})
	postBody := bytes.NewBuffer(loginBody)
	resp, _ := http.Post(
		endpt,
		"application/json",
		postBody,
	)
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	var responseMap map[string]string
	json.Unmarshal(response, &responseMap)
	// log.Println(responseMap["idToken"])
	return responseMap["idToken"]
}

func TestCheckAuthMiddleware(t *testing.T) {
	// Setup Firebase Auth connection and get a test ID
	app := InitFirebase()
	idToken := getUserIdToken()

	// Setup testing environment
	ginHandler := CheckAuthMiddleware(&app)
	httpWriter := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpWriter)
	ginContext.Request, _ = http.NewRequest("GET", "/debug/checkAuth", nil)

	// Test for valid user
	ginContext.Request.Header.Set("Authorization", idToken)
	ginHandler(ginContext)
	if httpWriter.Code == 401 {
		t.Errorf("Authentication failed for valid token")
	}
	// Test for invalid user
	ginContext.Request.Header.Set("Authorization", "foobar")
	ginHandler(ginContext)
	if httpWriter.Code != 401 {
		t.Errorf("Authentication worked for invalid token")
	}
}
