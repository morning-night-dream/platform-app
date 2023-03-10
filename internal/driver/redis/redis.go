package redis

import (
	"context"

	"github.com/dgraph-io/ristretto"
	"github.com/morning-night-dream/platform-app/internal/adapter/kvs"
	"github.com/morning-night-dream/platform-app/internal/driver/env"
	"github.com/redis/go-redis/v9"
)

var _ kvs.KVSFactory[any] = (*Redis[any])(nil)

type Redis[T any] struct{}

func New[T any]() *Redis[T] {
	return &Redis[T]{}
}

func (rds *Redis[T]) Of(
	cache *ristretto.Cache,
	client *redis.Client,
) (*kvs.KVS[T], error) {
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return &kvs.KVS[T]{
		Cache:  cache,
		Client: client,
	}, nil
}

func NewRedis(url string) *redis.Client {
	var opt *redis.Options
	if env.Env.IsProd() {
		var err error
		opt, err = redis.ParseURL(url)
		if err != nil {
			panic(err)
		}
	} else {
		opt = &redis.Options{
			Addr:     url,
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	}

	return redis.NewClient(opt)
}

func NewCache() (*ristretto.Cache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, err
	}

	return cache, nil
}
