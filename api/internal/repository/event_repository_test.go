package repository_test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/model"
	"github.com/paralleltree/go-leaderboard/internal/repository"
	mock_driver "github.com/paralleltree/go-leaderboard/mock/driver"
)

func TestEventRepository_RegisterEvent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	hashDriver := mock_driver.NewMockHashDriver(mockCtrl)

	eventRepository := repository.NewEventRepository(hashDriver)

	keyPrefix := "events:"
	event := model.Event{
		Name:    "test event",
		StartAt: time.Date(2022, 5, 1, 12, 0, 0, 0, time.UTC),
		EndAt:   time.Date(2022, 5, 1, 13, 0, 0, 0, time.UTC),
	}
	wantFields := map[string]string{
		"name":     event.Name,
		"start_at": strconv.FormatInt(event.StartAt.Unix(), 10),
		"end_at":   strconv.FormatInt(event.EndAt.Unix(), 10),
	}
	ctx := context.Background()

	hashDriver.EXPECT().
		Set(ctx, gomock.Any(), wantFields).
		Do(func(ctx context.Context, id string, fields map[string]string) {
			if !strings.HasPrefix(id, keyPrefix) {
				t.Fatalf("event data key does not start with `%s`: %s", keyPrefix, id)
			}
		}).
		Return(nil)

	id, err := eventRepository.RegisterEvent(ctx, event)
	if err != nil {
		t.Fatalf("unexpected error while registering event: %v", err)
	}
	if id == "" {
		t.Fatalf("no id returned")
	}
}

func TestEventRepository_GetEvent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	hashDriver := mock_driver.NewMockHashDriver(mockCtrl)

	eventRepository := repository.NewEventRepository(hashDriver)

	id := "111-111"
	wantName := "test event"
	wantStartAt := time.Date(2022, 5, 1, 12, 0, 0, 0, time.UTC)
	wantEndAt := time.Date(2022, 5, 1, 13, 0, 0, 0, time.UTC)
	wantEvent := model.Event{
		Name:    wantName,
		StartAt: wantStartAt,
		EndAt:   wantEndAt,
	}
	fields := map[string]string{
		"name":     wantName,
		"start_at": strconv.FormatInt(wantStartAt.Unix(), 10),
		"end_at":   strconv.FormatInt(wantEndAt.Unix(), 10),
	}
	wantOk := true
	ctx := context.Background()

	hashDriver.EXPECT().
		Get(ctx, fmt.Sprintf("events:%s", id)).
		Return(fields, true, nil)

	gotEvent, gotOk, err := eventRepository.GetEvent(ctx, id)
	if err != nil {
		t.Fatalf("unexpected error while getting event: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
	if wantEvent != gotEvent {
		t.Fatalf("unexpected result: expected %v, but got %v", wantEvent, gotEvent)
	}
}

func TestEventRepository_GetEvent_WhenEventNotExists(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	hashDriver := mock_driver.NewMockHashDriver(mockCtrl)

	eventRepository := repository.NewEventRepository(hashDriver)

	key := "111-111"
	wantOk := false
	ctx := context.Background()

	hashDriver.EXPECT().
		Get(ctx, fmt.Sprintf("events:%s", key)).
		Return(nil, false, nil)

	_, gotOk, err := eventRepository.GetEvent(ctx, key)
	if err != nil {
		t.Fatalf("unexpected error while getting event: %v", err)
	}
	if wantOk != gotOk {
		t.Fatalf("unexpected result(ok): expected %v, but got %v", wantOk, gotOk)
	}
}
