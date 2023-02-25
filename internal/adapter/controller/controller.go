package controller

import (
	"context"
	"errors"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
	"github.com/morning-night-dream/platform-app/internal/driver/public"
	"github.com/morning-night-dream/platform-app/internal/driver/store"
	"github.com/morning-night-dream/platform-app/pkg/openapi"
)

var _ openapi.ServerInterface = (*Controller)(nil)

type Controller struct {
	client   *Client
	store    *store.Store
	firebase *firebase.Client
	public   *public.Public
}

func New(
	client *Client,
	store *store.Store,
	firebase *firebase.Client,
	public *public.Public,
) *Controller {
	return &Controller{
		client:   client,
		store:    store,
		firebase: firebase,
		public:   public,
	}
}

const (
	uidKey = "UID"
	sidKey = "SID"
)

func (ctl *Controller) Authorize(ctx context.Context, header http.Header) (model.Auth, error) {
	cookie, err := ctl.getToken(header)
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	auth, err := ctl.store.Get(ctx, cookie.Value)
	if err != nil {
		return model.Auth{}, errors.New("error")
	}

	if err := ctl.firebase.VerifyIDToken(ctx, auth.IDToken); err != nil {
		return model.Auth{}, errors.New("error")
	}

	return auth, nil
}

func (ctl *Controller) getToken(header http.Header) (http.Cookie, error) {
	lines := header["Cookie"]
	if len(lines) == 0 {
		return http.Cookie{}, errors.New("error")
	}

	for _, line := range lines {
		line = textproto.TrimString(line)

		var part string

		for len(line) > 0 { // continue since we have rest
			part, line, _ = strings.Cut(line, ";")
			part = textproto.TrimString(part)
			if part == "" {
				continue
			}
			name, val, _ := strings.Cut(part, "=")
			if name != uidKey {
				return http.Cookie{}, errors.New("error")
			}
			return http.Cookie{Name: name, Value: val}, nil
		}
	}

	return http.Cookie{}, errors.New("error")
}
