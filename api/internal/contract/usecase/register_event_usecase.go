package usecase

import (
	"context"

	"github.com/paralleltree/go-leaderboard/internal/model"
)

type RegisterEventUsecase interface {
	// Registers given event then returns event id.
	RegisterEvent(ctx context.Context, event model.Event) (string, error)
}
