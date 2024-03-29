package handler

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	healthv1 "github.com/morning-night-dream/platform-app/pkg/connect/health/v1"
)

func (hdl *Handler) V1Health(w http.ResponseWriter, r *http.Request) {
	req := &healthv1.CheckRequest{}
	res, err := hdl.client.Health.Check(r.Context(), connect.NewRequest(req))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res.Msg.String()))
}
