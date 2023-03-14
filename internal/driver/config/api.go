package config

import (
	"os"
)

type api struct {
	AppCoreURL string
	APIKey     string
}

var API api

func init() {
	API = api{
		AppCoreURL: os.Getenv("APP_CORE_URL"),
		APIKey:     os.Getenv("API_KEY"),
	}
}
