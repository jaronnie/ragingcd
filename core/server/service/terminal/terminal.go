package terminal

import (
	"encoding/json"
	"strings"

	"github.com/jaronnie/ragingcd/core/pkg/websocat"
	"github.com/jaronnie/ragingcd/core/server/domain/bo"

	"github.com/gin-gonic/gin"
	"github.com/jaronnie/ragingcd/core/server/engine/response"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/handler"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/message"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/tty"
)

func Terminal(ctx *gin.Context) {
	if err := terminal(ctx); err != nil {
		response.Fail(ctx, err, 500)
	}
}

func terminal(ctx *gin.Context) error {
	var terminalBo bo.ExecTerminalBo
	err := ctx.BindQuery(&terminalBo)
	if err != nil {
		return err
	}

	// create websocket connection
	connection, err := websocat.NewConn(ctx.Writer, ctx.Request)
	if err != nil {
		return err
	}
	defer connection.Close()

	doneMsg := func() *message.TerminalMessage {
		findTarget, err := target.FindTarget(terminalBo)
		if err != nil {
			return message.Close(err.Error())
		}

		connHandler := handler.New(connection, findTarget)

		terminal, err := tty.New(findTarget, *connHandler, nil)
		if err != nil {
			return message.Close(err.Error())
		}

		go func() {
			for {
				winSize := connHandler.Next()
				if winSize == nil {
					break
				}
				err = terminal.Resize(uint(winSize.Height), uint(winSize.Width))
				if err != nil {
					break
				}
			}
		}()

		if err := terminal.Connect(connHandler); err != nil {
			return message.Close(err.Error())
		}
		defer terminal.Close()

		return convertErrToMsg(connHandler.Error(), err)
	}()

	if doneMsg != nil {
		msg, _ := json.Marshal(doneMsg)
		err := connection.WriteMsg(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func convertErrToMsg(connErr, shellErr error) *message.TerminalMessage {
	if connErr == nil && shellErr == nil {
		return nil
	}

	if connErr != nil {
		if isTimeout(connErr) {
			return message.Close("连接超时")
		}
		return message.Close(connErr.Error())
	}

	if shellErr != nil {
		return message.Close("shell error")
	}

	return nil
}

func isTimeout(err error) bool {
	return strings.Contains(err.Error(), "timeout")
}
