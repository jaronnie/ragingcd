package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/public"
	"github.com/jaronnie/ragingcd/core/server/api"
	"github.com/jaronnie/ragingcd/core/server/static"
)

func Router(e *gin.Engine) {
	// redirect 到 /ui
	e.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/ui")
	})

	ui := e.Group("/ui")
	static.Static(ui, public.Public)

	api.Router()
}
