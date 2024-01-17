package ssh

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/domain/po"
	"github.com/jaronnie/ragingcd/core/server/engine/db"
	"github.com/jaronnie/ragingcd/core/server/engine/response"
	"github.com/spf13/cast"
)

func Delete(ctx *gin.Context) {
	id := cast.ToInt(ctx.Param("id"))

	if _, err := db.Engine.Delete(&po.Ssh{
		Base: po.Base{ID: id},
	}); err != nil {
		response.Fail(ctx, err, 500)
		return
	}
	response.OK(ctx, true)
}
