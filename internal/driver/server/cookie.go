package server

import (
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/driver/config"
)

func Secure() bool {
	return config.API.Domain != ""
}

func SameSiteMode() http.SameSite {
	if config.API.Domain == "" {
		return http.SameSiteDefaultMode
	}

	return http.SameSiteNoneMode
}
