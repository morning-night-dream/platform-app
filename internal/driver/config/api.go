package config

import (
	"os"
)

type api struct {
	AppCoreURL          string
	APIKey              string
	CorsAllowOrigins    string
	Domain              string
	RedisURL            string
	FirebaseSecret      string
	FirebaseAPIEndpoint string
	FirebaseAPIKey      string
}

var API api

func init() {
	API = api{
		AppCoreURL:          os.Getenv("APP_CORE_URL"),
		APIKey:              os.Getenv("API_KEY"),
		CorsAllowOrigins:    os.Getenv("CORS_ALLOW_ORIGINS"),
		Domain:              os.Getenv("DOMAIN"),
		RedisURL:            os.Getenv("REDIS_URL"),
		FirebaseSecret:      os.Getenv("FIREBASE_SECRET"),
		FirebaseAPIEndpoint: os.Getenv("FIREBASE_API_ENDPOINT"),
		FirebaseAPIKey:      os.Getenv("FIREBASE_API_KEY"),
	}
}
