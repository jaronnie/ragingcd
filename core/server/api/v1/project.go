package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/middlewares"
)

func init() {
	projectApi := RouterGroup.Group("/project")
	BuildProjectRouter(projectApi)
}

func BuildProjectRouter(rg *gin.RouterGroup) {
	rg.POST("/create", middlewares.Auth())
	rg.GET("/list", middlewares.Auth())
	rg.GET("/:id", middlewares.Auth())
	rg.GET("/repo/list", middlewares.Auth())
	rg.GET("/environment/list", middlewares.Auth())
	rg.GET("/cicd/list", middlewares.Auth())
}
