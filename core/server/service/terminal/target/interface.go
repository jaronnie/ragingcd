package target

import (
	"fmt"

	"github.com/jaronnie/ragingcd/core/server/domain/bo"
)

// Interface to find resource: remote ssh or docker container or k8s pod
type Interface interface {
	Find(bo bo.ExecTerminalBo) (*Target, error)
}

type NewFunc func() (Interface, error)

var typeMap = map[string]NewFunc{}

func Register(typo string, f NewFunc) {
	typeMap[typo] = f
}

func New(typo string) (Interface, error) {
	f, ok := typeMap[typo]
	if !ok {
		return nil, fmt.Errorf("type %v not support", typo)
	}
	return f()
}

func FindTarget(bo bo.ExecTerminalBo) (*Target, error) {
	finder, err := New(bo.ResourceType)
	if err != nil {
		return nil, err
	}
	return finder.Find(bo)
}

type Target struct {
	ResourceType ResourceType
	ResourceName string
	ResourceID   string
	WorkingDir   string
	WinSize
	Pod
	Container
	Process
}

type ResourceType string

type Pod struct {
	Namespace *string
	PodName   *string
}

type Container struct {
	ContainerID *string
}

type Process struct {
	Pid *uint32
}

type WinSize struct {
	Rows uint16
	Cols uint16
}
