package repository

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type APISession interface {
	Save(context.Context, model.Session) error
	Find(context.Context, model.SessionID) (model.Session, error)
	Delete(context.Context, model.SessionID) error
}
