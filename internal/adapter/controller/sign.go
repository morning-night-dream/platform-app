package controller

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

// (POST /v1/sign)
func (ctl *Controller) V1Sign(w http.ResponseWriter, r *http.Request, params openapi.V1SignParams) {
	ctx := r.Context()

	sid, err := r.Cookie(sidKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		rs := openapi.UnauthorizedResponse{}

		if err := json.NewEncoder(w).Encode(rs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	uid, err := ctl.user.Get(ctx, sid.Value)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	auth, err := ctl.store.Get(ctx, uid)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	log.GetLogCtx(ctx).Info(fmt.Sprintf("pubkey: %s", auth.PublicKey))

	pub, err := x509.ParsePKIXPublicKey(auth.PublicKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to parse public key", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	log.GetLogCtx(ctx).Info(fmt.Sprintf("pubkey: %+v", pub))

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.GetLogCtx(ctx).Warn("failed to parse public key", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	log.GetLogCtx(ctx).Info(fmt.Sprintf("pubkey: %+v", publicKey))

	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(params.Code))
	hashed := h.Sum(nil)

	rsa.VerifyPSS(publicKey, crypto.SHA256, hashed, []byte(params.Signature), nil)
}
