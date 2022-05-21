package usecase

import (
	"context"
	"fmt"

	"github.com/paralleltree/go-leaderboard/internal/contract/repository"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
	"github.com/paralleltree/go-leaderboard/internal/model"
)

type getLeaderboardUsecase struct {
	scoreRepository repository.ScoreRepository
}

func NewGetLeaderboardUsecase(scoreRepository repository.ScoreRepository) usecase.GetLeaderboardUsecase {
	return &getLeaderboardUsecase{
		scoreRepository: scoreRepository,
	}
}

func (u *getLeaderboardUsecase) GetLeaderboard(ctx context.Context, eventId string, startRank, endRank int64) ([]model.UserRank, bool, error) {
	ranks, ok, err := u.scoreRepository.GetLeaderboard(ctx, eventId, startRank, endRank)
	if err != nil {
		return nil, false, fmt.Errorf("get leaderboard: %w", err)
	}
	if !ok {
		return nil, false, nil
	}
	return ranks, true, nil
}
