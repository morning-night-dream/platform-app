package redis

import (
	"context"

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
	client *redis.Client,
) (*kvs.KVS[T], error) {
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return &kvs.KVS[T]{
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
