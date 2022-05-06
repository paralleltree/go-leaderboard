package repository

import (
	"context"
	"fmt"

	"github.com/paralleltree/go-leaderboard/internal/contract/driver"
	"github.com/paralleltree/go-leaderboard/internal/contract/repository"
)

type scoreRepository struct {
	scoringStrategy driver.ScoringStrategy
	sortedSetDriver driver.SortedSetDriver
}

func NewScoreRepository(
	scoringStrategy driver.ScoringStrategy,
	sortedSetDriver driver.SortedSetDriver,
) repository.ScoreRepository {
	return &scoreRepository{
		scoringStrategy: scoringStrategy,
		sortedSetDriver: sortedSetDriver,
	}
}

func (r *scoreRepository) GetScore(ctx context.Context, eventId string, userId string) (int64, bool, error) {
	res, ok, err := r.sortedSetDriver.GetScore(ctx, buildScoreSetKey(eventId), userId)
	if err != nil {
		return 0, false, err
	}
	if !ok {
		return 0, false, nil
	}
	rawScore := int64(res)
	score := r.scoringStrategy.ExtractScore(rawScore)
	return score, true, nil
}

func (r *scoreRepository) SetScore(ctx context.Context, eventId, userId string, score int64, time int64) error {
	rawScore, err := r.scoringStrategy.ComposeScore(time, score)
	if err != nil {
		return fmt.Errorf("composing score: %w", err)
	}
	return r.sortedSetDriver.SetScore(ctx, buildScoreSetKey(eventId), userId, float64(rawScore))
}

func (r *scoreRepository) GetRank(ctx context.Context, eventId string, userId string) (int64, bool, error) {
	rank, ok, err := r.sortedSetDriver.GetRankByDescending(ctx, buildScoreSetKey(eventId), userId)
	if err != nil {
		return 0, false, err
	}
	if !ok {
		return 0, false, nil
	}
	return rank, true, nil
}

func (r *scoreRepository) GetLeaderboard(ctx context.Context, eventId string, startRank, endRank int64) ([]repository.UserRank, bool, error) {
	items, ok, err := r.sortedSetDriver.GetRankedList(ctx, buildScoreSetKey(eventId), startRank-1, endRank-1)
	if err != nil {
		return nil, false, err
	}
	if !ok {
		return nil, false, nil
	}
	result := make([]repository.UserRank, 0, len(items))
	for i, item := range items {
		item := repository.UserRank{
			Rank:   startRank + int64(i),
			UserId: item.Member,
			Score:  r.scoringStrategy.ExtractScore(int64(item.Score)),
		}
		result = append(result, item)
	}
	return result, true, nil
}

func buildScoreSetKey(eventId string) string {
	return fmt.Sprintf("events:%s:scores", eventId)
}
