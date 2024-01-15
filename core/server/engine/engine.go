package engine

import (
	"os"

	"github.com/jaronnie/ragingcd/core/pkg/logx"

	"github.com/gin-gonic/gin"
)

var ServerEngine *gin.Engine

func init() {
	if os.Args[1] != "server" {
		gin.SetMode(gin.ReleaseMode)
	}
	ServerEngine = gin.New()
	logx.Logger()
}
