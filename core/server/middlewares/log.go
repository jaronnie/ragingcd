package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/pkg/logx"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		logx.Infof("receive request from ip [%v]", c.ClientIP())
	}
}
