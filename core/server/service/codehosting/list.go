package codehosting

import (
	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/domain/po"
	"github.com/jaronnie/ragingcd/core/server/domain/query"
	"github.com/jaronnie/ragingcd/core/server/domain/vo"
	"github.com/jaronnie/ragingcd/core/server/engine/db"
	"github.com/jaronnie/ragingcd/core/server/engine/helper"
	"github.com/jaronnie/ragingcd/core/server/engine/response"
)

func List(ctx *gin.Context) {
	var q query.CodeHostingListQuery
	err := ctx.BindQuery(&q)
	if err != nil {
		response.Fail(ctx, err, 500)
		return
	}

	session := db.Engine.Limit(q.PageSize, q.PageSize*(q.PageNum-1))
	session.OrderBy("updated_at DESC")

	authHelper := helper.AuthHelper{Context: ctx}
	session.Where("user_id = ?", authHelper.UserID())

	var codehostingPos []*po.CodeHosting
	count, err := session.FindAndCount(&codehostingPos)
	if err != nil {
		response.Fail(ctx, err, 500)
	}

	var codehostingVos []*vo.CodeHostingVo

	for _, v := range codehostingPos {
		codehostingVos = append(codehostingVos, &vo.CodeHostingVo{
			ID:       v.ID,
			UUID:     v.UUID,
			Name:     v.Name,
			Type:     v.Type,
			Url:      v.Url,
			Username: v.Username,
		})
	}

	table := vo.CodeHostingTableData{
		Total: int(count),
		Rows:  codehostingVos,
	}

	response.OK(ctx, table)
}
