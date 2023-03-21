package interactor

import (
	"context"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/rpc"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
)

type APIAuthSignUp struct {
	authRPC rpc.Auth
}

func NewAPIAuthSignUp(
	authRPC rpc.Auth,
) port.APIAuthSignUp {
	return &APIAuthSignUp{
		authRPC: authRPC,
	}
}

func (aas *APIAuthSignUp) Execute(
	ctx context.Context,
	input port.APIAuthSignUpInput,
) (port.APIAuthSignUpOutput, error) {
	uid := uuid.New().String()

	if err := aas.authRPC.SignUp(ctx, model.UserID(uid), input.EMail, input.Password); err != nil {
		return port.APIAuthSignUpOutput{}, err
	}

	return port.APIAuthSignUpOutput{}, nil
}
