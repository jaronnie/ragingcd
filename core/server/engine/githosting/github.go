package githosting

import (
	"context"

	"github.com/jaronnie/ragingcd/core/pkg/restc"
)

type Github struct {
	Config Config
	Client restc.Interface
}

func (g *Github) VerifyToken() error {
	return g.Client.Get().SubPath("/user").Do(context.Background()).Error()
}
