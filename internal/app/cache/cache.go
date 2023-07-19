package cache

import (
	"context"
	"time"

	"github.com/HEEPOKE/go-echo-hexagonal-api/pkg/database"
)

func SetKey(key string, value interface{}, expiration time.Duration) error {
	client := database.ConnectToRedis()
	return client.Set(context.Background(), key, value, expiration).Err()
}
func GetKey(key string) (string, error) {
	client := database.ConnectToRedis()
	return client.Get(context.Background(), key).Result()
}
