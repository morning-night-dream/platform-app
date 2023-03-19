package config

import (
	"os"
)

type api struct {
	AppCoreURL       string
	APIKey           string
	CorsAllowOrigins string
}

var API api

func init() {
	API = api{
		AppCoreURL:       os.Getenv("APP_CORE_URL"),
		APIKey:           os.Getenv("API_KEY"),
		CorsAllowOrigins: os.Getenv("CORS_ALLOW_ORIGINS"),
	}
}
