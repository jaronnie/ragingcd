package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/middlewares"
	"github.com/jaronnie/ragingcd/core/server/service/ssh"
)

func init() {
	sshApi := RouterGroup.Group("/ssh")
	BuildSshRouter(sshApi)
}

func BuildSshRouter(rg *gin.RouterGroup) {
	rg.POST("/create", middlewares.Auth(), ssh.Create)
	rg.GET("/list", middlewares.Auth(), ssh.List)
}
