package external

import (
	"context"
	"fmt"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/rpc"
	userv1 "github.com/morning-night-dream/platform-app/pkg/connect/user/v1"
	"github.com/morning-night-dream/platform-app/pkg/connect/user/v1/userv1connect"
)

var _ rpc.User = (*User)(nil)

type UserFactory interface {
	User(string) (*User, error)
}

type User struct {
	connect userv1connect.UserServiceClient
}

func NewUser(
	connect userv1connect.UserServiceClient,
) *User {
	return &User{
		connect: connect,
	}
}

func (us *User) Create(ctx context.Context) (model.User, error) {
	req := NewRequestWithTID(ctx, &userv1.CreateRequest{})

	user, err := us.connect.Create(ctx, req)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to sign up: %w", err)
	}

	return model.User{
		UserID: model.UserID(user.Msg.User.Id),
	}, nil
}
