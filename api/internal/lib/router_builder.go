package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/paralleltree/go-leaderboard/internal/config"
)

func InflateRouterMiddleware(router *gin.Engine, env config.Env) {
	if env.AllowOrigin != "" {
		router.Use(BuildCorsMiddleware(env.AllowOrigin))
	}
}
