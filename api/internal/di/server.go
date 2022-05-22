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
	idGenerator := driver.NewUuidGenerator()

	eventRepository := repository.NewEventRepository(idGenerator, redisHashDriver, redisSortedSetDriver)
	scoreRepository := repository.NewScoreRepository(scoringStrategy, redisSortedSetDriver)

	getEventsUsecase := usecase.NewGetEventsUsecase(eventRepository)
	registerEventUsecase := usecase.NewRegisterEventUsecase(eventRepository)
	getLeaderboardUsecase := usecase.NewGetLeaderboardUsecase(eventRepository, scoreRepository)
	setScoreUsecase := usecase.NewSetScoreUsecase(eventRepository, scoreRepository, buildCurrentTimeProvider())

	// inflate handlers
	router.GET("/events", handler.BuildGetEventsHandler(getEventsUsecase))
	router.POST("/events", handler.BuildRegisterEventHandler(registerEventUsecase))
	router.GET("/events/:id/leaderboard", handler.BuildGetLeaderboardHandler(getLeaderboardUsecase))
	router.PUT("/events/:id/scores", handler.BuildSetScoreHandler(setScoreUsecase))
}

func buildCurrentTimeProvider() usecase.TimeProvider {
	return func() time.Time {
		return time.Now().UTC()
	}
}
