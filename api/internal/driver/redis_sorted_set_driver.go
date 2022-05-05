package driver

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/paralleltree/go-leaderboard/internal/contract/driver"
)

type redisSortedSetDriver struct {
	client *redis.Client
}

func NewRedisSortedSetDriver(endpoint string) driver.SortedSetDriver {
	return &redisSortedSetDriver{
		client: newRedisClient(endpoint),
	}
}

func (d *redisSortedSetDriver) GetScore(ctx context.Context, key, member string) (float64, bool, error) {
	res := d.client.ZScore(ctx, key, member)
	if err := res.Err(); err != nil {
		if err == redis.Nil {
			// key or member not found
			return 0, false, nil
		}
		return 0, false, err
	}
	return res.Val(), true, nil
}

func (d *redisSortedSetDriver) SetScore(ctx context.Context, key, member string, score float64) error {
	return d.client.ZAdd(ctx, key, &redis.Z{Score: score, Member: member}).Err()
}

func (d *redisSortedSetDriver) GetRankByDescending(ctx context.Context, key, member string) (int64, bool, error) {
	res, err := d.client.ZRevRank(ctx, key, member).Result()
	if err == redis.Nil {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}
	return res, true, nil
}

func (d *redisSortedSetDriver) GetRankedList(ctx context.Context, key string, start, stop int64) ([]driver.SortedSetItem, bool, error) {
	count, err := d.client.Exists(ctx, key).Result()
	if err != nil {
		return nil, false, err
	}
	if count == 0 {
		return nil, false, nil
	}
	res, err := d.client.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil, false, err
	}
	members := make([]driver.SortedSetItem, 0, len(res))
	for _, m := range res {
		item := driver.SortedSetItem{
			Score:  m.Score,
			Member: m.Member.(string),
		}
		members = append(members, item)
	}
	return members, true, nil
}
