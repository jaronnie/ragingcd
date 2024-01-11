package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/middlewares"
)

func init() {
	repoApi := RouterGroup.Group("/repo")
	BuildRepoRouter(repoApi)
}

func BuildRepoRouter(rg *gin.RouterGroup) {
	rg.POST("/create", middlewares.Auth())
	rg.GET("/list", middlewares.Auth())
}
