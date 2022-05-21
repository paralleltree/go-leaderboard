//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"context"

	"github.com/paralleltree/go-leaderboard/internal/model"
)

type GetLeaderboardUsecase interface {
	GetLeaderboard(ctx context.Context, eventId string, startRank, endRank int64) ([]model.UserRank, bool, error)
}
