package repository

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type APIAuth interface {
	SignUp(context.Context, model.UserID, model.EMail, model.Password) error
	SignIn(context.Context, model.EMail, model.Password) (model.Auth, error)
	Refresh(context.Context, model.RefreshToken) (model.RefreshToken, error)
	Verify(context.Context) error
	Delete(context.Context, model.UserID) error
	Save(context.Context, model.Auth) error
	Find(context.Context, model.UserID) (model.Auth, error)
}
