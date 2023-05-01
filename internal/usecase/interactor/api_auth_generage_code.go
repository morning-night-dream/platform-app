package interactor

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/cache"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
)

type ApiAuthGenerateCode struct {
	codeCache cache.Cache[model.Code]
}

func NewAPIAuthGenerateCode(
	codeCache cache.Cache[model.Code],
) port.APIAuthGenerateCode {
	return &ApiAuthGenerateCode{
		codeCache: codeCache,
	}
}

func (aac *ApiAuthGenerateCode) Execute(
	ctx context.Context,
	input port.APIAuthGenerateCodeInput,
) (port.APIAuthGenerateCodeOutput, error) {
	sessionID, err := model.GetID(string(input.SessionToken), "secret")
	if err != nil {
		return port.APIAuthGenerateCodeOutput{}, err
	}

	codeID := model.CodeID(uuid.NewString())

	code := &model.Code{
		CodeID:    codeID,
		SessionID: model.SessionID(sessionID),
	}

	if err := aac.codeCache.Set(ctx, string(codeID), code, time.Duration(time.Minute)); err != nil {
		return port.APIAuthGenerateCodeOutput{}, err
	}

	return port.APIAuthGenerateCodeOutput{
		CodeID: codeID,
	}, nil
}
