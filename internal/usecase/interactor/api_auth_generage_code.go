package interactor

import (
	"context"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
)

type ApiAuthGenerateCode struct {
	codeRepository repository.APICode
}

func NewAPIAuthGenerateCode(
	codeRepository repository.APICode,
) port.APIAuthGenerateCode {
	return &ApiAuthGenerateCode{
		codeRepository: codeRepository,
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

	if err := aac.codeRepository.Save(ctx, model.Code{
		CodeID:    codeID,
		SessionID: model.SessionID(sessionID),
	}); err != nil {
		return port.APIAuthGenerateCodeOutput{}, err
	}

	return port.APIAuthGenerateCodeOutput{
		CodeID: codeID,
	}, nil
}
