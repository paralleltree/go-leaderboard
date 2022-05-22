package repository_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/paralleltree/go-leaderboard/internal/contract/driver"
	"github.com/paralleltree/go-leaderboard/internal/model"
	"github.com/paralleltree/go-leaderboard/internal/repository"
	mock_driver "github.com/paralleltree/go-leaderboard/mock/driver"
)

const (
	eventListKey = "events_ordered_by_end_at"
)

func TestEventRepository_RegisterEvent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	idGenerator := mock_driver.NewMockUniqueIdGenerator(mockCtrl)
	hashDriver := mock_driver.NewMockHashDriver(mockCtrl)
	sortedSetDriver := mock_driver.NewMockSortedSetDriver(mockCtrl)

	eventRepository := repository.NewEventRepository(idGenerator, hashDriver, sortedSetDriver)

	keyPrefix := "events:"
	event := model.Event{
		Name:    "test event",
		StartAt: time.Date(2022, 5, 1, 12, 0, 0, 0, time.UTC),
		EndAt:   time.Date(2022, 5, 1, 13, 0, 0, 0, time.UTC),
	}
	issuedId := "test_event"
	wantFields := map[string]string{
		"name":     event.Name,
		"start_at": strconv.FormatInt(event.StartAt.Unix(), 10),
		"end_at":   strconv.FormatInt(event.EndAt.Unix(), 10),
	}
	ctx := context.Background()

	idGenerator.EXPECT().
		GenerateNewId().
		Return(issuedId, nil)
	hashDriver.EXPECT().
		Set(ctx, fmt.Sprintf("%s%s", keyPrefix, issuedId), wantFields).
		Return(nil)
	sortedSetDriver.EXPECT().
		SetScore(ctx, eventListKey, issuedId, float64(event.EndAt.Unix())).
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
	idGenerator := mock_driver.NewMockUniqueIdGenerator(mockCtrl)
	hashDriver := mock_driver.NewMockHashDriver(mockCtrl)
	sortedSetDriver := mock_driver.NewMockSortedSetDriver(mockCtrl)

	eventRepository := repository.NewEventRepository(idGenerator, hashDriver, sortedSetDriver)

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
	idGenerator := mock_driver.NewMockUniqueIdGenerator(mockCtrl)
	hashDriver := mock_driver.NewMockHashDriver(mockCtrl)
	sortedSetDriver := mock_driver.NewMockSortedSetDriver(mockCtrl)

	eventRepository := repository.NewEventRepository(idGenerator, hashDriver, sortedSetDriver)

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

func TestEventRepository_GetEvents(t *testing.T) {
	cases := []struct {
		Name          string
		Page          int64
		Count         int64
		WantStartRank int64
		WantEndRank   int64
		HasError      bool
	}{
		{
			Name:     "page is 0",
			Page:     0,
			Count:    10,
			HasError: true,
		},
		{
			Name:     "count is 0",
			Page:     10,
			Count:    0,
			HasError: true,
		},
		{
			Name:          "valid range",
			Page:          2,
			Count:         50,
			WantStartRank: 50,
			WantEndRank:   99,
			HasError:      false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			idGenerator := mock_driver.NewMockUniqueIdGenerator(mockCtrl)
			hashDriver := mock_driver.NewMockHashDriver(mockCtrl)
			sortedSetDriver := mock_driver.NewMockSortedSetDriver(mockCtrl)

			eventRepository := repository.NewEventRepository(idGenerator, hashDriver, sortedSetDriver)

			ctx := context.Background()
			wantEvent := model.Record[model.Event]{
				Id: "1",
				Item: model.Event{
					Name:    "test event",
					StartAt: time.Date(2022, 5, 1, 12, 0, 0, 0, time.UTC),
					EndAt:   time.Date(2022, 5, 1, 15, 0, 0, 0, time.UTC),
				},
			}
			wantEvents := []model.Record[model.Event]{wantEvent}
			registeredEvents := map[string]map[string]string{
				"events:1": {
					"name":     "test event",
					"start_at": strconv.FormatInt(wantEvent.Item.StartAt.Unix(), 10),
					"end_at":   strconv.FormatInt(wantEvent.Item.EndAt.Unix(), 10),
				},
			}

			if !tt.HasError {
				sortedSetDriver.EXPECT().
					GetRankedList(ctx, eventListKey, tt.WantStartRank, tt.WantEndRank).
					Return([]driver.SortedSetItem{{Member: "1", Score: float64(wantEvent.Item.EndAt.Unix())}}, true, nil)
				hashDriver.EXPECT().
					Get(ctx, gomock.Any()).
					DoAndReturn(func(ctx context.Context, key string) (map[string]string, bool, error) {
						h, ok := registeredEvents[key]
						return h, ok, nil
					})
			}

			gotEvents, err := eventRepository.GetEvents(ctx, tt.Page, tt.Count)
			if (err != nil) != tt.HasError {
				t.Fatalf("unexpected result(error): expected haserror: %v, but got %v", tt.HasError, err)
			}
			if tt.HasError {
				return
			}
			if len(wantEvents) != len(gotEvents) {
				t.Fatalf("unexpected result count: expected %v, but got %v", len(wantEvents), len(gotEvents))
			}
			for i, wantEvent := range wantEvents {
				gotEvent := gotEvents[i]
				if wantEvent != gotEvent {
					t.Fatalf("unexpected event: expected %v, but got %v", wantEvent, gotEvent)
				}
			}
		})
	}
}
