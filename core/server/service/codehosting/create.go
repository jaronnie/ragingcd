package codehosting

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jaronnie/ragingcd/core/server/domain/bo"
	"github.com/jaronnie/ragingcd/core/server/domain/po"
	"github.com/jaronnie/ragingcd/core/server/domain/vo"
	"github.com/jaronnie/ragingcd/core/server/engine/db"
	"github.com/jaronnie/ragingcd/core/server/engine/helper"
	"github.com/jaronnie/ragingcd/core/server/engine/response"
)

func Create(ctx *gin.Context) {
	var hosting bo.CreateCodeHostingBo
	err := ctx.ShouldBindJSON(&hosting)
	if err != nil {
		response.Fail(ctx, err, 500)
		return
	}

	authHelper := helper.AuthHelper{Context: ctx}

	insert := &po.CodeHosting{
		UUID:     "CH-" + uuid.New().String()[:6],
		Name:     hosting.Name,
		Type:     hosting.Type,
		Url:      hosting.Url,
		Username: hosting.Username,
		Token:    hosting.Token,
		UserID:   authHelper.UserID(),
	}
	_, err = db.Engine.Insert(insert)

	if err != nil {
		response.Fail(ctx, err, 500)
		return
	}
	response.OK(ctx, vo.CodeHostingVo{
		ID:       insert.ID,
		UUID:     insert.UUID,
		Name:     insert.Name,
		Type:     insert.Type,
		Url:      insert.Url,
		Username: insert.Username,
	})
}
