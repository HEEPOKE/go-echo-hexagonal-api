package database

import (
	"context"
	"fmt"
	"log"

	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/config"
	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Cfg.REDIS_URL, config.Cfg.REDIS_PORT),
		Password: config.Cfg.REDIS_PASSWORD,
		DB:       0,
	})
	return client
}

func CheckRedis() {
	redisClient := ConnectToRedis()

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Failed to connect to Redis:", err)
	}
}
