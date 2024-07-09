package storage

import (
	"fmt"
	"task_manager/internal/config"

	"github.com/go-redis/redis"
)

type RedisManager struct {
	rdb *redis.Client
}

func NewRedisManager() *RedisManager {
	addr := fmt.Sprintf("%s:%s",
		config.GetStrEnv("REDIS_HOST", "localhost"),
		config.GetStrEnv("REDIS_PORT", "6379"))
	return &RedisManager{
		rdb: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.GetStrEnv("REDIS_PASSWORD", ""),
			DB:       config.GetIntEnv("REDIS_DB", 0),
		}),
	}
}

func makeCustomerShippingId(id string) string {
	return fmt.Sprintf("customer_shipping:%s", id)
}

func makeCustomerShippingAddressId(id string) string {
	return fmt.Sprintf("customer_shipping_address:%s", id)
}
