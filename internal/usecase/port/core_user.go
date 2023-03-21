package port

import (
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/usecase"
)

type CoreUserSignUpInput struct {
	usecase.Input
}

type CoreUserSignUpOutput struct {
	usecase.Output
	model.User
}

type CoreUserSignUp interface {
	usecase.Usecase[CoreUserSignUpInput, CoreUserSignUpOutput]
}
