package usecase

import (
	"context"
	"fmt"

	"github.com/paralleltree/go-leaderboard/internal/contract/repository"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
	"github.com/paralleltree/go-leaderboard/internal/model"
)

type getLeaderboardUsecase struct {
	eventRepository repository.EventRepository
	scoreRepository repository.ScoreRepository
}

func NewGetLeaderboardUsecase(eventRepository repository.EventRepository, scoreRepository repository.ScoreRepository) usecase.GetLeaderboardUsecase {
	return &getLeaderboardUsecase{
		eventRepository: eventRepository,
		scoreRepository: scoreRepository,
	}
}

func (u *getLeaderboardUsecase) GetLeaderboard(ctx context.Context, eventId string, startRank, endRank int64) ([]model.UserRank, bool, error) {
	_, ok, err := u.eventRepository.GetEvent(ctx, eventId)
	if err != nil {
		return nil, false, fmt.Errorf("get event: %w", err)
	}
	if !ok {
		return nil, false, nil
	}
	ranks, ok, err := u.scoreRepository.GetLeaderboard(ctx, eventId, startRank, endRank)
	if err != nil {
		return nil, false, fmt.Errorf("get leaderboard: %w", err)
	}
	if !ok {
		return nil, true, nil
	}
	return ranks, true, nil
}
