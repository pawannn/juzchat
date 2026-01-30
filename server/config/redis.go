package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.DBAddr,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	return rdb
}
