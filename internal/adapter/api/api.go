package api

import (
	"errors"
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

var _ openapi.ServerInterface = (*API)(nil)

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

type API struct {
	version  string
	auth     *Auth
	client   *Client
	store    *store.Store
	firebase *firebase.Client
	public   *public.Public
	user     *user.User
}

func New(
	version string,
	auth *Auth,
	client *Client,
	store *store.Store,
	firebase *firebase.Client,
	public *public.Public,
	user *user.User,
) *API {
	return &API{
		version:  version,
		auth:     auth,
		client:   client,
		store:    store,
		firebase: firebase,
		public:   public,
		user:     user,
	}
}

func (ctl *API) Authorize(r *http.Request) (model.Auth, error) {
	ctx := r.Context()

	uid, err := r.Cookie(model.UIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		return model.Auth{}, errors.New("error")
	}

	auth, err := ctl.store.Get(ctx, uid.Value)
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	if err := ctl.firebase.VerifyIDToken(ctx, string(auth.IDToken)); err != nil {
		return model.Auth{}, errors.New("error")
	}

	return auth, nil
}

func (ctl *API) Refresh(r *http.Request) (model.Auth, error) {
	ctx := r.Context()

	uid, err := r.Cookie(model.UIDKey)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get auth", log.ErrorField(err))

		return model.Auth{}, errors.New("error")
	}

	auth, err := ctl.store.Get(ctx, uid.Value)
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	refresh, err := ctl.firebase.RefreshToken(ctx, string(auth.RefreshToken))
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	auth.RefreshToken = model.RefreshToken(refresh)

	if err := ctl.store.Set(ctx, uid.Value, auth); err != nil {
		return model.Auth{}, errors.New("error")
	}

	return auth, nil
}
