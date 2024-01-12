package target

import (
	"github.com/jaronnie/ragingcd/core/server/domain/bo"
)

const SSH ResourceType = "ssh"

func init() {
	Register(string(SSH), newRemoteSsh)
}

type RemoteSsh struct{}

func newRemoteSsh() (Interface, error) {
	return &RemoteSsh{}, nil
}

func (r *RemoteSsh) Find(bo bo.ExecTerminalBo) (*Target, error) {
	return &Target{
		ResourceType: SSH,
		ResourceID:   bo.ResourceID,
		WorkingDir:   "",
		WinSize: WinSize{
			Rows: bo.Rows,
			Cols: bo.Cols,
		},
	}, nil
}
