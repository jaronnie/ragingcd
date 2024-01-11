package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/middlewares"
)

func BuildSSHRouter(rg *gin.RouterGroup) {
	rg.POST("/create", middlewares.Auth())
	rg.GET("/list", middlewares.Auth())
	rg.GET("/terminal", middlewares.Auth())
}
