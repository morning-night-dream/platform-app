package gateway

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
	"github.com/morning-night-dream/platform-app/pkg/ent"
)

var _ repository.User = (*User)(nil)

type User struct {
	db *ent.Client
}

func NewUser(db *ent.Client) *User {
	return &User{
		db: db,
	}
}

func (us *User) Save(ctx context.Context, item model.User) error {
	if _, err := us.db.User.Create().SetID(uuid.MustParse(string(item.UserID))).Save(ctx); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}
