package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/contract/driver"
	"github.com/paralleltree/go-leaderboard/internal/model"
	repositoryImpl "github.com/paralleltree/go-leaderboard/internal/repository"
	mock_driver "github.com/paralleltree/go-leaderboard/mock/driver"
)

func TestScoreRepository_GetScore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sortedSetDriver := mock_driver.NewMockSortedSetDriver(mockCtrl)
	scoringStrategy := mock_driver.NewMockScoringStrategy(mockCtrl)

	scoreRepository := repositoryImpl.NewScoreRepository(scoringStrategy, sortedSetDriver)

	eventId := "1"
	wantKey := fmt.Sprintf("events:%s:scores", eventId)
	wantUserId := "10"
	wantUserScore := int64(10)
	wantComposedScore := int64(10)
	wantOk := true
	ctx := context.Background()

	sortedSetDriver.EXPECT().
		GetScore(ctx, wantKey, wantUserId).
		Return(float64(wantComposedScore), true, nil)
	scoringStrategy.EXPECT().
		ExtractScore(wantComposedScore).
		Return(wantUserScore)

	gotScore, gotOk, err := scoreRepository.GetScore(ctx, eventId, wantUserId)
	if err != nil {
		t.Fatalf("unexpected error while getting score: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
	if wantUserScore != gotScore {
		t.Fatalf("unexpected score: expected %v, but got %v", wantUserScore, gotScore)
	}
}

func TestScoreRepository_SetScore_WhenScoreIsValid_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sortedSetDriver := mock_driver.NewMockSortedSetDriver(mockCtrl)
	scoringStrategy := mock_driver.NewMockScoringStrategy(mockCtrl)

	scoreRepository := repositoryImpl.NewScoreRepository(scoringStrategy, sortedSetDriver)

	eventId := "1"
	wantKey := fmt.Sprintf("events:%s:scores", eventId)
	wantUserId := "10"
	wantUserScore := int64(10)
	wantTime := int64(0)
	wantComposedScore := int64(10)
	ctx := context.Background()

	scoringStrategy.EXPECT().
		ComposeScore(wantTime, wantUserScore).
		Return(wantComposedScore, nil)

	sortedSetDriver.EXPECT().
		SetScore(ctx, wantKey, wantUserId, float64(wantComposedScore)).
		Return(nil)

	if err := scoreRepository.SetScore(ctx, eventId, wantUserId, wantUserScore, wantTime); err != nil {
		t.Fatalf("unexpected error while setting score: %v", err)
	}
}

func TestScoreRepository_GetRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sortedSetDriver := mock_driver.NewMockSortedSetDriver(mockCtrl)
	scoringStrategy := mock_driver.NewMockScoringStrategy(mockCtrl)

	scoreRepository := repositoryImpl.NewScoreRepository(scoringStrategy, sortedSetDriver)

	eventId := "1"
	wantKey := fmt.Sprintf("events:%s:scores", eventId)
	wantUserId := "10"
	wantUserRank := int64(20)
	wantOk := true
	ctx := context.Background()

	sortedSetDriver.EXPECT().
		GetRankByDescending(ctx, wantKey, wantUserId).
		Return(wantUserRank, true, nil)

	gotRank, gotOk, err := scoreRepository.GetRank(ctx, eventId, wantUserId)
	if err != nil {
		t.Fatalf("unexpected error while getting score: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
	if wantUserRank != gotRank {
		t.Fatalf("unexpected score: expected %v, but got %v", wantUserRank, gotRank)
	}
}

func TestScoreRepository_GetLeaderboard(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	sortedSetDriver := mock_driver.NewMockSortedSetDriver(mockCtrl)
	scoringStrategy := mock_driver.NewMockScoringStrategy(mockCtrl)

	scoreRepository := repositoryImpl.NewScoreRepository(scoringStrategy, sortedSetDriver)

	eventId := "1"
	wantKey := fmt.Sprintf("events:%s:scores", eventId)
	wantOk := true
	wantStartRank := int64(10)
	wantEndRank := int64(11)
	wantSortedSetItems := []driver.SortedSetItem{
		{Score: 20, Member: "user1"},
		{Score: 10, Member: "user2"},
	}
	wantRanks := []model.UserRank{
		{Rank: 10, UserId: "user1", Score: 20},
		{Rank: 11, UserId: "user2", Score: 10},
	}
	ctx := context.Background()

	sortedSetDriver.EXPECT().
		GetRankedList(ctx, wantKey, wantStartRank-1, wantEndRank-1).
		Return(wantSortedSetItems, true, nil)
	scoringStrategy.EXPECT().
		ExtractScore(gomock.Any()).AnyTimes().
		DoAndReturn(func(rawScore int64) int64 { return rawScore })

	gotResult, gotOk, err := scoreRepository.GetLeaderboard(ctx, eventId, wantStartRank, wantEndRank)
	if err != nil {
		t.Fatalf("unexpected error while getting score: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
	if len(wantRanks) != len(gotResult) {
		t.Fatalf("unexpected result: expected length %d, but got %d", len(wantRanks), len(gotResult))
	}
	for i, gotItem := range gotResult {
		wantItem := wantRanks[i]
		if wantItem != gotItem {
			t.Fatalf("unexpected result: expected %v, but got %v", wantItem, gotItem)
		}
	}
}
