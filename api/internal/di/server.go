package di

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paralleltree/go-leaderboard/internal/config"
	"github.com/paralleltree/go-leaderboard/internal/driver"
	"github.com/paralleltree/go-leaderboard/internal/handler"
	"github.com/paralleltree/go-leaderboard/internal/repository"
	"github.com/paralleltree/go-leaderboard/internal/usecase"
)

const (
	timeBitWidth = 24
)

func InflateHandlers(env config.Env, router *gin.Engine) {
	redisHashDriver := driver.NewRedisHashDriver(env.RedisEndpoint)
	redisSortedSetDriver := driver.NewRedisSortedSetDriver(env.RedisEndpoint)
	scoringStrategy := driver.NewScoringWithRemainingTimeStrategy(timeBitWidth)

	eventRepository := repository.NewEventRepository(redisHashDriver)
	scoreRepository := repository.NewScoreRepository(scoringStrategy, redisSortedSetDriver)

	registerEventUsecase := usecase.NewRegisterEventUsecase(eventRepository)
	setScoreUsecase := usecase.NewSetScoreUsecase(eventRepository, scoreRepository, buildCurrentTimeProvider())

	// inflate handlers
	router.POST("/events", handler.BuildRegisterEventHandler(registerEventUsecase))
	router.PUT("/events/:id/scores", handler.BuildSetScoreHandler(setScoreUsecase))
}

func buildCurrentTimeProvider() usecase.TimeProvider {
	return func() time.Time {
		return time.Now().UTC()
	}
}
