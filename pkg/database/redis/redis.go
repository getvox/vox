package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func NewClient(c *Config) (*redis.Client, error) {
	rc := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	if _, err := rc.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return rc, nil
}
