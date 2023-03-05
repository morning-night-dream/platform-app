package api

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	versionv1 "github.com/morning-night-dream/platform-app/pkg/connect/version/v1"
)

func (api *API) V1APIVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(api.version))
}

func (api *API) V1CoreVersion(w http.ResponseWriter, r *http.Request) {
	req := &versionv1.ConfirmRequest{}
	res, err := api.client.Version.Confirm(r.Context(), connect.NewRequest(req))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)

		return
	}

	w.Write([]byte(res.Msg.Version))
}
