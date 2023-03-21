package interactor

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
)

type CoreUserCreate struct {
	userRepo repository.User
}

func NewCoreUserCreate(
	userRepo repository.User,
) port.CoreUserCreate {
	return &CoreUserCreate{
		userRepo: userRepo,
	}
}

func (aas *CoreUserCreate) Execute(
	ctx context.Context,
	input port.CoreUserCreateInput,
) (port.CoreUserCreateOutput, error) {
	uid := uuid.New().String()

	user := model.User{
		UserID: model.UserID(uid),
	}

	if err := aas.userRepo.Save(ctx, user); err != nil {
		return port.CoreUserCreateOutput{}, fmt.Errorf("failed to save user: %w", err)
	}

	return port.CoreUserCreateOutput{
		User: user,
	}, nil
}
