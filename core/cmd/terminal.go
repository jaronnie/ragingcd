/*
Copyright © 2024 jaronnie <jaron@jaronnie.com>

*/

package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/jaronnie/ragingcd/core/pkg/restc"

	"github.com/jaronnie/ragingcd/core/pkg/server"
	"github.com/pkg/errors"
	"golang.org/x/term"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var (
	TerminalWSClient *websocket.Conn
	TerminalWsServer = "terminal-ws"
)

type (
	Op string

	MessageData struct {
		MessageType int `json:"-"`
		Op          Op
		Data        string
		Rows        int
		Cols        int
	}
)

const (
	Ping      = "ping\r"
	Stdin  Op = "stdin"
	Stdout Op = "stdout"
	Close  Op = "close"
)

// terminalCmd represents the terminal command
var terminalCmd = &cobra.Command{
	Use:   "terminal",
	Short: "core terminal",
	Long:  `core terminal.`,
	RunE:  execOperations,
}

func execOperations(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.Errorf("invlid exec args, please check")
	}

	cmd.SilenceUsage = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	resourceType := args[0]
	resourceID := args[1]

	// 退出信号
	exit := make(chan os.Signal, 1)
	var err error

	var client restc.Interface
	headers := make(http.Header, 0)
	headers.Set("Authorization", "Bearer kZBdVc4ZLliTeVOggkbE2Mkvo0TjbAR3AjEg3PnNqj3g6M53udbwfqtDjtq5IAakvsT05MjawCKXt9dQbGOSZ7un5acUERQFszzoPrCgRlj1Qkx4lkDAbiu2fnGJj4eP")
	client, err = restc.New(restc.WithUrl("ws://localhost:8081"), restc.WithHeaders(headers))
	TerminalWSClient, _, err = client.Get().SubPath("/api/v1/terminal/exec").Params(restc.QueryParam{
		Name:  "resource_type",
		Value: resourceType,
	}, restc.QueryParam{
		Name:  "resource_id",
		Value: resourceID,
	}).WsConn(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if TerminalWSClient != nil {
			defer TerminalWSClient.Close()
		}
		server.StopServer(TerminalWsServer)
	}()

	// keep alive websocket conn and write message to websocket
	err = terminalWs()
	if err != nil {
		exit <- syscall.SIGQUIT
	}

	go func() {
		if err = pushOsStdinToServer(); err != nil {
			exit <- syscall.SIGQUIT
		}
	}()

	go func() {
		var et bool
		if et, err = pullTerminalMsgToConsole(); err != nil {
			exit <- syscall.SIGQUIT
		} else if et {
			exit <- syscall.SIGQUIT
		}
	}()

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	<-exit
	return err
}

func pushOsStdinToServer() error {
	for {
		// 读取 stdin 输入
		buf := make([]byte, 1024)
		reader := bufio.NewReader(os.Stdin)

		n, _ := reader.Read(buf)

		if n <= 0 {
			continue
		}

		input := buf[0:n]
		if _, err := server.PushMsgToServer(TerminalWsServer, &MessageData{
			MessageType: websocket.TextMessage,
			Op:          Stdin,
			Data:        string(input),
		}); err != nil {
			return err
		}
	}
}

func pullTerminalMsgToConsole() (close bool, err error) {
	for {
		_, message, err := TerminalWSClient.ReadMessage()
		if err != nil {
			return false, err
		}

		var resp MessageData
		_ = json.Unmarshal(message, &resp)

		switch resp.Op {
		case Stdout:
			fmt.Print(resp.Data)
		case Close:
			// TODO: server 端修复 docker 关闭的问题
			return true, nil
		}
	}
}

func terminalWs() (err error) {
	wt, _ := server.NewSvr(TerminalWsServer,
		func(msg interface{}, num int) (resp interface{}, err error) {
			// 给 websocket 发送消息
			sendMsg, ok := msg.(*MessageData)
			if !ok {
				return nil, errors.Errorf("assert error")
			}
			marshal, _ := json.Marshal(sendMsg)
			if err = TerminalWSClient.WriteMessage(sendMsg.MessageType, marshal); err != nil {
				// write message meet error, exit terminal
				return nil, err
			}
			return
		},
		[]server.TimedTask{
			{
				Task: func(num int) error {
					// ping to keepalive
					_, err = server.PushMsgToServer(TerminalWsServer, &MessageData{
						MessageType: websocket.PingMessage,
						Op:          Ping,
					})
					if err != nil {
						return err
					}
					return nil
				},
				Time: 1 * time.Second,
			}}, &server.Option{ActionGoroutineNum: 1}) // gorilla websocket 不允许多 goroutine 并发写

	wt.Go()
	return
}

func init() {
	rootCmd.AddCommand(terminalCmd)
}
