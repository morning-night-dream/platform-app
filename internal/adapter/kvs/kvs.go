package kvs

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/morning-night-dream/platform-app/internal/domain/cache"
	"github.com/redis/go-redis/v9"
)

type KVSFactory[T any] interface {
	Of(*redis.Client) (*KVS[T], error)
}

var _ cache.Cache[any] = (*KVS[any])(nil)

type KVS[T any] struct {
	Client *redis.Client
}

func (kvs *KVS[T]) Get(ctx context.Context, key string) (*T, error) {
	var value T

	str, err := kvs.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, errors.New("1 failed to get cache")
	}

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		return nil, errors.New("failed to unmarshal json")
	}

	return nil, nil
}

func (kvs *KVS[T]) Set(ctx context.Context, key string, value *T, ttl time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return errors.New("failed to marshal json")
	}

	if err := kvs.Client.Set(ctx, key, string(val), ttl).Err(); err != nil {
		return errors.New("failed to set redis")
	}

	return nil
}

func (kvs *KVS[T]) Del(ctx context.Context, key string) error {
	kvs.Client.Del(ctx, key)

	return nil
}

// gob を試した形跡。うまく動かせなかった。。。

// func (kvs *KVS[T]) encode(value T) ([]byte, error) {
// 	buf := bytes.NewBuffer(nil)

// 	if err := gob.NewEncoder(buf).Encode(&value); err != nil {
// 		return nil, err
// 	}

// 	return buf.Bytes(), nil
// }

// func (kvs *KVS[T]) decode(data []byte) (T, error) {
// 	var value T

// 	buf := bytes.NewBuffer(data)

// 	if err := gob.NewDecoder(buf).Decode(&value); err != nil {
// 		return value, err
// 	}

// 	return value, nil
// }
