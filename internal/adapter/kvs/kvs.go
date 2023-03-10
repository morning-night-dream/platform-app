package kvs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/morning-night-dream/platform-app/internal/domain/cache"
	"github.com/morning-night-dream/platform-app/pkg/log"
	"github.com/redis/go-redis/v9"
)

type KVSFactory[T any] interface {
	Of(*ristretto.Cache, *redis.Client) (*KVS[T], error)
}

var _ cache.Cache[any] = (*KVS[any])(nil)

type KVS[T any] struct {
	Cache  *ristretto.Cache
	Client *redis.Client
}

func (kvs *KVS[T]) Get(ctx context.Context, key string) (T, error) {
	var value T

	// インメモリキャッシュから取得した値は[]byte型になっている

	val, ok := kvs.Cache.Get(key)
	if !ok {
		log.Log().Info(fmt.Sprintf("miss cache key: %s", key))
	}

	log.GetLogCtx(ctx).Info(fmt.Sprintf("cache value: %T: %+v", val, val))

	if _, ok := val.(T); ok {
		return val.(T), nil
	}

	// redisから取得した値は文字列になっている

	str, err := kvs.Client.Get(ctx, key).Result()
	if err != nil {
		return value, errors.New("1 failed to get cache")
	}

	log.GetLogCtx(ctx).Info(fmt.Sprintf("cache value: %T: %+v", val, val))

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		return value, errors.New("failed to unmarshal json")
	}

	return value, nil
}

func (kvs *KVS[T]) Set(ctx context.Context, key string, value T, ttl time.Duration) error {
	if ok := kvs.Cache.SetWithTTL(key, value, 1, ttl); !ok {
		return errors.New("failed to set in memory")
	}

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
	kvs.Cache.Del(key)

	kvs.Client.Del(ctx, key)

	return nil
}

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
