package interactor

import (
	"context"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/cache"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

type APIAuthSignIn struct {
	authCache      cache.Cache[model.Auth]
	authRepository repository.APIAuth
	sessionCache   cache.Cache[model.Session]
}

func NewAPIAuthSignIn(
	authRepository repository.APIAuth,
	authCache cache.Cache[model.Auth],
	sessionCache cache.Cache[model.Session],
) port.APIAuthSignIn {
	return &APIAuthSignIn{
		authRepository: authRepository,
		authCache:      authCache,
		sessionCache:   sessionCache,
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
	if err := aas.sessionCache.Set(ctx, sid, session, model.Age); err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	if err := aas.authCache.Set(ctx, string(auth.UserID), auth, model.Age); err != nil {
		if err := aas.sessionCache.Del(ctx, sid); err != nil {
			log.GetLogCtx(ctx).Warn("failed to delete session", log.ErrorField(err))
		}

		return port.APIAuthSignInOutput{}, err
	}

	sidToken, err := model.GenerateToken(sid, "secret")
	if err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	idToken, err := model.GenerateToken(string(auth.UserID), sid)
	if err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	return port.APIAuthSignInOutput{
		UserID:       auth.UserID,
		SessionID:    model.SessionID(sid),
		IDToken:      model.IDToken(idToken),
		SessionToken: model.SessionToken(sidToken),
	}, nil
}
