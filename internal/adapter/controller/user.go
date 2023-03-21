package controller

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	userv1 "github.com/morning-night-dream/platform-app/pkg/connect/user/v1"
	"github.com/morning-night-dream/platform-app/pkg/connect/user/v1/userv1connect"
)

var _ userv1connect.UserServiceHandler = (*User)(nil)

type User struct {
	ctl    *Controller
	create port.CoreUserCreate
}

func NewUser(
	ctl *Controller,
	create port.CoreUserCreate,
) *User {
	return &User{
		ctl:    ctl,
		create: create,
	}
}

func (us *User) Create(
	ctx context.Context,
	req *connect.Request[userv1.CreateRequest],
) (*connect.Response[userv1.CreateResponse], error) {
	input := port.CoreUserCreateInput{}

	output, err := us.create.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to execute sign up: %w", err)
	}

	return connect.NewResponse(&userv1.CreateResponse{
		User: &userv1.User{
			Id: string(output.User.UserID),
		},
	}), nil
}

func (us *User) Update(
	ctx context.Context,
	req *connect.Request[userv1.UpdateRequest],
) (*connect.Response[userv1.UpdateResponse], error) {
	return connect.NewResponse(&userv1.UpdateResponse{}), nil
}
