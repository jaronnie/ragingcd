package v1

import (
	"github.com/gin-gonic/gin"
)

func ApiRouter(rg *gin.RouterGroup) {
	rg.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, struct {
			Message string `json:"message"`
		}{Message: "ok"})
	})

	codeHostingApi := rg.Group("/codehosting")
	BuildCodeHostingRouter(codeHostingApi)

	repoApi := rg.Group("/repo")
	BuildRepoRouter(repoApi)

	projectApi := rg.Group("/project")
	BuildProjectRouter(projectApi)

	sshApi := rg.Group("/ssh")
	BuildSSHRouter(sshApi)
}
