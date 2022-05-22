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
	if tmp == "" {
		tmp = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDSGWJKPA0sIKeh\n8FrmFGeOtmL+XLsSIibIQZio+NoKHBAcU2cls46UHSIxCTJq0ZDXzF4ei6gm0Uqz\nMrUPMsMZv2em6754g3+42tVx9ykkbh3+8fIxzBS6v+nxA7P43X4Noh0FJ1/7uqIz\ni62ULdnblhpphHhpJ8LjqVVGsoloPakwfKNUh0SDz+XsyyxI0ufpcF1aNUZu5CHn\nqXikubSBvdmnTD+pQKc0FxaEqRaHL7h6Xzu2g6RERPpaWr4+HUkDJAaW9u4Vlwhb\nXGmzUyp0HVRXhfkhRbiG40jykyZ+oKD/td4SwCOFCzf4k0G/Jt3oEN6s4kVJ10LC\ntaH+3N0zAgMBAAECggEAHfaNzH18+W6cyZ0QMaD2VeWP/6u06DSjqEqmnW6EFg4D\nhC6m1rshWeE/x5OCs7Y4fHZCdAPB0utlRmI0bTr1lR31h9o2G1TRqcjXyP2RSgdE\nUuAphM2QpUOKdxtqltLrz8Dvd5UyfKGU0VoZwri5SbZCBQtl6sVHZ5V2OnNq4kkq\nmhhegQECZUzIAyRX3WbKf8Yu8eAM9tWsvNkO+W2wF+dIq106Hw15XjuJE6Z/kgam\nGYk1b1y6gEjgD+ZoPonLIRv+jWgOiQCCscSTXHnduQJAENtCNyEmBH6lF4qFufZu\n2ss1xH/5BrtdnZ4V3a0s+EUYQAtSHDCi2nmGYFdKUQKBgQD6BMrnDx/OOSaCZUMc\n9/kDxL/wxhnbmEiF2V2T2tl0EaO4kYQxynlZhZVeWPiXRVhjRVdzUs9yBZOb0vy/\nbu9WdyS8hDeOQRY1BuAmpWIUF9cSfQaMRFIbq7/P0zyw5XqaQT3lxQEA7rXIQ+d3\n+drhziRatqELIl+b+hX/1Sb6EQKBgQDXIBvv6xXiSBwMVe6pjjbJb7SG0QW+G1pP\nXQjOk2EX5fbc2ksaCZKbfaVN3YC5FyJdNrFXR+KLNQANEEbOSh4/mO+CISNSYUlF\n+gXTBzmYxHBohnFfS2C67R2YepIJ8Uj0GVU84eDkeZdJMiqMmg+H+68CxgTpCeEo\nBSGaTkH/AwKBgFNfEcInGvYLvLmyxsR8ND97dn31sV221Eg+CaRqUCUSVMQRUkHA\nQOMHVp3VkV/wMd84mkbMkHx3O5e0ra+wcIMmy8tJU7VOIvefyVNZxvDoWkHCC1Lu\n3Wp4xUeqKwzaGR4jL17VaNZEw716V098s+6kbR8K030BA1zh8kATdiHhAoGARXYH\nr0L/8O2JqO4CPts9k3MvHizVptmcIm4OzuzFd/r3573QbBrVLMG4I1k3HAx9Ow3S\n2zTJ0FsPpigwRKGn/K77/s+GYS4qg57ETKxTi6E6DnYCm1tyY0j2umoxR2aSQMcB\nP8RLYlpkX+0D0hxYkXbRvpqDsV9QRSTLAdDs/FMCgYEA6YI7Y4KwtbkOokwwMIkU\nAO8aJ0i9SLqm0nOIXHbIcKh3egCZXbnnoBI/NoSIlUWOp41M/dQqF3KQQBVAjalm\n4SfWhBnJ+t15aYu/Rgu9M6XxSCjgrThOFH4VHIC7Z8G2GFWXnTmUNYzKo700sGqA\nxBM3dmursyu6psdMc2oQk2U="
	}
	// log.Println(tmp[0:100])
	// log.Println(tmp[100:])
	private_key += tmp + "\n-----END PRIVATE KEY-----\n"
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
	log.Println("Hello, ", string(credbyte))
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
