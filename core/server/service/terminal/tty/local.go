package tty

import (
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/message"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/handler"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"

	"github.com/creack/pty"
)

func init() {
	Register("local", NewLocal)
}

type Local struct {
	target     *target.Target
	ptyHandler handler.Handler
	wsHandler  handler.WsHandler
	errChan    chan error
	doneChan   chan struct{}
	shell      *os.File
	cmd        *exec.Cmd
}

func NewLocal(target *target.Target, wsHandler handler.WsHandler, ptyHandler handler.Handler) (TTY, error) {
	return &Local{
		target:     target,
		ptyHandler: ptyHandler,
		wsHandler:  wsHandler,
		errChan:    make(chan error),
		doneChan:   make(chan struct{}),
	}, nil
}

func (s *Local) Connect(ws *handler.WsHandler) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.cmd = exec.CommandContext(ctx, "bash")
	var err error
	s.shell, err = s.create(s.cmd)
	if err != nil {
		return err
	}

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		defer func() {
			if err := recover(); err != nil {
			}
		}()

		if _, err := io.Copy(ws.Stdout(), s.shell); err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
			}
		}()

		if _, err := io.Copy(s.shell, ws.Stdin()); err != nil {
			errChan <- err
		}
	}()
	return <-errChan
}

func (s *Local) create(cmd *exec.Cmd) (*os.File, error) {
	result, err := pty.StartWithSize(cmd, &pty.Winsize{Rows: s.target.Rows, Cols: s.target.Cols})
	if err != nil {
		return nil, err
	}
	s.shell = result
	return result, nil
}

func (s *Local) Close() error {
	s.cmd.Wait()
	s.shell.Write([]byte(handler.EndOfTransmission))
	s.wsHandler.Write(message.Close("").Bytes())
	if err := s.shell.Close(); err != nil {
		return err
	}
	return nil
}

func (s *Local) Write(p []byte) (int, error) {
	return s.shell.Write(p)
}

func (s *Local) Read(p []byte) (int, error) {
	return s.shell.Read(p)
}

func (s *Local) Resize(rows, cols uint) error {
	return pty.Setsize(s.shell, &pty.Winsize{Rows: uint16(rows), Cols: uint16(cols)})
}
