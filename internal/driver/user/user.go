package user

import (
	"context"
	"sync"
	"time"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type User struct {
	lock sync.Mutex
	key  map[string]string
}

type Cache struct {
	model.Auth
	CreatedAt time.Time
}

func New() *User {
	return &User{
		key: make(map[string]string),
	}
}

func (user *User) Get(ctx context.Context, key string) (string, error) {
	user.lock.Lock()
	defer user.lock.Unlock()

	// && val.CreatedAt.Before(time.Now().Add(time.Duration(val.ExpiresIn)))

	if val, ok := user.key[key]; ok {
		return val, nil
	}

	return "", nil
}

func (user *User) Set(ctx context.Context, key string, val string) error {
	user.lock.Lock()
	defer user.lock.Unlock()

	user.key[key] = val

	return nil
}

func (user *User) Delete(ctx context.Context, key string) error {
	user.lock.Lock()
	defer user.lock.Unlock()

	delete(user.key, key)

	return nil
}
