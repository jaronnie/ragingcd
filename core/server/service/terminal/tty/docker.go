package tty

import (
	"context"
	"io"

	"github.com/jaronnie/ragingcd/core/pkg/logx"

	"github.com/jaronnie/ragingcd/core/server/domain/po"
	"github.com/jaronnie/ragingcd/core/server/engine/db"
	"github.com/spf13/cast"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/handler"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func init() {
	Register("docker", NewDocker)
}

type Docker struct {
	target      *target.Target
	ContainerID string
	ptyHandler  handler.Handler
	wsHandler   handler.WsHandler
	client      *client.Client
	conn        *types.HijackedResponse
}

func NewDocker(target *target.Target, wsHandler handler.WsHandler, ptyHandler handler.Handler) (TTY, error) {
	return &Docker{
		target:     target,
		ptyHandler: ptyHandler,
		wsHandler:  wsHandler,
	}, nil
}

func (s *Docker) Connect(ws *handler.WsHandler) error {
	shell, err := s.create()
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
		if _, err := io.Copy(ws.Stdout(), shell.Conn); err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
			}
		}()
		if _, err := io.Copy(ws.Stderr(), shell.Conn); err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
			}
		}()

		if _, err := io.Copy(shell.Conn, ws.Stdin()); err != nil {
			errChan <- err
		}
	}()
	s.conn = shell
	return <-errChan
}

func (s *Docker) create() (*types.HijackedResponse, error) {
	var err error
	dockerPo := po.Docker{
		Base: po.Base{ID: cast.ToInt(s.target.ResourceID)},
	}
	b, err := db.Engine.Get(&dockerPo)
	if !b {
		return nil, ErrTargetNotFound
	}
	if err != nil {
		return nil, err
	}

	s.client, err = client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.40"))
	if err != nil {
		return nil, err
	}

	s.ContainerID = dockerPo.ContainerID
	if s.target.WorkingDir == "" {
		c, err := s.client.ContainerInspect(context.Background(), s.ContainerID)
		if err != nil {
			return nil, err
		}
		s.target.WorkingDir = c.Config.WorkingDir
	}

	ctx := context.Background()
	exec, err := s.client.ContainerExecCreate(ctx, s.ContainerID, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
		Tty:          true,
	})
	if err != nil {
		return nil, err
	}
	result, err := s.client.ContainerExecAttach(ctx, exec.ID, types.ExecStartCheck{Detach: false, Tty: true})
	return &result, err
}

func (s *Docker) Close() error {
	logx.Debugf("close docker terminal")
	s.wsHandler.WriteClose([]byte("exit"))
	s.conn.CloseWrite()
	return s.conn.Conn.Close()
}

func (s *Docker) Resize(rows, cols uint) error {
	return s.client.ContainerExecResize(context.Background(), s.ContainerID,
		types.ResizeOptions{
			Height: rows,
			Width:  cols,
		},
	)
}
