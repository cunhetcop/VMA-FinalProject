// utils/redis.go
package utils

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), 
		Password: os.Getenv("REDIS_PASSWORD"), 
		DB:       0,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis")
	}
}

func DeleteKey(key string) error {
	err := RedisClient.Del(context.Background(), key).Err()
	return err
}

func FlushAll() error {
	err := RedisClient.FlushAll(context.Background()).Err()
	return err
}