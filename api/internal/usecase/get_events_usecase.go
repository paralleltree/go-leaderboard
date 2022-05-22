package usecase

import (
	"context"
	"fmt"

	"github.com/paralleltree/go-leaderboard/internal/contract/repository"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
	"github.com/paralleltree/go-leaderboard/internal/model"
)

type getEventsUsecase struct {
	eventRepository repository.EventRepository
}

func NewGetEventsUsecase(eventRepository repository.EventRepository) usecase.GetEventsUsecase {
	return &getEventsUsecase{
		eventRepository: eventRepository,
	}
}

func (u *getEventsUsecase) GetEvents(ctx context.Context, page, count int64) ([]model.Record[model.Event], error) {
	res, err := u.eventRepository.GetEvents(ctx, page, count)
	if err != nil {
		return nil, fmt.Errorf("get event: %w", err)
	}
	return res, nil
}
