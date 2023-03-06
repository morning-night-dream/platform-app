package repository

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type APICode interface {
	Save(context.Context, model.Code) error
	Find(context.Context, model.CodeID) (model.Code, error)
	Delete(context.Context, model.CodeID) error
}
