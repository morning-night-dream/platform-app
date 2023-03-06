package interactor

import (
	"context"
	"crypto"
	"crypto/rsa"
	"encoding/base64"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

type APIAuthRefresh struct {
	authRepository    repository.APIAuth
	sessionRepository repository.APISession
	codeRepository    repository.APICode
}

func NewAPIAuthRefresh(
	authRepository repository.APIAuth,
	sessionRepository repository.APISession,
	codeRepository repository.APICode,
) port.APIAuthRefresh {
	return &APIAuthRefresh{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
		codeRepository:    codeRepository,
	}
}

func (aar *APIAuthRefresh) Execute(
	ctx context.Context,
	input port.APIAuthRefreshInput,
) (port.APIAuthRefreshOutput, error) {
	// code から session id 取得
	code, err := aar.codeRepository.Find(ctx, input.CodeID)
	if err != nil {
		return port.APIAuthRefreshOutput{}, err
	}

	// session id から public key を取得
	session, err := aar.sessionRepository.Find(ctx, code.SessionID)
	if err != nil {
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

	if err := rsa.VerifyPSS(session.PublicKey, crypto.SHA256, hashed, []byte(sig), &rsa.PSSOptions{
		Hash: crypto.SHA256,
	}); err != nil {
		log.GetLogCtx(ctx).Warn("failed to verify signature", log.ErrorField(err))

		return port.APIAuthRefreshOutput{}, err
	}

	// code の削除
	if err := aar.codeRepository.Delete(ctx, input.CodeID); err != nil {
		return port.APIAuthRefreshOutput{}, err
	}

	uidToken, err := model.GenerateToken(string(session.UserID), string(session.SessionID))
	if err != nil {
		return port.APIAuthRefreshOutput{}, err
	}

	return port.APIAuthRefreshOutput{
		UserToken: model.UserToken(uidToken),
	}, nil
}
