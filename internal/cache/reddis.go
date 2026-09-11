package cache

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
}
