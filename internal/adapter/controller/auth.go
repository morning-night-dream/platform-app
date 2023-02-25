package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

// (GET /v1/auth/refresh)
func (ctl *Controller) V1AuthRefresh(w http.ResponseWriter, r *http.Request, params openapi.V1AuthRefreshParams) {
	// リフレッシュに失敗したらキャッシュトークンは削除する
}

// (POST /v1/auth/signin)
func (ctl *Controller) V1AuthSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body openapi.V1AuthSignInJSONBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode request body", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	email, err := body.Email.MarshalJSON()
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to marshal email", log.ErrorField(err))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := ctl.firebase.Login(ctx, string(email), body.Password)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to login", log.ErrorField(err))

		w.WriteHeader(http.StatusForbidden)

		return
	}

	exp, _ := strconv.Atoi(res.ExpiresIn)

	strs := strings.Split(res.IDToken, ".")

	tmpPayload, err := base64.RawStdEncoding.DecodeString(strs[1])
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	type Payload struct {
		UserID string `json:"user_id"`
	}

	var payload Payload

	if err := json.Unmarshal(tmpPayload, &payload); err != nil {
		log.GetLogCtx(ctx).Warn("failed to unmarshal json "+string(tmpPayload), log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	pub, err := base64.StdEncoding.DecodeString(body.PublicKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode public key", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	log.GetLogCtx(ctx).Info(fmt.Sprintf("public key: %s", string(pub)))

	uid := payload.UserID

	sid := uuid.New().String()

	log.GetLogCtx(ctx).Info(fmt.Sprintf("id: %s", uid))

	log.GetLogCtx(ctx).Info(fmt.Sprintf("exp: %d", exp))

	if err := ctl.store.Set(ctx, uid, model.Auth{
		ID:           uid,
		UserID:       uid,
		IDToken:      res.IDToken,
		PublicKey:    pub,
		RefreshToken: res.RefreshToken,
		SessionToken: uid,
		ExpiresIn:    60,
		Expires:      time.Now().Add(60 * time.Second),
	}); err != nil {
		log.GetLogCtx(ctx).Warn("failed to set auth", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := ctl.user.Set(ctx, sid, uid); err != nil {
		log.GetLogCtx(ctx).Warn("failed to set public key", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     uidKey,
		Value:    payload.UserID,
		Expires:  time.Now().Add(60 * time.Second),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     sidKey,
		Value:    sid,
		Expires:  time.Now().Add(168 * time.Hour),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	_, _ = w.Write([]byte("OK"))
}

// (POST /v1/auth/signout)
func (ctl *Controller) V1AuthSignOut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.GetLogCtx(ctx).Info(fmt.Sprintf("header: %+v", r.Header))

	uid, err := r.Cookie(uidKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if err := ctl.store.Delete(ctx, uid.Value); err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// (POST /v1/auth/signup)
func (ctl *Controller) V1AuthSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body openapi.V1AuthSignUpJSONBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode request body", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	uid := uuid.New().String()

	log.GetLogCtx(ctx).Info(fmt.Sprintf("uid: %s", uid))

	email, err := body.Email.MarshalJSON()
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to marshal email", log.ErrorField(err))

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := ctl.firebase.CreateUser(ctx, uid, string(email), body.Password); err != nil {
		log.GetLogCtx(ctx).Warn("failed to create user", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// (GET /v1/auth/verify)
func (ctl Controller) V1AuthVerify(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.GetLogCtx(ctx).Info(fmt.Sprintf("header: %+v", r.Header))

	uid, err := r.Cookie(uidKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		rs := openapi.UnauthorizedResponse{}

		if err := json.NewEncoder(w).Encode(rs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	auth, err := ctl.store.Get(ctx, uid.Value)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	log.GetLogCtx(ctx).Info(fmt.Sprintf("expires: %s", auth.Expires))
}
