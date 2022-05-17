package usecase

import (
	"context"
	"fmt"

	"github.com/paralleltree/go-leaderboard/internal/contract/repository"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
)

type getRankUsecase struct {
	scoreRepository repository.ScoreRepository
}

func NewGetRankUsecase(scoreRepository repository.ScoreRepository) usecase.GetRankUsecase {
	return &getRankUsecase{
		scoreRepository: scoreRepository,
	}
}

func (u *getRankUsecase) GetRank(ctx context.Context, eventId string, userId string) (int64, bool, error) {
	rank, ok, err := u.scoreRepository.GetRank(ctx, eventId, userId)
	if err != nil {
		return 0, false, fmt.Errorf("get rank: %w", err)
	}
	if !ok {
		return 0, false, nil
	}
	return rank, true, nil
}
