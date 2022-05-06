package repository

import "context"

type UserRank struct {
	Rank   int64
	UserId string
	Score  int64
}

type ScoreRepository interface {
	GetScore(ctx context.Context, eventId string, userId string) (int64, bool, error)
	SetScore(ctx context.Context, eventId string, userId string, score int64, time int64) error
	GetRank(ctx context.Context, eventId string, userId string) (int64, bool, error)
	GetLeaderboard(ctx context.Context, eventId string, startRank, endRank int64) ([]UserRank, bool, error)
}
