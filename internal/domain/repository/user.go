package repository

import (
	"context"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type User interface {
	Save(context.Context, model.User) error
}
