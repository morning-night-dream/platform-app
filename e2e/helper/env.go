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

	return os.Getenv("CORE_ENDPOINT")
}

func GetAPIEndpoint(t *testing.T) string {
	t.Helper()

	return os.Getenv("API_ENDPOINT")
}

func GetDSN(t *testing.T) string {
	t.Helper()

	return os.Getenv("DATABASE_URL")
}
