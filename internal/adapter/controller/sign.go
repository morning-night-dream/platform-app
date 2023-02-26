package controller

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

// (GET /v1/sign).
func (ctl *Controller) V1Sign(w http.ResponseWriter, r *http.Request, params openapi.V1SignParams) {
	ctx := r.Context()

	sidToken, err := r.Cookie(model.SIDKey)
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

	uid, err := ctl.user.Get(ctx, sid)
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

	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(params.Code))
	hashed := h.Sum(nil)

	sig, err := base64.StdEncoding.DecodeString(params.Signature)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode signature", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := rsa.VerifyPKCS1v15(auth.PublicKey, crypto.SHA256, hashed, sig); err != nil {
		log.GetLogCtx(ctx).Warn("failed to verify signature", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	log.GetLogCtx(ctx).Info("signature verified")

	http.SetCookie(w, &http.Cookie{
		Name:     model.SIDKey,
		Value:    uuid.NewString(),
		Expires:  time.Now().Add(168 * time.Hour),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	_, _ = w.Write([]byte("OK"))
}
