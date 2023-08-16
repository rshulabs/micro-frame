package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rshulabs/micro-frame/internal/pkg/response"
)

func Ping(c *gin.Context) {
	response.WriteResponse(c, nil, "pong")
}
