package ssh

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jaronnie/ragingcd/core/pkg/sshx"
	"github.com/jaronnie/ragingcd/core/server/domain/bo"
	"github.com/jaronnie/ragingcd/core/server/domain/po"
	"github.com/jaronnie/ragingcd/core/server/domain/vo"
	"github.com/jaronnie/ragingcd/core/server/engine/db"
	"github.com/jaronnie/ragingcd/core/server/engine/helper"
	"github.com/jaronnie/ragingcd/core/server/engine/response"
)

func Create(ctx *gin.Context) {
	var sshBo bo.CreateSshBo
	err := ctx.ShouldBindJSON(&sshBo)
	if err != nil {
		response.Fail(ctx, err, 500)
		return
	}

	authHelper := helper.AuthHelper{Context: ctx}

	// 校验是否能登录
	if _, err = sshx.NewSshClient(&sshx.Config{
		IP:         sshBo.IP,
		Port:       sshBo.Port,
		Username:   sshBo.Username,
		Type:       sshBo.Type,
		Password:   sshBo.Password,
		PrivateKey: sshBo.PrivateKey,
		Timeout:    10,
	}); err != nil {
		response.Fail(ctx, err, 500)
		return
	}

	sshPo := &po.Ssh{
		Name:       sshBo.Name,
		Type:       sshBo.Type,
		IP:         sshBo.IP,
		Port:       sshBo.Port,
		Username:   sshBo.Username,
		Password:   sshBo.Password,
		PrivateKey: sshBo.PrivateKey,
		Base: po.Base{
			UserID: authHelper.UserID(),
			UUID:   "SSH-" + uuid.New().String()[:6],
		},
	}
	_, err = db.Engine.Insert(sshPo)
	if err != nil {
		response.Fail(ctx, err, 500)
		return
	}
	response.OK(ctx, vo.SshVo{
		ID:       sshPo.ID,
		UUID:     sshPo.UUID,
		Name:     sshPo.Name,
		Type:     sshPo.Type,
		Username: sshPo.Username,
	})
}
