package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/model"
	"github.com/paralleltree/go-leaderboard/internal/usecase"
	mock_repository "github.com/paralleltree/go-leaderboard/mock/repository"
)

func TestSetScoreUsecase_SetScore(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	eventRepository := mock_repository.NewMockEventRepository(mockCtrl)
	scoreRepository := mock_repository.NewMockScoreRepository(mockCtrl)
	wantCurrentTime := time.Date(2022, 5, 1, 13, 0, 0, 0, time.UTC)
	timeProvider := func() time.Time { return wantCurrentTime }

	wantEventId := "1"
	event := model.Event{
		Name:    "test",
		StartAt: time.Date(2022, 5, 1, 12, 0, 0, 0, time.UTC),
		EndAt:   time.Date(2022, 5, 1, 14, 0, 0, 0, time.UTC),
	}
	wantUserId := "2"
	wantScore := int64(10)
	ctx := context.Background()

	eventRepository.EXPECT().
		GetEvent(ctx, wantEventId).
		Return(event, true, nil)
	scoreRepository.EXPECT().
		SetScore(ctx, wantEventId, wantUserId, wantScore, int64(event.EndAt.Sub(wantCurrentTime).Seconds())).
		Return(nil)

	usecase := usecase.NewSetScoreUsecase(eventRepository, scoreRepository, timeProvider)

	if err := usecase.SetScore(ctx, wantEventId, wantUserId, wantScore); err != nil {
		t.Fatalf("unexpected error while setting score: %v", err)
	}
}
