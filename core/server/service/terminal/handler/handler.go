package handler

import (
	"fmt"
	"io"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/conn"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/message"

	"github.com/pkg/errors"
	cmd "k8s.io/client-go/tools/remotecommand"
)

const (
	// EndOfTransmission 终端结束符
	EndOfTransmission = "\u0004"
	// NullCharacter 空操作
	NullCharacter = "\u0000"
)

type Handler interface {
	cmd.TerminalSizeQueue
	Done()
	Tty() bool
	Stdin() io.Reader
	Stdout() io.Writer
	Stderr() io.Writer
	Error() error
}

type WsHandler struct {
	conn     conn.Connection
	target   *target.Target
	sizeChan chan cmd.TerminalSize
	doneChan chan struct{}
	tty      bool
	err      error
}

func New(conn conn.Connection, target *target.Target) *WsHandler {
	session := &WsHandler{
		conn:     conn,
		tty:      true,
		sizeChan: make(chan cmd.TerminalSize),
		doneChan: make(chan struct{}),
		target:   target,
	}
	return session
}

func (t *WsHandler) Done() {
	defer close(t.doneChan)
	t.doneChan <- struct{}{}
}

func (t *WsHandler) Next() *cmd.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}

func (t *WsHandler) Stdin() io.Reader {
	return t
}

func (t *WsHandler) Stdout() io.Writer {
	return t
}

func (t *WsHandler) Stderr() io.Writer {
	return t
}

func (t *WsHandler) Read(p []byte) (int, error) {
	msg, err := t.conn.ReadMsg()
	if err != nil {
		t.err = err
		return copy(p, EndOfTransmission), err
	}
	if msg == nil {
		return 0, nil
	}
	count, err := t.handleMsg(p, msg)
	if err != nil {
		t.err = err
		return count, errors.Wrap(err, "Fail to handle msg")
	}
	return count, nil
}

func (t *WsHandler) handleMsg(terminalStdin []byte, msg *message.TerminalMessage) (int, error) {
	return msgAction(msg.Operation)(terminalStdin, msg, t.sizeChan, t.target)
}

func (t *WsHandler) Write(p []byte) (int, error) {
	fmt.Printf("websocket write. [%s]", p)
	if err := t.conn.WriteMsg(message.Stdout(string(p))); err != nil {
		t.err = err
		return 0, err
	}
	return len(p), nil
}

func (t *WsHandler) Tty() bool {
	return t.tty
}

func (t *WsHandler) Error() error {
	return t.err
}
