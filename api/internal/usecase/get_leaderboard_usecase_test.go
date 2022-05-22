package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/model"
	"github.com/paralleltree/go-leaderboard/internal/usecase"
	mock_repository "github.com/paralleltree/go-leaderboard/mock/repository"
)

func TestGetLeaderboardUsecase_GetLeaderboard_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	eventRepository := mock_repository.NewMockEventRepository(mockCtrl)
	scoreRepository := mock_repository.NewMockScoreRepository(mockCtrl)

	wantEventId := "1"
	wantStartRank := int64(1)
	wantEndRank := int64(10)
	wantRanks := []model.UserRank{
		{Rank: 1, UserId: "1", Score: 42},
	}
	wantOk := true
	ctx := context.Background()

	eventRepository.EXPECT().
		GetEvent(ctx, wantEventId).
		Return(model.Event{}, true, nil)
	scoreRepository.EXPECT().
		GetLeaderboard(ctx, wantEventId, wantStartRank, wantEndRank).
		Return(wantRanks, wantOk, nil)

	usecase := usecase.NewGetLeaderboardUsecase(eventRepository, scoreRepository)
	gotRanks, gotOk, err := usecase.GetLeaderboard(ctx, wantEventId, wantStartRank, wantEndRank)
	if err != nil {
		t.Fatalf("unexpected result(error): %v", err)
	}
	if gotOk != wantOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
	if len(wantRanks) != len(gotRanks) {
		t.Fatalf("unexpected result count: expected %v, but got %v", len(wantRanks), len(gotRanks))
	}
	if wantOk {
		for i, wantRank := range wantRanks {
			if wantRank != gotRanks[i] {
				t.Fatalf("unexpected result[%d]: expected %v, but got %v", i, wantRank, gotRanks[i])
			}
		}
	}
}

func TestGetLeaderboardUsecase_GetLeaderboard_WhenNoScoreExists_ReturnsEmpty(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	eventRepository := mock_repository.NewMockEventRepository(mockCtrl)
	scoreRepository := mock_repository.NewMockScoreRepository(mockCtrl)

	wantEventId := "1"
	wantStartRank := int64(1)
	wantEndRank := int64(10)
	wantRanks := []model.UserRank{}
	wantOk := true
	ctx := context.Background()

	eventRepository.EXPECT().
		GetEvent(ctx, wantEventId).
		Return(model.Event{}, true, nil)
	scoreRepository.EXPECT().
		GetLeaderboard(ctx, wantEventId, wantStartRank, wantEndRank).
		Return(nil, false, nil)

	usecase := usecase.NewGetLeaderboardUsecase(eventRepository, scoreRepository)
	gotRanks, gotOk, err := usecase.GetLeaderboard(ctx, wantEventId, wantStartRank, wantEndRank)
	if err != nil {
		t.Fatalf("unexpected result(error): %v", err)
	}
	if gotOk != wantOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
	if len(wantRanks) != len(gotRanks) {
		t.Fatalf("unexpected result count: expected %v, but got %v", len(wantRanks), len(gotRanks))
	}
	for i, wantRank := range wantRanks {
		if wantRank != gotRanks[i] {
			t.Fatalf("unexpected result[%d]: expected %v, but got %v", i, wantRank, gotRanks[i])
		}
	}
}

func TestGetLeaderboardUsecase_GetLeaderboard_WhenEventDoesNotExists_ReturnsNotOK(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	eventRepository := mock_repository.NewMockEventRepository(mockCtrl)
	scoreRepository := mock_repository.NewMockScoreRepository(mockCtrl)

	wantEventId := "1"
	wantStartRank := int64(1)
	wantEndRank := int64(10)
	wantOk := false
	ctx := context.Background()

	eventRepository.EXPECT().
		GetEvent(ctx, wantEventId).
		Return(model.Event{}, false, nil)

	usecase := usecase.NewGetLeaderboardUsecase(eventRepository, scoreRepository)
	_, gotOk, err := usecase.GetLeaderboard(ctx, wantEventId, wantStartRank, wantEndRank)
	if err != nil {
		t.Fatalf("unexpected result(error): %v", err)
	}
	if gotOk != wantOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
}
