package driver

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/paralleltree/go-leaderboard/internal/contract/driver"
)

type redisHashDriver struct {
	client *redis.Client
}

func NewRedisHashDriver(endpoint string) driver.HashDriver {
	return &redisHashDriver{
		client: newRedisClient(endpoint),
	}
}

func (d *redisHashDriver) Get(ctx context.Context, key string) (map[string]string, bool, error) {
	count, err := d.client.Exists(ctx, key).Result()
	if err != nil {
		return nil, false, err
	}
	if count == 0 {
		return nil, false, nil
	}
	res, err := d.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (d *redisHashDriver) Set(ctx context.Context, key string, fields map[string]string) error {
	return d.client.HSet(ctx, key, fields).Err()
}
