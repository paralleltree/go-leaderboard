package model_test

import (
	"testing"
	"time"

	"github.com/paralleltree/go-leaderboard/internal/model"
)

func TestEvent_IsInSession(t *testing.T) {
	tz := time.UTC
	startAt := time.Date(2022, 5, 1, 12, 0, 0, 0, tz)
	endAt := time.Date(2022, 5, 1, 12, 30, 0, 0, tz)
	event := model.Event{
		StartAt: startAt,
		EndAt:   endAt,
	}
	cases := []struct {
		name       string
		now        time.Time
		wantResult bool
	}{
		{
			name:       "before startAt",
			now:        time.Date(2022, 5, 1, 11, 59, 59, 0, tz),
			wantResult: false,
		},
		{
			name:       "after endAt",
			now:        time.Date(2022, 5, 1, 12, 30, 1, 0, tz),
			wantResult: false,
		},
		{
			name:       "at startAt",
			now:        startAt,
			wantResult: true,
		},
		{
			name:       "at endAt",
			now:        endAt,
			wantResult: false,
		},
		{
			name:       "in period",
			now:        time.Date(2022, 5, 1, 12, 15, 0, 0, tz),
			wantResult: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := event.IsInSession(tt.now)
			if tt.wantResult != gotResult {
				t.Fatalf("invalid result with value %v: expected %v, but got %v", tt.now, tt.wantResult, gotResult)
			}
		})
	}
}
