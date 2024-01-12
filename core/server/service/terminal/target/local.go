package target

import (
	"github.com/jaronnie/ragingcd/core/server/domain/bo"
)

const Local ResourceType = "local"

func init() {
	Register(string(Local), newLocal)
}

type LocalExec struct{}

func newLocal() (Interface, error) {
	return &LocalExec{}, nil
}

func (r *LocalExec) Find(bo bo.ExecTerminalBo) (*Target, error) {
	return &Target{
		ResourceType: Local,
		ResourceID:   bo.ResourceID,
		WorkingDir:   "",
		WinSize: WinSize{
			Rows: bo.Rows,
			Cols: bo.Cols,
		},
	}, nil
}
