package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/middlewares"
	"github.com/jaronnie/ragingcd/core/server/service/terminal"
)

func init() {
	terminalApi := RouterGroup.Group("/terminal")
	BuildTerminalRouter(terminalApi)
}

func BuildTerminalRouter(rg *gin.RouterGroup) {
	rg.GET("/exec", middlewares.Auth(), terminal.Terminal)
}
