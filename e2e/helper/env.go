package helper

import (
	"os"
	"testing"
)

func GetAPIKey(t *testing.T) string {
	t.Helper()

	return os.Getenv("API_KEY")
}

func GetCoreEndpoint(t *testing.T) string {
	t.Helper()

	endpoint := os.Getenv("CORE_ENDPOINT")

	if endpoint == "" {
		return "http://localhost:8081"
	}

	return endpoint
}

func GetAPIEndpoint(t *testing.T) string {
	t.Helper()

	endpoint := os.Getenv("API_ENDPOINT")

	if endpoint == "" {
		return "http://localhost:8082"
	}

	return endpoint
}

func GetDSN(t *testing.T) string {
	t.Helper()

	return os.Getenv("DATABASE_URL")
}
