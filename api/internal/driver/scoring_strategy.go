package driver

import (
	"fmt"
	"math"

	"github.com/paralleltree/go-leaderboard/internal/contract/driver"
)

const (
	significandBitWidth = 53
)

type scoringWithRemainingTimeStrategy struct {
	timeBitWidth int64
}

func NewScoringWithRemainingTimeStrategy(timeBitWidth int64) driver.ScoringStrategy {
	return &scoringWithRemainingTimeStrategy{
		timeBitWidth: timeBitWidth,
	}
}

func (s scoringWithRemainingTimeStrategy) ComposeScore(time, score int64) (int64, error) {
	if !isValidBitWidth(time, s.timeBitWidth) || time < 0 {
		return 0, fmt.Errorf("time out of range: %d", time)
	}
	if !isValidBitWidth(score, significandBitWidth-s.timeBitWidth) || score < 0 {
		return 0, fmt.Errorf("score out of range: %d", score)
	}
	return (score << s.timeBitWidth) | time, nil
}

func (s scoringWithRemainingTimeStrategy) ExtractScore(rawScore int64) int64 {
	return rawScore >> int64(s.timeBitWidth)
}

func isValidBitWidth(value int64, bitWidth int64) bool {
	max := int64(math.Pow(2, float64(bitWidth))) - 1
	return value <= max
}
