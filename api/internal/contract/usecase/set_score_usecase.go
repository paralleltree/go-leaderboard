package usecase

import "context"

type SetScoreUsecase interface {
	SetScore(ctx context.Context, eventId string, userId string, score int64) error
}
