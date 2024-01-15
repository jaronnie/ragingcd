package tty

import (
	"context"
	"io"

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
	handler     handler.Handler
	client      *client.Client
	conn        *types.HijackedResponse
}

func NewDocker(target *target.Target, ptyHandler handler.Handler) (TTY, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}
	return NewDockerShellWithClient(target, ptyHandler, cli)
}

func NewDockerShellWithClient(target *target.Target, ptyHandler handler.Handler, client *client.Client) (TTY, error) {
	shell := &Docker{
		target:  target,
		handler: ptyHandler,
		client:  client,
	}

	return shell, nil
}

func (s *Docker) Connect(stdout io.Writer, stdin io.Reader) error {
	shell, err := s.createShell()
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
		if _, err := io.Copy(stdout, shell.Conn); err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
			}
		}()

		if _, err := io.Copy(shell.Conn, stdin); err != nil {
			errChan <- err
		}
	}()
	s.conn = shell
	return <-errChan
}

func (s *Docker) createShell() (*types.HijackedResponse, error) {
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
	return s.conn.Conn.Close()
}

func (s *Docker) Resize(rows, cols uint) error {
	return s.client.ContainerExecResize(context.Background(), s.ContainerID,
		types.ResizeOptions{
			Height: rows,
			Width:  cols,
		})
}
