package interactor

import (
	"context"
	"fmt"

	"github.com/morning-night-dream/platform-app/internal/domain/rpc"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
)

type APIAuthSignUp struct {
	authRPC rpc.Auth
	userRPC rpc.User
}

func NewAPIAuthSignUp(
	authRPC rpc.Auth,
	userRPC rpc.User,
) port.APIAuthSignUp {
	return &APIAuthSignUp{
		authRPC: authRPC,
		userRPC: userRPC,
	}
}

func (aas *APIAuthSignUp) Execute(
	ctx context.Context,
	input port.APIAuthSignUpInput,
) (port.APIAuthSignUpOutput, error) {
	res, err := aas.userRPC.SignUp(ctx)
	if err != nil {
		return port.APIAuthSignUpOutput{}, fmt.Errorf("failed to user sign up: %w", err)
	}

	if err := aas.authRPC.SignUp(ctx, res.UserID, input.EMail, input.Password); err != nil {
		return port.APIAuthSignUpOutput{}, fmt.Errorf("failed to auth sign up: %w", err)
	}

	return port.APIAuthSignUpOutput{}, nil
}
