package interactor

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/cache"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

type APIAuthVerify struct {
	authCache cache.Cache[model.Auth]
}

func NewAPIAuthVerify(
	authCache cache.Cache[model.Auth],
) port.APIAuthVerify {
	return &APIAuthVerify{
		authCache: authCache,
	}
}

func (aav *APIAuthVerify) Execute(
	ctx context.Context,
	input port.APIAuthVerifyInput,
) (port.APIAuthVerifyOutput, error) {
	// サーバーの署名を確認
	sid, err := model.GetID(string(input.SessionToken), "secret")
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get sid", log.ErrorField(err))

		return port.APIAuthVerifyOutput{}, err
	}

	uid, err := model.GetID(string(input.IDToken), sid)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get uid", log.ErrorField(err))

		return port.APIAuthVerifyOutput{}, err
	}

	auth, err := aav.authCache.Get(ctx, uid)
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to find auth", log.ErrorField(err))

		return port.APIAuthVerifyOutput{}, err
	}

	if err := auth.Verify(); err != nil {
		return port.APIAuthVerifyOutput{}, err
	}

	return port.APIAuthVerifyOutput{}, nil
}
