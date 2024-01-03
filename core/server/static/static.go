package static

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/public"
)

func Static(r *gin.RouterGroup, f fs.FS) {
	// 将 static 目录下的文件作为静态文件提供服务
	staticHandler, err := fs.Sub(public.Public, "dist")
	if err != nil {
		log.Fatal("Unable to load static files: ", err)
	}
	r.StaticFS("/", http.FS(staticHandler))
}
