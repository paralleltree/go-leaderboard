package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
	"github.com/paralleltree/go-leaderboard/internal/lib"
	"github.com/paralleltree/go-leaderboard/internal/model"
)

type getEventsParams struct {
	Page  int64 `form:"page,default=1"`
	Count int64 `form:"count,default=20"`
}

func BuildGetEventsHandler(usecase usecase.GetEventsUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := getEventsParams{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		events, err := usecase.GetEvents(c.Request.Context(), params.Page, params.Count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, lib.Map(events, eventRecordToEventResponse))
	}
}

type event struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}

func eventRecordToEventResponse(r model.Record[model.Event]) event {
	return event{
		Id:      r.Id,
		Name:    r.Item.Name,
		StartAt: lib.FormatDateTime(r.Item.StartAt),
		EndAt:   lib.FormatDateTime(r.Item.EndAt),
	}
}
