package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/fronted-backend-admin/backend-go/public"
	apiv1 "github.com/jaronnie/fronted-backend-admin/backend-go/server/api/v1"
	"github.com/jaronnie/fronted-backend-admin/backend-go/server/static"
)

func Router(e *gin.Engine) {
	// redirect 到 /ui
	e.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/ui")
	})

	ui := e.Group("/ui")
	static.Static(ui, public.Public)

	apiV1 := e.Group("/api/v1")
	apiv1.ApiRouter(apiV1)
}

func Cors(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	r.Use(cors.New(config))
}
