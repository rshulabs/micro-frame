package controller

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/rshulabs/micro-frame/internal/pkg/middlewares"
)

func installHttpRouter(g *gin.Engine) *gin.Engine {
	api := g.Group("/api").Use(middleware.Cors())
	{
		api.GET("/ping", Ping)
	}
	return g
}
