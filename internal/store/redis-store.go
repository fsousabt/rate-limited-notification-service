package store

import (
	"context"
	"encoding/json"

	"github.com/fsousabt/rate-limiter/infra/cache"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore(addr string) *RedisStore {
	return &RedisStore{
		client: cache.NewRedisClient(addr),
		ctx:    context.Background(),
	}
}

func (s *RedisStore) Get(key string) (Bucket, bool) {
	val, err := s.client.Get(s.ctx, key).Result()
	if err == redis.Nil {
		return Bucket{}, false
	} else if err != nil {
		return Bucket{}, false
	}

	bucket, err := deserializeBucket(val)
	if err != nil {
		return Bucket{}, false
	}

	return bucket, true
}

func (s *RedisStore) Set(key string, item Bucket) bool {
	for {
		err := s.client.Watch(s.ctx, func(tx *redis.Tx) error {
			val, err := serializeBucket(item)
			if err != nil {
				return err
			}

			_, err = tx.TxPipelined(s.ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(s.ctx, key, val, 0)
				return nil
			})
			return err
		}, key)

		if err == redis.TxFailedErr {
			continue
		}
		return err == nil
	}
}

func serializeBucket(b Bucket) (string, error) {
	data, err := json.Marshal(b)
	return string(data), err
}

func deserializeBucket(s string) (Bucket, error) {
	var b Bucket
	err := json.Unmarshal([]byte(s), &b)
	return b, err
}
