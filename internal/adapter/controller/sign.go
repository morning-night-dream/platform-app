package controller

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

// (GET /v1/sign)
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

	// log.GetLogCtx(ctx).Info(fmt.Sprintf("pubkey: %s", auth.PublicKey))

	// publicKey, err := x509.ParsePKCS1PublicKey(auth.PublicKey)
	// if err != nil {
	// 	log.GetLogCtx(ctx).Warn("failed to parse public key", log.ErrorField(err))

	// 	w.WriteHeader(http.StatusInternalServerError)

	// 	return
	// }

	// pub, err := x509.ParsePKIXPublicKey(auth.PublicKey)
	// if err != nil {
	// 	log.GetLogCtx(ctx).Warn("failed to parse public key", log.ErrorField(err))

	// 	w.WriteHeader(http.StatusInternalServerError)

	// 	return
	// }

	// log.GetLogCtx(ctx).Info(fmt.Sprintf("pubkey: %+v", pub))

	// publicKey, ok := pub.(*rsa.PublicKey)
	// if !ok {
	// 	log.GetLogCtx(ctx).Warn("failed to parse public key", log.ErrorField(err))

	// 	w.WriteHeader(http.StatusInternalServerError)

	// 	return
	// }

	// log.GetLogCtx(ctx).Info(fmt.Sprintf("pubkey: %+v", publicKey))

	var public *rsa.PublicKey
	if err := json.Unmarshal(auth.PublicKey, &public); err != nil {
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

	if err := rsa.VerifyPKCS1v15(public, crypto.SHA256, hashed, sig); err != nil {
		log.GetLogCtx(ctx).Warn("failed to verify signature", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	log.GetLogCtx(ctx).Info("signature verified")

	_, _ = w.Write([]byte("OK"))
}
