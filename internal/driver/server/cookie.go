package server

import (
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/driver/config"
)

func CookieSecure() bool {
	return config.API.Domain != ""
}

func CookieSameSiteMode() http.SameSite {
	if config.API.Domain == "" {
		return http.SameSiteDefaultMode
	}

	return http.SameSiteNoneMode
}
