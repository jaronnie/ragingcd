package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/middlewares"
)

func BuildCodeHostingRouter(rg *gin.RouterGroup) {
	rg.GET("/list", middlewares.Auth(), func(c *gin.Context) {
		c.String(200, "hello")
	})
}
