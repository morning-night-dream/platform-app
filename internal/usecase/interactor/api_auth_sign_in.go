package interactor

import (
	"context"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/cache"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/rpc"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

type APIAuthSignIn struct {
	authCache    cache.Cache[model.Auth]
	authRPC      rpc.Auth
	sessionCache cache.Cache[model.Session]
}

func NewAPIAuthSignIn(
	authRPC rpc.Auth,
	authCache cache.Cache[model.Auth],
	sessionCache cache.Cache[model.Session],
) port.APIAuthSignIn {
	return &APIAuthSignIn{
		authRPC:      authRPC,
		authCache:    authCache,
		sessionCache: sessionCache,
	}
}

func (aas *APIAuthSignIn) Execute(
	ctx context.Context,
	input port.APIAuthSignInInput,
) (port.APIAuthSignInOutput, error) {
	user, err := aas.authRPC.SignIn(ctx, input.EMail, input.Password)
	if err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	sid := uuid.New().String()

	key, err := model.PublicKeyToString(input.PublicKey)
	if err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	session := &model.Session{
		SessionId: sid,
		UserId:    user.UserId,
		PublicKey: key,
	}

	// トランザクション必要か
	if err := aas.sessionCache.Set(ctx, sid, session, model.Age); err != nil {
		return port.APIAuthSignInOutput{}, err
	}

	auth := &model.Auth{
		ID:     user.UserId,
		UserID: user.UserId,
	}

	if err := aas.authCache.Set(ctx, user.UserId, auth, model.Age); err != nil {
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
