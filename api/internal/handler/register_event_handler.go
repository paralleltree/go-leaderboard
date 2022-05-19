package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
	"github.com/paralleltree/go-leaderboard/internal/lib"
	"github.com/paralleltree/go-leaderboard/internal/model"
)

type eventParams struct {
	Name    string `json:"name" binding:"required"`
	StartAt string `json:"start_at" binding:"required"`
	EndAt   string `json:"end_at" binding:"required"`
}

func BuildRegisterEventHandler(registerEventUsecase usecase.RegisterEventUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := eventParams{}
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		event, err := buildEvent(params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := registerEventUsecase.RegisterEvent(c.Request.Context(), event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func buildEvent(params eventParams) (model.Event, error) {
	startAt, err := lib.ParseDateTime(params.StartAt)
	if err != nil {
		return model.Event{}, fmt.Errorf("parse start_at: %w", err)
	}
	endAt, err := lib.ParseDateTime(params.EndAt)
	if err != nil {
		return model.Event{}, fmt.Errorf("parse end_at: %w", err)
	}
	return model.Event{
		Name:    params.Name,
		StartAt: startAt,
		EndAt:   endAt,
	}, nil
}
