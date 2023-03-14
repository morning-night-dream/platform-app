package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
	"github.com/morning-night-dream/platform-app/internal/driver/public"
	"github.com/morning-night-dream/platform-app/internal/driver/store"
	"github.com/morning-night-dream/platform-app/internal/driver/user"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

var _ openapi.ServerInterface = (*Handler)(nil)

const HeaderApiKey = "Api-Key"

type Auth struct {
	signIn  port.APIAuthSignIn
	signOut port.APIAuthSignOut
	signUp  port.APIAuthSignUp
	verify  port.APIAuthVerify
	refresh port.APIAuthRefresh
	code    port.APIAuthGenerateCode
}

func NewAuth(
	signIn port.APIAuthSignIn,
	signOut port.APIAuthSignOut,
	signUp port.APIAuthSignUp,
	verify port.APIAuthVerify,
	refresh port.APIAuthRefresh,
	code port.APIAuthGenerateCode,
) *Auth {
	return &Auth{
		signIn:  signIn,
		signOut: signOut,
		signUp:  signUp,
		verify:  verify,
		refresh: refresh,
		code:    code,
	}
}

type Handler struct {
	version  string
	key      string
	auth     *Auth
	client   *Client
	store    *store.Store
	firebase *firebase.Client
	public   *public.Public
	user     *user.User
}

func New(
	version string,
	key string,
	auth *Auth,
	client *Client,
	store *store.Store,
	firebase *firebase.Client,
	public *public.Public,
	user *user.User,
) *Handler {
	return &Handler{
		version:  version,
		key:      key,
		auth:     auth,
		client:   client,
		store:    store,
		firebase: firebase,
		public:   public,
		user:     user,
	}
}

func (hdl *Handler) IsUnauthorizedAPIKey(w http.ResponseWriter, r *http.Request) bool {
	key := r.Header.Get(HeaderApiKey)

	if key == hdl.key {
		return false
	}

	w.WriteHeader(http.StatusUnauthorized)

	log.GetLogCtx(r.Context()).Warn(fmt.Sprintf("invalid api key: %s", key))

	return true
}

func (hdl *Handler) Authorize(r *http.Request) (model.Auth, error) {
	ctx := r.Context()

	uid, err := r.Cookie(model.IDTokenKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		return model.Auth{}, errors.New("error")
	}

	auth, err := hdl.store.Get(ctx, uid.Value)
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	if err := hdl.firebase.VerifyIDToken(ctx, string(auth.IDToken)); err != nil {
		return model.Auth{}, errors.New("error")
	}

	return auth, nil
}

func (hdl *Handler) Refresh(r *http.Request) (model.Auth, error) {
	ctx := r.Context()

	uid, err := r.Cookie(model.IDTokenKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		return model.Auth{}, errors.New("error")
	}

	auth, err := hdl.store.Get(ctx, uid.Value)
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	refresh, err := hdl.firebase.RefreshToken(ctx, string(auth.RefreshToken))
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	auth.RefreshToken = model.RefreshToken(refresh)

	if err := hdl.store.Set(ctx, uid.Value, auth); err != nil {
		return model.Auth{}, errors.New("error")
	}

	return auth, nil
}
