package ssh

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
	var q query.SshListQuery
	err := ctx.BindQuery(&q)
	if err != nil {
		response.Fail(ctx, err, 500)
		return
	}

	session := db.Engine.Limit(q.PageSize, q.PageSize*(q.PageNum-1))
	session.OrderBy("updated_at DESC")

	authHelper := helper.AuthHelper{Context: ctx}
	session.Where("user_id = ?", authHelper.UserID())

	var sshPos []*po.Ssh
	count, err := session.FindAndCount(&sshPos)
	if err != nil {
		response.Fail(ctx, err, 500)
	}

	var sshVos []*vo.SshVo

	for _, v := range sshPos {
		sshVos = append(sshVos, &vo.SshVo{
			ID:       v.ID,
			UUID:     v.UUID,
			Name:     v.Name,
			Type:     v.Type,
			IP:       v.IP,
			Port:     v.Port,
			Username: v.Username,
		})
	}

	table := vo.PageData{
		Total: int(count),
		Rows:  sshVos,
	}

	response.OK(ctx, table)
}
