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
	scoreRepository := mock_repository.NewMockScoreRepository(mockCtrl)

	wantEventId := "1"
	wantStartRank := int64(1)
	wantEndRank := int64(10)
	wantRanks := []model.UserRank{
		{Rank: 1, UserId: "1", Score: 42},
	}
	wantOk := true
	ctx := context.Background()

	scoreRepository.EXPECT().
		GetLeaderboard(ctx, wantEventId, wantStartRank, wantEndRank).
		Return(wantRanks, wantOk, nil)

	usecase := usecase.NewGetLeaderboardUsecase(scoreRepository)
	gotRanks, gotOk, err := usecase.GetLeaderboard(ctx, wantEventId, wantStartRank, wantEndRank)
	if err != nil {
		t.Fatalf("unexpected result(error): %v", err)
	}
	if gotOk != wantOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
	if wantOk {
		for i, wantRank := range gotRanks {
			if wantRank != gotRanks[i] {
				t.Fatalf("unexpected result[%d]: expected %v, but got %v", i, wantRank, gotRanks[i])
			}
		}
	}
}
