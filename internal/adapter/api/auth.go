package api

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

// GET /v1/auth/refresh.
func (api *API) V1AuthRefresh(w http.ResponseWriter, r *http.Request, params openapi.V1AuthRefreshParams) {
	// リフレッシュに失敗したらキャッシュトークンは削除する
	ctx := r.Context()

	sidToken, err := r.Cookie(model.SIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		// session が存在しないのでcodeは生成しない
		rs := openapi.V1UnauthorizedResponse{}

		if err := json.NewEncoder(w).Encode(rs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	input := port.APIAuthRefreshInput{
		CodeID:       model.CodeID(params.Code),
		Signature:    model.Signature(params.Signature),
		SessionToken: model.SessionToken(sidToken.Value),
	}

	output, err := api.auth.refresh.Execute(ctx, input)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to execute", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	expires := model.DefaultExpires * time.Second
	if params.ExpiresIn != nil {
		if *params.ExpiresIn < 0 || *params.ExpiresIn > model.DefaultExpires {
			log.GetLogCtx(ctx).Warn("invalid expires_in", log.ErrorField(err))

			w.WriteHeader(http.StatusBadRequest)

			return
		}
		expires = time.Duration(*params.ExpiresIn) * time.Second
	}

	http.SetCookie(w, &http.Cookie{
		Name:     model.UIDKey,
		Value:    string(output.UserToken),
		Expires:  time.Now().Add(expires),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	_, _ = w.Write([]byte("OK"))
}

// POST /v1/auth/signin.
func (api *API) V1AuthSignIn(w http.ResponseWriter, r *http.Request) {
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

	// --------------------------------------------------
	// from frontend
	pubkey, err := base64.StdEncoding.DecodeString(body.PublicKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	pub, err := x509.ParsePKIXPublicKey(pubkey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to parse public key", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	key, ok := pub.(*rsa.PublicKey)
	if !ok {
		log.GetLogCtx(ctx).Warn("failed to parse public key", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	// --------------------------------------------------

	// --------------------------------------------------
	// from go code
	// pub, err := base64.StdEncoding.DecodeString(body.PublicKey)
	// if err != nil {
	// 	log.GetLogCtx(ctx).Warn("failed to decode public key", log.ErrorField(err))

	// 	w.WriteHeader(http.StatusInternalServerError)

	// 	return
	// }

	// var key *rsa.PublicKey
	// if err := json.Unmarshal(pub, &key); err != nil {
	// 	log.GetLogCtx(ctx).Warn("failed to unmarshal public key", log.ErrorField(err))

	// 	w.WriteHeader(http.StatusInternalServerError)

	// 	return
	// }
	// --------------------------------------------------

	expires := model.DefaultExpires * time.Second
	if body.ExpiresIn != nil {
		if *body.ExpiresIn < 0 || *body.ExpiresIn > model.DefaultExpires {
			log.GetLogCtx(ctx).Warn("invalid expires_in", log.ErrorField(err))

			w.WriteHeader(http.StatusBadRequest)

			return
		}
		expires = time.Duration(*body.ExpiresIn) * time.Second
	}

	input := port.APIAuthSignInInput{
		EMail:     model.EMail(email),
		Password:  model.Password(body.Password),
		PublicKey: key,
		ExpiresIn: model.ExpiresIn(expires),
	}

	output, err := api.auth.signIn.Execute(ctx, input)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to execute", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     model.UIDKey,
		Value:    string(output.UserToken),
		Expires:  time.Now().Add(expires),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     model.SIDKey,
		Value:    string(output.SessionToken),
		Expires:  time.Now().Add(model.Age),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	_, _ = w.Write([]byte("OK"))
}

// GET /v1/auth/signout.
func (api *API) V1AuthSignOut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sidToken, _ := r.Cookie(model.SIDKey)

	uidToken, _ := r.Cookie(model.UIDKey)

	input := port.APIAuthSignOutInput{
		IDToken:      model.IDToken(uidToken.Value),
		SessionToken: model.SessionToken(sidToken.Value),
	}

	_, _ = api.auth.signOut.Execute(ctx, input)

	http.SetCookie(w, &http.Cookie{
		Name:     model.UIDKey,
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     model.SIDKey,
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	_, _ = w.Write([]byte("OK"))
}

// POST /v1/auth/signup.
func (api *API) V1AuthSignUp(w http.ResponseWriter, r *http.Request) {
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

	input := port.APIAuthSignUpInput{
		EMail:    model.EMail(string(email)),
		Password: model.Password(body.Password),
	}

	if _, err := api.auth.signUp.Execute(ctx, input); err != nil {
		log.GetLogCtx(ctx).Warn("failed to sign up", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// GET /v1/auth/verify.
func (api *API) V1AuthVerify(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.GetLogCtx(ctx).Info(fmt.Sprintf("header: %+v", r.Header))

	sidToken, err := r.Cookie(model.SIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	uidToken, err := r.Cookie(model.UIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		api.unauthorize(w, r, ctx, model.SessionToken(sidToken.Value))

		return
	}

	input := port.APIAuthVerifyInput{
		IDToken:      model.IDToken(uidToken.Value),
		SessionToken: model.SessionToken(sidToken.Value),
	}

	if _, err := api.auth.verify.Execute(ctx, input); err != nil {
		log.GetLogCtx(ctx).Warn("failed to execute", log.ErrorField(err))

		api.unauthorize(w, r, ctx, model.SessionToken(sidToken.Value))
	}

	_, _ = w.Write([]byte("OK"))
}

// DELETE /v1/auth.
func (api *API) V1AuthResign(w http.ResponseWriter, r *http.Request) {
	// TODO: not implemented
}

func (api *API) unauthorize(w http.ResponseWriter, r *http.Request, ctx context.Context, stk model.SessionToken) {
	input := port.APIAuthGenerateCodeInput{
		SessionToken: stk,
	}

	output, err := api.auth.code.Execute(ctx, input)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to execute", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusUnauthorized)

	rs := openapi.V1UnauthorizedResponse{
		Code: uuid.MustParse(string(output.CodeID)),
	}

	if err := json.NewEncoder(w).Encode(rs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
