package repository

import (
	"context"

	"github.com/paralleltree/go-leaderboard/internal/model"
)

type EventRepository interface {
	// Registers given event then returns assigned id.
	RegisterEvent(ctx context.Context, event model.Event) (string, error)
	GetEvent(ctx context.Context, id string) (model.Event, bool, error)
}
