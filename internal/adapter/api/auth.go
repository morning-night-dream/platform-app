package api

import (
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

// (GET /v1/auth/refresh).
func (api *API) V1AuthRefresh(w http.ResponseWriter, r *http.Request, params openapi.V1AuthRefreshParams) {
	// リフレッシュに失敗したらキャッシュトークンは削除する
	ctx := r.Context()

	sidToken, err := r.Cookie(model.SIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		// sessin が存在しないのでcodeは生成しない
		rs := openapi.UnauthorizedResponse{}

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

	http.SetCookie(w, &http.Cookie{
		Name:     model.UIDKey,
		Value:    string(output.UserToken),
		Expires:  time.Now().Add(60 * time.Second),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	_, _ = w.Write([]byte("OK"))
}

// (POST /v1/auth/signin).
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

	input := port.APIAuthSignInInput{
		EMail:     model.EMail(email),
		Password:  model.Password(body.Password),
		PublicKey: key,
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
		Expires:  time.Now().Add(60 * time.Second),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     model.SIDKey,
		Value:    string(output.SessionToken),
		Expires:  time.Now().Add(168 * time.Hour),
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
	})

	_, _ = w.Write([]byte("OK"))
}

// (POST /v1/auth/signout).
func (api *API) V1AuthSignOut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.GetLogCtx(ctx).Info(fmt.Sprintf("header: %+v", r.Header))

	uid, err := r.Cookie(model.UIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	if err := api.store.Delete(ctx, uid.Value); err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// (POST /v1/auth/signup).
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

// (GET /v1/auth/verify).
func (api *API) V1AuthVerify(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.GetLogCtx(ctx).Info(fmt.Sprintf("header: %+v", r.Header))

	sidToken, err := r.Cookie(model.SIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		w.WriteHeader(http.StatusUnauthorized)

		// sessin が存在しないのでcodeは生成しない
		rs := openapi.UnauthorizedResponse{}

		if err := json.NewEncoder(w).Encode(rs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	uidToken, err := r.Cookie(model.UIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		// TODO 共通化
		input := port.APIAuthGenerateCodeInput{
			SessionToken: model.SessionToken(sidToken.Value),
		}

		output, err := api.auth.code.Execute(ctx, input)
		if err != nil {
			log.GetLogCtx(ctx).Warn("failed to execute", log.ErrorField(err))

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusUnauthorized)

		rs := openapi.UnauthorizedResponse{
			Code: uuid.MustParse(string(output.CodeID)),
		}

		if err := json.NewEncoder(w).Encode(rs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	input := port.APIAuthVerifyInput{
		IDToken:      model.IDToken(uidToken.Value),
		SessionToken: model.SessionToken(sidToken.Value),
	}

	if _, err := api.auth.verify.Execute(ctx, input); err != nil {
		log.GetLogCtx(ctx).Warn("failed to execute", log.ErrorField(err))

		// TODO 共通化
		input := port.APIAuthGenerateCodeInput{
			SessionToken: model.SessionToken(sidToken.Value),
		}

		output, err := api.auth.code.Execute(ctx, input)
		if err != nil {
			log.GetLogCtx(ctx).Warn("failed to execute", log.ErrorField(err))

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusUnauthorized)

		rs := openapi.UnauthorizedResponse{
			Code: uuid.MustParse(string(output.CodeID)),
		}

		if err := json.NewEncoder(w).Encode(rs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}
}

func (*API) V1AuthResign(w http.ResponseWriter, r *http.Request) {
	// TODO: not implemented
}
