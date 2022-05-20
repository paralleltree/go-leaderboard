package di

import (
	"github.com/gin-gonic/gin"
	"github.com/paralleltree/go-leaderboard/internal/config"
	"github.com/paralleltree/go-leaderboard/internal/driver"
	"github.com/paralleltree/go-leaderboard/internal/handler"
	"github.com/paralleltree/go-leaderboard/internal/repository"
	"github.com/paralleltree/go-leaderboard/internal/usecase"
)

func InflateHandlers(env config.Env, router *gin.Engine) {
	redisHashDriver := driver.NewRedisHashDriver(env.RedisEndpoint)
	eventRepository := repository.NewEventRepository(redisHashDriver)
	registerEventUsecase := usecase.NewRegisterEventUsecase(eventRepository)

	// inflate handlers
	router.POST("/events", handler.BuildRegisterEventHandler(registerEventUsecase))
}
