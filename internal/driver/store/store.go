package store

import (
	"context"
	"sync"
	"time"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
)

type Store struct {
	lock  sync.Mutex
	cache map[string]Cache
}

const ttl = time.Second

type Cache struct {
	model.Auth
	CreatedAt time.Time
}

func New() *Store {
	return &Store{
		cache: make(map[string]Cache),
	}
}

func (store *Store) Get(ctx context.Context, key string) (model.Auth, error) {
	store.lock.Lock()
	defer store.lock.Unlock()

	// && val.CreatedAt.Before(time.Now().Add(time.Duration(val.ExpiresIn)))

	if val, ok := store.cache[key]; ok {
		return val.Auth, nil
	}

	return model.Auth{}, nil
}

func (store *Store) Set(ctx context.Context, key string, val model.Auth) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	store.cache[key] = Cache{
		Auth:      val,
		CreatedAt: time.Now(),
	}

	return nil
}

func (store *Store) Delete(ctx context.Context, key string) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	delete(store.cache, key)

	return nil
}
