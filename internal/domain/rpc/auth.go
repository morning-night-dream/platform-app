package rpc

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type Auth interface {
	SignUp(context.Context, model.UserID, model.EMail, model.Password) error
	SignIn(context.Context, model.EMail, model.Password) (model.Auth, error)
}
