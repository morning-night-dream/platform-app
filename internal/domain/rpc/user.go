package rpc

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type User interface {
	SignUp(context.Context) (model.User, error)
}
