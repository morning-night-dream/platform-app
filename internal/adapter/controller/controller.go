package controller

import (
	"context"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/driver/firebase"
	"github.com/morning-night-dream/platform-app/internal/driver/redis"
)

type Controller struct {
	firebase *firebase.Client
	cache    *redis.Client
}

func New(
	firebase *firebase.Client,
	cache *redis.Client,
) *Controller {
	return &Controller{
		firebase: firebase,
		cache:    cache,
	}
}

func (ctl *Controller) Authorize(ctx context.Context, header http.Header) (model.Auth, error) {
	cookie, err := ctl.getToken(header)
	if err != nil {
		return model.Auth{}, ErrUnauthorized
	}

	auth, err := ctl.cache.Get(ctx, cookie.Value)
	if err != nil {
		return model.Auth{}, ErrUnauthorized
	}

	if err := ctl.firebase.VerifyIDToken(ctx, string(auth.IDToken)); err != nil {
		return model.Auth{}, ErrUnauthorized
	}

	return auth, nil
}

func (ctl *Controller) getToken(header http.Header) (http.Cookie, error) {
	lines := header["Cookie"]
	if len(lines) == 0 {
		return http.Cookie{}, ErrUnauthorized
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
			if name != "token" {
				return http.Cookie{}, ErrUnauthorized
			}
			return http.Cookie{Name: name, Value: val}, nil
		}
	}

	return http.Cookie{}, ErrUnauthorized
}

func (ctl *Controller) GetSession(header http.Header) (string, error) {
	lines := header["Cookie"]
	if len(lines) == 0 {
		return "", ErrUnauthorized
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
			if name != "token" {
				return "", ErrUnauthorized
			}
			return val, nil
		}
	}

	return "", ErrUnauthorized
}
