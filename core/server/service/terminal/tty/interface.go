package tty

import (
	"fmt"
	"io"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/handler"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"
)

type Server interface {
	Connect(stdout io.Writer, stdin io.Reader) error
	Resize(rows, cols uint) error

	io.Closer
}

type NewFunc func(target *target.Target, ptyHandler handler.Handler) (Server, error)

var cacheCreator = map[target.ResourceType]NewFunc{}

func Register(resourceType target.ResourceType, f NewFunc) {
	cacheCreator[resourceType] = f
}

func New(target *target.Target, ptyHandler handler.Handler) (Server, error) {
	f, ok := cacheCreator[target.ResourceType]
	if !ok {
		return nil, fmt.Errorf("resource type %v not support", target.ResourceType)
	}
	return f(target, ptyHandler)
}
