package driver

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/paralleltree/go-leaderboard/internal/contract/driver"
)

type uuidGenerator struct{}

func NewUuidGenerator() driver.UniqueIdGenerator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) GenerateNewId() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("generate new uuid: %w", err)
	}
	return id.String(), nil
}
