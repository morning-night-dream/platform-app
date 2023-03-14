package handler

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

// (GET /v1/sign).
func (hdl *Handler) V1Sign(w http.ResponseWriter, r *http.Request, params openapi.V1SignParams) {
	ctx := r.Context()

	sidToken, err := r.Cookie(model.SessionTokenKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	sid, err := model.GetID(sidToken.Value, "secret")
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get sid", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	uid, err := hdl.user.Get(ctx, sid)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	auth, err := hdl.store.Get(ctx, uid)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(params.Code))
	hashed := h.Sum(nil)

	sig, err := base64.StdEncoding.DecodeString(params.Signature)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode signature", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := rsa.VerifyPSS(auth.PublicKey, crypto.SHA256, hashed, []byte(sig), &rsa.PSSOptions{
		Hash: crypto.SHA256,
	}); err != nil {
		log.GetLogCtx(ctx).Warn("failed to verify signature", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	log.GetLogCtx(ctx).Info("signature verified")

	_, _ = w.Write([]byte("OK"))
}
