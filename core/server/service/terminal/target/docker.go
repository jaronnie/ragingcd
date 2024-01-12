package target

import (
	"github.com/jaronnie/ragingcd/core/server/domain/bo"
)

const Docker ResourceType = "docker"

func init() {
	Register(string(Docker), newDocker)
}

type DockerExec struct{}

func newDocker() (Interface, error) {
	return &DockerExec{}, nil
}

func (r *DockerExec) Find(bo bo.ExecTerminalBo) (*Target, error) {
	var containerID string
	containerID = "6f4ae812cd81"
	return &Target{
		ResourceType: Docker,
		ResourceID:   bo.ResourceID,
		WorkingDir:   "",
		WinSize: WinSize{
			Rows: bo.Rows,
			Cols: bo.Cols,
		},
		Container: Container{
			ContainerID: &containerID,
		},
	}, nil
}
