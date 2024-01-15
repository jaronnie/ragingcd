package middlewares

import "github.com/gin-gonic/gin"

func Use(engine *gin.Engine) {
	engine.Use(Cors(), Log())
}
