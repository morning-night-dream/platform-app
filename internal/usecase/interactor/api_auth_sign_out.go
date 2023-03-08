package interactor

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

type APIAuthSignOut struct {
	authRepository    repository.APIAuth
	sessionRepository repository.APISession
}

func NewAPIAuthSignOut(
	authRepository repository.APIAuth,
	sessionRepository repository.APISession,
) port.APIAuthSignOut {
	return &APIAuthSignOut{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
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

	if err := aas.sessionRepository.Delete(ctx, model.SessionID(sid)); err != nil {
		return port.APIAuthSignOutOutput{}, err
	}

	if err := aas.authRepository.Delete(ctx, model.UserID(uid)); err != nil {
		return port.APIAuthSignOutOutput{}, err
	}

	return port.APIAuthSignOutOutput{}, nil
}
