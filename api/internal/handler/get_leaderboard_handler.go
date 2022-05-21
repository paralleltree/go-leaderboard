package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paralleltree/go-leaderboard/internal/contract/usecase"
	"github.com/paralleltree/go-leaderboard/internal/model"
)

type getLeaderboardParams struct {
	StartRank int64 `form:"start,default=1"`
	EndRank   int64 `form:"end,default=100"`
}

func BuildGetLeaderboardHandler(usecase usecase.GetLeaderboardUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId := c.Param("id")
		params := getLeaderboardParams{}
		if err := c.ShouldBind(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ranks, ok, err := usecase.GetLeaderboard(c.Request.Context(), eventId, params.StartRank, params.EndRank)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !ok {
			// TOOD: Replace own error interface
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
			return
		}
		c.JSON(http.StatusOK, convertToJsonUserRank(ranks))
	}
}

type userRank struct {
	Rank   int64  `json:"rank"`
	UserId string `json:"user_id"`
	Score  int64  `json:"score"`
}

func convertToJsonUserRank(src []model.UserRank) []userRank {
	res := make([]userRank, 0, len(src))
	for _, v := range src {
		res = append(res, userRank(v))
	}
	return res
}
