package usecase

import (
	"context"

	"github.com/paralleltree/go-leaderboard/internal/model"
)

type GetEventsUsecase interface {
	GetEvents(ctx context.Context, page, count int64) ([]model.Record[model.Event], error)
}
