package module

import (
	"fmt"

	"github.com/go-redis/redis/v9"
)

func NewRedis(config *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		DB:   config.Redis.DB,
	})
}
