package tty

import (
	"io"

	"github.com/pkg/errors"

	"github.com/jaronnie/ragingcd/core/server/domain/po"
	"github.com/jaronnie/ragingcd/core/server/engine/db"
	"github.com/spf13/cast"

	"github.com/jaronnie/ragingcd/core/pkg/sshx"

	"golang.org/x/crypto/ssh"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/handler"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"
)

var (
	ErrTargetNotFound = errors.New("not found")
)

const (
	ShellTypeRemoteSsh target.ResourceType = "ssh"
)

func init() {
	Register(ShellTypeRemoteSsh, NewRemoteSsh)
}

type RemoteSshShell struct {
	target   *target.Target
	handler  handler.Handler
	errChan  chan error
	doneChan chan struct{}
	client   *ssh.Client
}

func NewRemoteSsh(target *target.Target, ptyHandler handler.Handler) (Server, error) {
	return &RemoteSshShell{
		target:   target,
		handler:  ptyHandler,
		errChan:  make(chan error),
		doneChan: make(chan struct{}),
	}, nil
}

func (r RemoteSshShell) Connect(stdout io.Writer, stdin io.Reader) error {
	var err error
	sshPo := po.Ssh{
		Base: po.Base{ID: cast.ToInt(r.target.ResourceID)},
	}
	b, err := db.Engine.Get(&sshPo)
	if !b {
		return ErrTargetNotFound
	}
	if err != nil {
		return err
	}
	r.client, err = sshx.NewSshClient(&sshx.Config{
		IP:         sshPo.IP,
		Port:       sshPo.Port,
		Username:   sshPo.Username,
		Type:       sshPo.Type,
		Password:   sshPo.Password,
		PrivateKey: sshPo.PrivateKey,
		Timeout:    10,
	})
	if err != nil {
		return err
	}
	session, err := r.client.NewSession()
	if err != nil {
		r.client.Close()
		return err
	}
	defer session.Close()
	errChan := make(chan error)
	defer close(errChan)
	err = session.Shell()

	go func() {
		// Obtain the reader for the session's stdout
		// Copy data from stdoutReader to stdout
		_, err = io.Copy(stdout, session.Stdin)
		if err != nil {
			r.errChan <- err
			return
		}

	}()

	go func() {
		// Copy data from stdin to the session's stdin
		_, err = io.Copy(session.Stdout, stdin)
		r.errChan <- err
	}()
	return <-errChan
}

func (r RemoteSshShell) Resize(rows, cols uint) error {
	return nil
}

func (r RemoteSshShell) Close() error {
	return r.client.Close()
}
