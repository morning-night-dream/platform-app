package public

import (
	"context"
	"sync"
	"time"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type Public struct {
	lock sync.Mutex
	key  map[string]string
}

type Cache struct {
	model.Auth
	CreatedAt time.Time
}

func New() *Public {
	return &Public{
		key: make(map[string]string),
	}
}

func (pub *Public) Get(ctx context.Context, key string) (string, error) {
	pub.lock.Lock()
	defer pub.lock.Unlock()

	// && val.CreatedAt.Before(time.Now().Add(time.Duration(val.ExpiresIn)))

	if val, ok := pub.key[key]; ok {
		return val, nil
	}

	return "", nil
}

func (pub *Public) Set(ctx context.Context, key string, val string) error {
	pub.lock.Lock()
	defer pub.lock.Unlock()

	pub.key[key] = val

	return nil
}

func (pub *Public) Delete(ctx context.Context, key string) error {
	pub.lock.Lock()
	defer pub.lock.Unlock()

	delete(pub.key, key)

	return nil
}
