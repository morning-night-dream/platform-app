package interactor

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
)

type CoreUserSignUp struct {
	userRepo repository.User
}

func NewCoreUserSignUp(
	userRepo repository.User,
) port.CoreUserSignUp {
	return &CoreUserSignUp{
		userRepo: userRepo,
	}
}

func (aas *CoreUserSignUp) Execute(
	ctx context.Context,
	input port.CoreUserSignUpInput,
) (port.CoreUserSignUpOutput, error) {
	uid := uuid.New().String()

	user := model.User{
		UserID: model.UserID(uid),
	}

	if err := aas.userRepo.Save(ctx, user); err != nil {
		return port.CoreUserSignUpOutput{}, fmt.Errorf("failed to save user: %w", err)
	}

	return port.CoreUserSignUpOutput{
		User: user,
	}, nil
}
