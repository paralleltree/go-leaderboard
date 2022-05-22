package driver_test

import (
	"fmt"
	"testing"

	"github.com/paralleltree/go-leaderboard/internal/driver"
)

func TestScoringWithRemainingTimeStrategy_ComposeScore(t *testing.T) {
	timeBitWidth := int64(24)

	cases := []struct {
		name          string
		score         int64
		remainingTime int64
		wantResult    int64
		expectError   bool
	}{
		{
			name:          "score overflows",
			score:         int64(2 << (53 - timeBitWidth)),
			remainingTime: 0,
			expectError:   true,
		},
		{
			name:          "negative score",
			score:         -10,
			remainingTime: 0,
			expectError:   true,
		},
		{
			name:          "remaining time overflows",
			score:         0,
			remainingTime: int64(2 << timeBitWidth),
			expectError:   true,
		},
		{
			name:          "negative time",
			score:         0,
			remainingTime: -10,
			expectError:   true,
		},
		{
			name:          "valid maximum time and score",
			score:         int64((1<<(53-timeBitWidth) - 1)),
			remainingTime: int64((1 << timeBitWidth) - 1),
			wantResult:    int64((1<<(53-timeBitWidth)-1)<<timeBitWidth) | int64((1<<timeBitWidth)-1),
			expectError:   false,
		},
	}

	strategy := driver.NewScoringWithRemainingTimeStrategy(timeBitWidth)

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := strategy.ComposeScore(tt.remainingTime, tt.score)
			if tt.expectError != (err != nil) {
				t.Fatalf("unexpected result: expected hasError: %v, but got %v", tt.expectError, err)
			}
			if !tt.expectError && tt.wantResult != gotResult {
				t.Fatalf("unexpected result: expected %v, but got %v", tt.wantResult, gotResult)
			}
		})
	}
}

func TestScoringWithRemainingTimeStrategy_ExtractScore(t *testing.T) {
	timeBitWidth := int64(24)

	cases := []struct {
		rawScore  int64
		wantScore int64
	}{
		{
			rawScore:  int64(10 << timeBitWidth),
			wantScore: 10,
		},
	}

	strategy := driver.NewScoringWithRemainingTimeStrategy(timeBitWidth)

	for _, tt := range cases {
		t.Run(fmt.Sprintf("raw score %d extracted to %d", tt.rawScore, tt.wantScore), func(t *testing.T) {
			gotScore := strategy.ExtractScore(tt.rawScore)

			if tt.wantScore != gotScore {
				t.Fatalf("unexpected score: expected %v, but got %v", tt.wantScore, gotScore)
			}
		})
	}
}
