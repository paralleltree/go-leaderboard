package usecase

import "context"

type GetRankUsecase interface {
	GetRank(ctx context.Context, eventId string, userId string) (int64, bool, error)
}
