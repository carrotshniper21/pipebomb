package cache

import (
	"context"
	"encoding/json"
	redis "github.com/go-redis/redis/v8"
	"time"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func cacheData(key string, fetchFunc func(string) (interface{}, error), query string) (interface{}, error) {
	ctx := context.Background()
	cachedData, err := redisClient.Get(ctx, key).Result()
	if err == nil {
		var data interface{}
		err = json.Unmarshal([]byte(cachedData), &data)
		if err == nil {
			return data, nil
		}
	}

	data, err := fetchFunc(query)
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(data)
	if err == nil {
		redisClient.Set(ctx, key, dataBytes, 10*time.Minute)
	}

	return data, err
}
