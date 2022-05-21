//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/paralleltree/go-leaderboard/internal/model"
)

type ScoreRepository interface {
	GetScore(ctx context.Context, eventId string, userId string) (int64, bool, error)
	SetScore(ctx context.Context, eventId string, userId string, score int64, time int64) error
	GetRank(ctx context.Context, eventId string, userId string) (int64, bool, error)
	GetLeaderboard(ctx context.Context, eventId string, startRank, endRank int64) ([]model.UserRank, bool, error)
}
