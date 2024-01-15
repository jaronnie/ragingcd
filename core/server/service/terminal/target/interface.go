package target

import (
	"github.com/jaronnie/ragingcd/core/server/domain/bo"
)

type ResourceType string

type Target struct {
	ResourceType ResourceType
	ResourceID   string
	WorkingDir   string
	Rows         uint16
	Cols         uint16
}

func FindTarget(bo bo.ExecTerminalBo) (*Target, error) {
	return &Target{
		ResourceType: ResourceType(bo.ResourceType),
		ResourceID:   bo.ResourceID,
		WorkingDir:   bo.WorkingDir,
		Rows:         bo.Rows,
		Cols:         bo.Cols,
	}, nil
}
