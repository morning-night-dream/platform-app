package interactor

import (
	"context"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
)

type APIAuthSignUp struct {
	authRepository    repository.APIAuth
	sessionRepository repository.APISession
}

func NewAPIAuthSignUp(
	authRepository repository.APIAuth,
	sessionRepository repository.APISession,
) port.APIAuthSignUp {
	return &APIAuthSignUp{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
	}
}

func (aas *APIAuthSignUp) Execute(
	ctx context.Context,
	input port.APIAuthSignUpInput,
) (port.APIAuthSignUpOutput, error) {
	uid := uuid.New().String()

	if err := aas.authRepository.SignUp(ctx, model.UserID(uid), input.EMail, input.Password); err != nil {
		return port.APIAuthSignUpOutput{}, err
	}

	return port.APIAuthSignUpOutput{}, nil
}
