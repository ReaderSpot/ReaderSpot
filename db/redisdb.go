package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var redisClient *redis.Client

func RedisClient() *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	}
	return redisClient
}
