package controller

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-dream/platform-app/internal/usecase/port"
	userv1 "github.com/morning-night-dream/platform-app/pkg/connect/user/v1"
)

type User struct {
	ctl    *Controller
	signUp port.CoreUserSignUp
}

func NewUser(
	ctl *Controller,
	signUp port.CoreUserSignUp,
) *User {
	return &User{
		ctl:    ctl,
		signUp: signUp,
	}
}

func (us *User) SignUp(
	ctx context.Context,
	req *connect.Request[userv1.SignUpRequest],
) (*connect.Response[userv1.SignUpResponse], error) {
	input := port.CoreUserSignUpInput{}

	output, err := us.signUp.Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to execute sign up: %w", err)
	}

	return connect.NewResponse(&userv1.SignUpResponse{
		User: &userv1.User{
			Id: string(output.User.UserID),
		},
	}), nil
}

func (us *User) SignIn(
	ctx context.Context,
	req *connect.Request[userv1.SignInRequest],
) (*connect.Response[userv1.SignInResponse], error) {
	return connect.NewResponse(&userv1.SignInResponse{}), nil
}
