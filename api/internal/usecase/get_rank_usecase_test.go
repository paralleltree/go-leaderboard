package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/usecase"
	mock_repository "github.com/paralleltree/go-leaderboard/mock/repository"
)

func TestGetRankUsecase_GetRank(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	scoreRepository := mock_repository.NewMockScoreRepository(mockCtrl)

	eventId := "1"
	userId := "2"
	wantRank := int64(1)
	wantOk := true
	ctx := context.Background()

	scoreRepository.EXPECT().
		GetRank(ctx, eventId, userId).
		Return(wantRank, true, nil)

	usecase := usecase.NewGetRankUsecase(scoreRepository)
	gotRank, gotOk, err := usecase.GetRank(ctx, eventId, userId)
	if err != nil {
		t.Fatalf("unexpected error while getting rank: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
	if wantRank != gotRank {
		t.Fatalf("unexpected result: expected rank %v, but got %v", wantRank, gotRank)
	}
}
