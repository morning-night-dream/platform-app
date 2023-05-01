package interactor

import (
	"context"
	"crypto"
	"crypto/rsa"
	"encoding/base64"

	"github.com/morning-night-dream/platform-app/internal/domain/cache"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

type APIAuthRefresh struct {
	sessionCache cache.Cache[model.Session]
	codeCache    cache.Cache[model.Code]
}

func NewAPIAuthRefresh(
	sessionCache cache.Cache[model.Session],
	codeCache cache.Cache[model.Code],
) port.APIAuthRefresh {
	return &APIAuthRefresh{
		sessionCache: sessionCache,
		codeCache:    codeCache,
	}
}

func (aar *APIAuthRefresh) Execute(
	ctx context.Context,
	input port.APIAuthRefreshInput,
) (port.APIAuthRefreshOutput, error) {
	// code から session id 取得
	code, err := aar.codeCache.Get(ctx, string(input.CodeID))
	if err != nil {
		return port.APIAuthRefreshOutput{}, err
	}

	// session id から public key を取得
	session, err := aar.sessionCache.Get(ctx, string(code.SessionID))
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get session", log.ErrorField(err))

		return port.APIAuthRefreshOutput{}, err
	}

	// code の署名検証
	// TODO ビジネスロジックなのでdomainモデルに移行する
	h := crypto.Hash.New(crypto.SHA256)

	h.Write([]byte(input.CodeID))

	hashed := h.Sum(nil)

	sig, err := base64.StdEncoding.DecodeString(string(input.Signature))
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to decode signature", log.ErrorField(err))

		return port.APIAuthRefreshOutput{}, err
	}

	key, err := session.RSAPublicKey()
	if err != nil {
		log.GetLogCtx(ctx).Warn("failed to get rsa public key", log.ErrorField(err))

		return port.APIAuthRefreshOutput{}, err
	}

	if err := rsa.VerifyPSS(key, crypto.SHA256, hashed, sig, &rsa.PSSOptions{
		Hash: crypto.SHA256,
	}); err != nil {
		log.GetLogCtx(ctx).Warn("failed to verify signature", log.ErrorField(err))

		return port.APIAuthRefreshOutput{}, err
	}

	// code の削除
	if err := aar.codeCache.Del(ctx, string(input.CodeID)); err != nil {
		return port.APIAuthRefreshOutput{}, err
	}

	idToken, err := model.GenerateToken(session.UserId, session.SessionId)
	if err != nil {
		return port.APIAuthRefreshOutput{}, err
	}

	return port.APIAuthRefreshOutput{
		IDToken: model.IDToken(idToken),
	}, nil
}
