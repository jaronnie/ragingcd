package engine

import "github.com/gin-gonic/gin"

var ServerEngine *gin.Engine

func init() {
	ServerEngine = gin.Default()
}
