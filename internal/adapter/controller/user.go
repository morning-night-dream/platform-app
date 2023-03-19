package controller

import (
	"context"

	"github.com/bufbuild/connect-go"
	userv1 "github.com/morning-night-dream/platform-app/pkg/connect/user/v1"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (us *User) SignUp(
	ctx context.Context,
	req *connect.Request[userv1.SignUpRequest],
) (*connect.Response[userv1.SignUpResponse], error) {
	return connect.NewResponse(&userv1.SignUpResponse{}), nil
}

func (us *User) SignIn(
	ctx context.Context,
	req *connect.Request[userv1.SignInRequest],
) (*connect.Response[userv1.SignInResponse], error) {
	return connect.NewResponse(&userv1.SignInResponse{}), nil
}
