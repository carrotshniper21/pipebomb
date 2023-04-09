// cache/redis.go
package cache

import (
	"context"
	"encoding/json"

	redis "github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	redisClient *redis.Client
	Addr        string
	Password    string
	DB          int
}

func NewRedisCache(addr string, password string, db int) *RedisCache {
	return &RedisCache{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	ctx := context.Background()
	cachedData, err := r.redisClient.Get(ctx, key).Result()
	if err == nil {
		var data interface{}
		err = json.Unmarshal([]byte(cachedData), &data)
		if err == nil {
			return data, nil
		}
	}

	return nil, err
}

func (r *RedisCache) Set(key string, data interface{}, expiration time.Duration) error {
	ctx := context.Background()
	dataBytes, err := json.Marshal(data)
	if err == nil {
		err = r.redisClient.Set(ctx, key, dataBytes, expiration).Err()
	}

	return err
}

func (r *RedisCache) Delete(key string) error {
	ctx := context.Background()
	return r.redisClient.Del(ctx, key).Err()
}

func (r *RedisCache) Close() error {
	return r.redisClient.Close()
}

func NewCache(addr string, password string, db int) *RedisCache {
	return NewRedisCache(addr, password, db)
}

func CacheData(r *RedisCache, key string, fetchFunc func(string) (interface{}, error), query string) (interface{}, error) {
	data, err := r.Get(key)
	if err == nil {
		return data, nil
	}

	data, err = fetchFunc(query)
	if err != nil {
		return nil, err
	}

	err = r.Set(key, data, 10*time.Minute)
	if err != nil {
		return nil, err
	}

	return data, err
}
