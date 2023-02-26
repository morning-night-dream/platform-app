package controller

import (
	"errors"
	"net/http"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
	"github.com/morning-night-dream/platform-app/internal/driver/public"
	"github.com/morning-night-dream/platform-app/internal/driver/store"
	"github.com/morning-night-dream/platform-app/internal/driver/user"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

var _ openapi.ServerInterface = (*Controller)(nil)

type Controller struct {
	client   *Client
	store    *store.Store
	firebase *firebase.Client
	public   *public.Public
	user     *user.User
}

func New(
	client *Client,
	store *store.Store,
	firebase *firebase.Client,
	public *public.Public,
	user *user.User,
) *Controller {
	return &Controller{
		client:   client,
		store:    store,
		firebase: firebase,
		public:   public,
		user:     user,
	}
}

func (ctl *Controller) Authorize(r *http.Request) (model.Auth, error) {
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

	if err := ctl.firebase.VerifyIDToken(ctx, auth.IDToken); err != nil {
		return model.Auth{}, errors.New("error")
	}

	return auth, nil
}

func (ctl *Controller) Refresh(r *http.Request) (model.Auth, error) {
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

	refresh, err := ctl.firebase.RefreshToken(ctx, auth.RefreshToken)
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	auth.RefreshToken = refresh

	if err := ctl.store.Set(ctx, uid.Value, auth); err != nil {
		return model.Auth{}, errors.New("error")
	}

	return auth, nil
}
