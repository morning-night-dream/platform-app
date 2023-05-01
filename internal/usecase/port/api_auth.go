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
	model.ExpiresIn
}

type APIAuthSignInOutput struct {
	usecase.Output
	UserID string
	model.SessionID
	model.IDToken
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
	model.IDToken
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
	model.IDToken
	model.SessionToken
}

type APIAuthSignOutOutput struct {
	usecase.Output
}

type APIAuthSignOut interface {
	usecase.Usecase[APIAuthSignOutInput, APIAuthSignOutOutput]
}

type APIAuthGenerateCodeInput struct {
	model.SessionToken
}

type APIAuthGenerateCodeOutput struct {
	model.CodeID
}

type APIAuthGenerateCode interface {
	usecase.Usecase[APIAuthGenerateCodeInput, APIAuthGenerateCodeOutput]
}
