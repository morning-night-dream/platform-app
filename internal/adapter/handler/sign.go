package handler

import (
	"net/http"

	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

// (GET /v1/sign).
func (hdl *Handler) V1Sign(w http.ResponseWriter, r *http.Request, params openapi.V1SignParams) {
	_, _ = w.Write([]byte("OK"))
}
