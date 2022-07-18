/*
	This file describes logic to take care of firebase authentication
*/
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Programatically recreates the google service account JSON
// Do this because Github and Heroku do not have a native way of storing JSON secrets
func GetGoogleCredJson(isProd bool) map[string]string {
	var projectId string
	// Get private key from environment and replace unescaped newlines
	private_key, _ := os.LookupEnv("FIREBASE_KEY")
	private_key = strings.ReplaceAll(private_key, "\\n", "\n")
	private_key_id, _ := os.LookupEnv("FIREBASE_KEY_ID")
	var client_email string
	var client_id string
	var client_x509_cert_url string
	if isProd {
		projectId = "findnus-prod"
		client_email = "firebase-adminsdk-ssly0@findnus-prod.iam.gserviceaccount.com"
		client_id = "117830696574462735012"
		client_x509_cert_url = "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-ssly0%40findnus-prod.iam.gserviceaccount.com"
	} else {
		projectId = "findnus-dev"
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

var firebaseApp firebase.App

// Creates an Admin SDK instance for the backend to do Firebase Auth Operations.
func InitFirebase() firebase.App {
	var app *firebase.App
	var err error
	config := &firebase.Config{}
	// Check if we are running locally with the existence of secrets folder
	if _, err := os.Stat("../../secrets"); err == nil {
		opt := option.WithCredentialsFile("../../secrets/findnus-dev-firebase-adminsdk-9zxcr-0cdf90c387.json")
		app, err = firebase.NewApp(context.Background(), config, opt)
	} else {
		// We are in a staged environment (Github || Heroku)
		// Get our credentials and load up firebase
		prodVar, _ := os.LookupEnv("PRODUCTION")
		var credbyte []byte
		if prodVar == "true" {
			credbyte, _ = json.Marshal(GetGoogleCredJson(true))
		} else {
			credbyte, _ = json.Marshal(GetGoogleCredJson(false))
		}
		opt := option.WithCredentialsJSON(credbyte)
		app, err = firebase.NewApp(context.Background(), config, opt)
	}
	if err != nil {
		log.Fatalf("Error init-ing firebase, %v\n", err)
	}
	return *app
}

// Lookup a Lostee's email based on the User_id
func FirebaseGetEmailFromUser(user_id string) string {
	authClient, _ := firebaseApp.Auth(context.Background())
	userRecord, err := authClient.GetUser(context.TODO(), user_id)
	if err != nil {
		log.Println("Error getting firebase email:", err.Error())
		return ""
	}
	return userRecord.Email
}
