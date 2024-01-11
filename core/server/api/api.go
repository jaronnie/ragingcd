package api

import (
	"github.com/gin-gonic/gin"

	_ "github.com/jaronnie/ragingcd/core/server/api/v1"
	"github.com/jaronnie/ragingcd/core/server/engine"
)

func Router() {
	engine.ServerEngine.GET("/api/health", func(ctx *gin.Context) {
		ctx.JSON(200, struct {
			Message string `json:"message"`
		}{Message: "ok"})
	})
}
