package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/middlewares"
	"github.com/jaronnie/ragingcd/core/server/service/codehosting"
)

func BuildCodeHostingRouter(rg *gin.RouterGroup) {
	rg.POST("/create", middlewares.Auth(), codehosting.Create)
	rg.GET("/list", middlewares.Auth(), codehosting.List)
}
