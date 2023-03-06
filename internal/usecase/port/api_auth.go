package port

import (
	"crypto/rsa"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/usecase"
)

type APIAuthSignUpInput struct {
	usecase.Input
	model.EMail
	model.Password
}

type APIAuthSignUpOutput struct {
	usecase.Output
}

type APIAuthSignUp interface {
	usecase.Usecase[APIAuthSignUpInput, APIAuthSignUpOutput]
}

type APIAuthSignInInput struct {
	usecase.Input
	model.EMail
	model.Password
	PublicKey *rsa.PublicKey
}

type APIAuthSignInOutput struct {
	usecase.Output
	model.UserID
	model.SessionID
	model.UserToken
	model.SessionToken
}

type APIAuthSignIn interface {
	usecase.Usecase[APIAuthSignInInput, APIAuthSignInOutput]
}

type APIAuthRefreshInput struct {
	usecase.Input
	model.CodeID
	model.Signature
	model.SessionToken
}

type APIAuthRefreshOutput struct {
	usecase.Output
	model.UserToken
}

type APIAuthRefresh interface {
	usecase.Usecase[APIAuthRefreshInput, APIAuthRefreshOutput]
}

type APIAuthVerifyInput struct {
	usecase.Input
	model.IDToken
	model.SessionToken
}

type APIAuthVerifyOutput struct {
	usecase.Output
}

type APIAuthVerify interface {
	usecase.Usecase[APIAuthVerifyInput, APIAuthVerifyOutput]
}

type APIAuthSignOutInput struct {
	usecase.Input
}

type APIAuthSignOutOutput struct {
	usecase.Output
}

type APIAuthSignOut interface {
	usecase.Usecase[APIAuthSignOutInput, APIAuthSignOutOutput]
}
