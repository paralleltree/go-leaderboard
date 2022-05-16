package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/paralleltree/go-leaderboard/internal/contract/repository"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
)

type TimeProvider func() time.Time

type setScoreUsecase struct {
	eventRepository     repository.EventRepository
	scoreRepository     repository.ScoreRepository
	currentTimeProvider TimeProvider
}

func NewSetScoreUsecase(eventRepository repository.EventRepository, scoreRepository repository.ScoreRepository, currentTimeProvider TimeProvider) usecase.SetScoreUsecase {
	return &setScoreUsecase{
		eventRepository:     eventRepository,
		scoreRepository:     scoreRepository,
		currentTimeProvider: currentTimeProvider,
	}
}

func (u *setScoreUsecase) SetScore(ctx context.Context, eventId string, userId string, score int64) error {
	now := u.currentTimeProvider()
	event, ok, err := u.eventRepository.GetEvent(ctx, eventId)
	if err != nil {
		return fmt.Errorf("get event: %w", err)
	}
	if !ok {
		return fmt.Errorf("get event: event not found: %s", eventId)
	}
	remainingTime := int64(event.EndAt.Sub(now).Seconds())
	if err := u.scoreRepository.SetScore(ctx, eventId, userId, score, remainingTime); err != nil {
		return fmt.Errorf("set score: %w", err)
	}
	return nil
}
