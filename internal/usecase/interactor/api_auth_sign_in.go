package interactor

import (
	"context"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

type APIAuthSignIn struct {
	authRepository    repository.APIAuth
	sessionRepository repository.APISession
}

func NewAPIAuthSignIn(
	authRepository repository.APIAuth,
	sessionRepository repository.APISession,
) port.APIAuthSignIn {
	return &APIAuthSignIn{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
	}
}

func (aas *APIAuthSignIn) Execute(
	ctx context.Context,
	input port.APIAuthSignInInput,
) (port.APIAuthSignInOutput, error) {
	auth, err := aas.authRepository.SignIn(ctx, input.EMail, input.Password)
	if err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	sid := uuid.New().String()

	session := model.Session{
		SessionID: model.SessionID(sid),
		UserID:    auth.UserID,
		PublicKey: input.PublicKey,
	}

	// トランザクション必要か

	if err := aas.sessionRepository.Save(ctx, session); err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	if err := aas.authRepository.Save(ctx, auth); err != nil {
		if err := aas.sessionRepository.Delete(ctx, model.SessionID(sid)); err != nil {
			log.GetLogCtx(ctx).Warn("failed to delete session", log.ErrorField(err))
		}

		return port.APIAuthSignInOutput{}, err
	}

	sidToken, err := model.GenerateToken(sid, "secret")
	if err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	uidToken, err := model.GenerateToken(string(auth.UserID), sid)
	if err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	return port.APIAuthSignInOutput{
		UserToken:    model.UserToken(uidToken),
		SessionToken: model.SessionToken(sidToken),
	}, nil
}
