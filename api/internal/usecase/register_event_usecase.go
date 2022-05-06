package usecase

import (
	"context"
	"fmt"

	"github.com/paralleltree/go-leaderboard/internal/contract/repository"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
	"github.com/paralleltree/go-leaderboard/internal/model"
)

type registerEventUsecase struct {
	eventRepository repository.EventRepository
}

func NewRegisterEventUsecase(eventRepository repository.EventRepository) usecase.RegisterEventUsecase {
	return &registerEventUsecase{
		eventRepository: eventRepository,
	}
}

func (u *registerEventUsecase) RegisterEvent(ctx context.Context, event model.Event) (string, error) {
	id, err := u.eventRepository.RegisterEvent(ctx, event)
	if err != nil {
		return "", fmt.Errorf("register event: %w", err)
	}
	return id, nil
}
