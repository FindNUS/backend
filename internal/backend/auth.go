/*
	This file describles logic to take care of firebase authentication
*/
package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// Creates an Admin SDK instance for the backend to do Firebase Auth Operations.
func InitFirebase() firebase.App {
	var app *firebase.App
	var err error
	// Check if we are running locally with the existence of secrets folder
	if _, err := os.Stat("../../secrets"); err == nil {
		opt := option.WithCredentialsFile("../../secrets/findnus-dev-firebase-adminsdk-9zxcr-0cdf90c387.json")
		config := &firebase.Config{}
		app, err = firebase.NewApp(context.Background(), config, opt)
	} else {
		// We are in a docker container on UAT/Prod. Env var to service account should exist
		app, err = firebase.NewApp(context.Background(), nil)
	}
	if err != nil {
		log.Fatalf("Error init firebase, %v\n", err)
	}
	log.Printf("Session ok!")
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
