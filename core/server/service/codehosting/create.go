package codehosting

import (
	"errors"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/domain/bo"
	"github.com/jaronnie/ragingcd/core/server/domain/po"
	"github.com/jaronnie/ragingcd/core/server/domain/vo"
	"github.com/jaronnie/ragingcd/core/server/engine/db"
	"github.com/jaronnie/ragingcd/core/server/engine/githosting"
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

	gh, err := githosting.New(githosting.Config{
		Type:     hosting.Type,
		Url:      hosting.Url,
		Username: hosting.Username,
		Token:    hosting.Token,
	})
	if err != nil {
		response.Fail(ctx, err, 500)
		return
	}

	// 校验 token 是否有效
	if gh.VerifyToken() != nil {
		response.Fail(ctx, errors.New("token is invalid"), 500)
		return
	}

	if hosting.Type == githosting.GITHUB {
		hosting.Url = "https://github.com"
	}
	insert := &po.CodeHosting{

		Name:     hosting.Name,
		Type:     hosting.Type,
		Url:      hosting.Url,
		Username: hosting.Username,
		Token:    hosting.Token,
		Base: po.Base{
			UserID: authHelper.UserID(),
			UUID:   "CH-" + uuid.New().String()[:6],
		},
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
