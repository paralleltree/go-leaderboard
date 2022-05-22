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

func TestGetEventsUsecase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	eventRepository := mock_repository.NewMockEventRepository(mockCtrl)

	wantPage := int64(1)
	wantCount := int64(50)
	wantEvents := []model.Record[model.Event]{
		{
			Id:   "1",
			Item: model.Event{Name: "test event", StartAt: time.Now(), EndAt: time.Now()},
		},
	}
	ctx := context.Background()

	eventRepository.EXPECT().
		GetEvents(ctx, wantPage, wantCount).
		Return(wantEvents, nil)

	usecase := usecase.NewGetEventsUsecase(eventRepository)

	gotEvents, err := usecase.GetEvents(ctx, wantPage, wantCount)
	if err != nil {
		t.Fatalf("unexpected error while getting events: %v", err)
	}
	if len(wantEvents) != len(gotEvents) {
		t.Fatalf("unexpected result count: expected length %v, but got %v", len(wantEvents), len(gotEvents))
	}
	for i, wantEvent := range wantEvents {
		gotEvent := gotEvents[i]
		if wantEvent != gotEvent {
			t.Fatalf("unexpected event: expected %v, but got %v", wantEvent, gotEvent)
		}
	}
}
