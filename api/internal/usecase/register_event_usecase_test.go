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

func TestRegisterEventUsecase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	eventRepository := mock_repository.NewMockEventRepository(mockCtrl)

	usecase := usecase.NewRegisterEventUsecase(eventRepository)

	event := model.Event{
		Name:    "test",
		StartAt: time.Date(2022, 5, 1, 12, 0, 0, 0, time.UTC),
		EndAt:   time.Date(2022, 5, 1, 13, 0, 0, 0, time.UTC),
	}
	wantEventId := "1"
	ctx := context.Background()

	eventRepository.EXPECT().
		RegisterEvent(ctx, event).
		Return(wantEventId, nil)

	gotEventId, err := usecase.RegisterEvent(ctx, event)
	if err != nil {
		t.Fatalf("unexpected error while registering event: %v", err)
	}
	if wantEventId != gotEventId {
		t.Fatalf("unexpected result: expected id %s, but got %s", wantEventId, gotEventId)
	}
}
