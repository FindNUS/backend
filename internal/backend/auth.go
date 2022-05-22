/*
	This file describles logic to take care of firebase authentication
*/
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// Programatically recreates the google service account JSON
// Do this because Github and Heroku do not have a native way of storing JSON secrets
func GetGoogleCredJson(isProd bool) map[string]string {
	var projectId string
	var private_key_id string
	private_key := "-----BEGIN PRIVATE KEY-----\n"
	tmp, _ := os.LookupEnv("FIREBASE_KEY")
	tmp = strings.ReplaceAll(tmp, "\\n", "")
	// log.Println(tmp[0:100])
	// log.Println(tmp[100:])
	private_key += tmp + "\n-----END PRIVATE KEY-----\n"
	log.Println("Key details: ", len(private_key), ", ", string(private_key[0:100]), string(private_key[100:]))
	strings.Trim(private_key, " ")
	var client_email string
	var client_id string
	var client_x509_cert_url string
	if isProd {
		projectId = "findnus-prod"
		private_key_id = "8c637a14442a5a1848af1da952d859148cfc063c"
		client_email = "firebase-adminsdk-ssly0@findnus-prod.iam.gserviceaccount.com"
		client_id = "117830696574462735012"
		client_x509_cert_url = "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-ssly0%40findnus-prod.iam.gserviceaccount.com"
	} else {
		projectId = "findnus-dev"
		private_key_id = "0cdf90c387f5d81121bcccb8ba1f9403e77cf2a4"
		client_email = "firebase-adminsdk-9zxcr@findnus-dev.iam.gserviceaccount.com"
		client_id = "116839976717740702813"
		client_x509_cert_url = "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-9zxcr%40findnus-dev.iam.gserviceaccount.com"
	}
	return map[string]string{
		"type":                        "service_account",
		"project_id":                  projectId,
		"private_key_id":              private_key_id,
		"private_key":                 private_key,
		"client_email":                client_email,
		"client_id":                   client_id,
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        client_x509_cert_url,
	}
}

// Creates an Admin SDK instance for the backend to do Firebase Auth Operations.
func InitFirebase() firebase.App {
	var app *firebase.App
	var err error
	config := &firebase.Config{}
	// Check if we are running locally with the existence of secrets folder
	// if _, err := os.Stat("../../secrets"); err == nil {
	// 	opt := option.WithCredentialsFile("../../secrets/findnus-dev-firebase-adminsdk-9zxcr-0cdf90c387.json")
	// 	app, err = firebase.NewApp(context.Background(), config, opt)
	// } else {
	// We are in a staged environment (Github || Heroku)
	// Get our credentials and load up firebase
	prodVar, _ := os.LookupEnv("PRODUCTION")
	var credbyte []byte
	if prodVar == "true" {
		credbyte, _ = json.Marshal(GetGoogleCredJson(true))
	} else {
		println("hi")
		credbyte, _ = json.Marshal(GetGoogleCredJson(false))
	}
	opt := option.WithCredentialsJSON(credbyte)
	app, err = firebase.NewApp(context.Background(), config, opt)
	// }
	if err != nil {
		log.Fatalf("Error init-ing firebase, %v\n", err)
	}
	return *app
}

// AuthGuard Middleware that checks if the client-side requester can use a priviledged API handler.
// Aborts executing the priviledged handler if user is not authorised.
func CheckAuthMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appCtx := context.Background()
		authClient, err := app.Auth(appCtx)
		if err != nil {
			log.Fatal(err)
		}
		idToken := ctx.GetHeader("Authorization")
		// WARN: Does not check for revocation status
		decodedToken, _ := authClient.VerifyIDToken(ctx, idToken)
		// Docs: https://github.com/firebase/firebase-admin-go/blob/v3.13.0/auth/auth.go#L277
		if decodedToken == nil {
			// Token has issues. Stop processing the request.
			ctx.AbortWithStatusJSON(
				401, gin.H{
					"message": "Unauthorised. Have you logged in?",
				},
			)
		}
	}
}
