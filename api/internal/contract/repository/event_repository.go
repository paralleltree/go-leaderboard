//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"context"

	"github.com/paralleltree/go-leaderboard/internal/model"
)

type EventRepository interface {
	// Registers given event then returns assigned id.
	RegisterEvent(ctx context.Context, event model.Event) (string, error)
	GetEvent(ctx context.Context, id string) (model.Event, bool, error)
	// Returns Events in descending order of EndAt.
	// `page`` is 1-based.
	GetEvents(ctx context.Context, page, count int64) ([]model.Record[model.Event], error)
}
