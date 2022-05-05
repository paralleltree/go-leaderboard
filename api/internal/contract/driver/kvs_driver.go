//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package driver

import "context"

type HashDriver interface {
	Get(ctx context.Context, key string) (map[string]string, bool, error)
	Set(ctx context.Context, key string, fields map[string]string) error
}

type SortedSetItem struct {
	Score  float64
	Member string
}

type SortedSetDriver interface {
	GetScore(ctx context.Context, key, member string) (float64, bool, error)
	SetScore(ctx context.Context, key, member string, score float64) error
	GetRankByDescending(ctx context.Context, key string, member string) (int64, bool, error)
	GetRankedList(ctx context.Context, key string, start, stop int64) ([]SortedSetItem, bool, error)
}
