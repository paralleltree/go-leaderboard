package driver

import "github.com/go-redis/redis/v8"

func newRedisClient(endpoint string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: "",
		DB:       0,
	})
}
