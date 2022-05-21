package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
)

type setScoreParams struct {
	UserId string `json:"user_id" binding:"required"`
	Score  int64  `json:"score" binding:"required"`
}

func BuildSetScoreHandler(usecase usecase.SetScoreUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := setScoreParams{}
		eventId := c.Param("id")
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := usecase.SetScore(c.Request.Context(), eventId, params.UserId, params.Score); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	}
}
