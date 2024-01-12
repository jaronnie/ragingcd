package tty

import (
	"context"
	"io"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/handler"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var (
	cmd = []string{"/bin/sh"}
)

func init() {
	Register("docker", NewDocker)
}

type DockerShell struct {
	target  *target.Target
	handler handler.Handler
	client  *client.Client
	execID  string
	conn    *types.HijackedResponse
}

func NewDocker(target *target.Target, ptyHandler handler.Handler) (Server, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}
	return NewDockerShellWithClient(target, ptyHandler, cli)
}

func NewDockerShellWithClient(target *target.Target, ptyHandler handler.Handler, client *client.Client) (Server, error) {
	shell := &DockerShell{
		target:  target,
		handler: ptyHandler,
		client:  client,
	}

	if target.WorkingDir == "" {
		c, err := client.ContainerInspect(context.Background(), *target.ContainerID)
		if err != nil {
			return nil, err
		}
		target.WorkingDir = c.Config.WorkingDir
	}

	return shell, nil
}

func (s *DockerShell) Connect(stdout io.Writer, stdin io.Reader) error {
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

func (s *DockerShell) createShell() (*types.HijackedResponse, error) {
	ctx := context.Background()
	exec, err := s.client.ContainerExecCreate(ctx, *s.target.ContainerID, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
		Tty:          true,
	})
	if err != nil {
		return nil, err
	}
	s.execID = exec.ID
	result, err := s.client.ContainerExecAttach(ctx, exec.ID, types.ExecStartCheck{Detach: false, Tty: true})
	return &result, err
}

func (s *DockerShell) Close() error {
	return s.conn.Conn.Close()
}

func (s *DockerShell) Resize(rows, cols uint) error {
	return s.client.ContainerExecResize(context.Background(), s.execID,
		types.ResizeOptions{
			Height: rows,
			Width:  cols,
		})
}
