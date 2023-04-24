// utils/scheduler.go
package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

var ctx = context.Background()
var redisClient *redis.Client

func deleteExpiredTokens() {

	keys, err := redisClient.Keys(ctx, "*").Result()
	if err != nil {
		fmt.Println("Error fetching keys:", err)
		return
	}

	for _, key := range keys {
		ttl, err := redisClient.TTL(ctx, key).Result()
		if err != nil {
			fmt.Println("Error fetching key TTL:", err)
			continue
		}

		if ttl <= 0 {
			err = redisClient.Del(ctx, key).Err()
			if err != nil {
				fmt.Println("Error deleting key:", err)
				continue
			}
			fmt.Printf("Deleted expired token: %s\n", key)
		}
	}
}

func ScheduleDeleteExpiredTokens() {
	c := cron.New(cron.WithLocation(time.FixedZone("Bangkok", 7*60*60)))
	_, err := c.AddFunc("0 17 * * 6", deleteExpiredTokens)
	if err != nil {
		fmt.Println("Error scheduling deleteExpiredTokens:", err)
		return
	}
	c.Start()
}
