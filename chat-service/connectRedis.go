package chat

import (
	"chat-app/util"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	ctx         context.Context
	config      util.Config
)

func ConnectRedis(config *util.Config) {
	ctx = context.TODO()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err := RedisClient.Set(ctx, "test", "How to Refresh Access Tokens the Right Way in Golang", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Redis client connected successfully...")
}
