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

	log.GetLogCtx(ctx).Info(fmt.Sprintf("id: %s", payload.UserID))

	log.GetLogCtx(ctx).Info(fmt.Sprintf("exp: %d", exp))

	ctl.store.Set(ctx, payload.UserID, model.Auth{
		ID:           payload.UserID,
		UserID:       payload.UserID,
		IDToken:      res.IDToken,
		RefreshToken: res.RefreshToken,
		SessionToken: payload.UserID,
		ExpiresIn:    60,
		Expires:      time.Now().Add(60 * time.Second),
	})

	cookie := &http.Cookie{
		Name:     key,
		Value:    payload.UserID,
		Expires:  time.Now().Add(60 * time.Second),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, cookie)

	_, _ = w.Write([]byte("OK"))
}

// (POST /v1/auth/signout)
func (ctl *Controller) V1AuthSignOut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.GetLogCtx(ctx).Info(fmt.Sprintf("header: %+v", r.Header))

	uid, err := ctl.GetUserID(r.Header)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if err := ctl.store.Delete(ctx, uid); err != nil {
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

	uid, err := r.Cookie(key)
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
