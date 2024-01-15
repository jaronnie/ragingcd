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

type RemoteSsh struct {
	target     *target.Target
	handler    handler.Handler
	errChan    chan error
	doneChan   chan struct{}
	client     *ssh.Client
	session    *ssh.Session   //ssh会话
	stdinPipe  io.WriteCloser //标准输入管道
	stdoutPipe io.Reader      //标准输入管道
}

func NewRemoteSsh(target *target.Target, ptyHandler handler.Handler) (TTY, error) {
	return &RemoteSsh{
		target:   target,
		handler:  ptyHandler,
		errChan:  make(chan error),
		doneChan: make(chan struct{}),
	}, nil
}

func (r RemoteSsh) Connect(stdout io.Writer, stdin io.Reader) error {
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
	r.session = session

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          1, // Enable echoing
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	r.stdinPipe, _ = session.StdinPipe()
	r.stdoutPipe, _ = session.StdoutPipe()

	// Request pseudo-terminal
	if err := session.RequestPty("xterm", 80, 60, modes); err != nil {
		return err
	}

	// Shell starts a login shell on the remote host
	if err := session.Shell(); err != nil {
		return err
	}

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		// Copy data from stdin to the session's stdin
		_, err = io.Copy(stdout, r.stdoutPipe)
		errChan <- err
	}()

	go func() {
		// Obtain the reader for the session's stdout
		// Copy data from stdoutReader to stdout
		_, err = io.Copy(r.stdinPipe, stdin)
		errChan <- err
	}()
	return <-errChan
}

func (r RemoteSsh) Resize(rows, cols uint) error {
	return nil
}

func (r RemoteSsh) Close() error {
	return r.client.Close()
}
