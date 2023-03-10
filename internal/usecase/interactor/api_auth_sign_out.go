package interactor

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/cache"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

type APIAuthSignOut struct {
	authCache    cache.Cache[model.Auth]
	sessionCache cache.Cache[model.Session]
}

func NewAPIAuthSignOut(
	authCache cache.Cache[model.Auth],
	sessionCache cache.Cache[model.Session],
) port.APIAuthSignOut {
	return &APIAuthSignOut{
		authCache:    authCache,
		sessionCache: sessionCache,
	}
}

func (aas *APIAuthSignOut) Execute(
	ctx context.Context,
	input port.APIAuthSignOutInput,
) (port.APIAuthSignOutOutput, error) {
	sid, err := model.GetID(string(input.SessionToken), "secret")
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get sid", log.ErrorField(err))

		return port.APIAuthSignOutOutput{}, err
	}

	uid, err := model.GetID(string(input.IDToken), sid)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get uid", log.ErrorField(err))

		return port.APIAuthSignOutOutput{}, err
	}

	// トランザクション必要か

	if err := aas.sessionCache.Del(ctx, sid); err != nil {
		return port.APIAuthSignOutOutput{}, err
	}

	if err := aas.authCache.Del(ctx, uid); err != nil {
		return port.APIAuthSignOutOutput{}, err
	}

	return port.APIAuthSignOutOutput{}, nil
}
