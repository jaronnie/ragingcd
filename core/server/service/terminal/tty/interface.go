package tty

import (
	"fmt"
	"io"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/handler"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"
)

var (
	cmd = []string{"/bin/sh"}
)

type TTY interface {
	Connect(*handler.WsHandler) error
	Resize(rows, cols uint) error

	io.Closer
}

type NewFunc func(target *target.Target, wsHandler handler.WsHandler, ptyHandler handler.Handler) (TTY, error)

var cacheCreator = map[target.ResourceType]NewFunc{}

func Register(resourceType target.ResourceType, f NewFunc) {
	cacheCreator[resourceType] = f
}

func New(target *target.Target, wsHandler handler.WsHandler, ptyHandler handler.Handler) (TTY, error) {
	f, ok := cacheCreator[target.ResourceType]
	if !ok {
		return nil, fmt.Errorf("resource type %v not support", target.ResourceType)
	}
	return f(target, wsHandler, ptyHandler)
}
