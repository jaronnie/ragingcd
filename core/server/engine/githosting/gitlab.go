package githosting

import (
	"context"

	"github.com/jaronnie/ragingcd/core/pkg/restc"
)

type Gitlab struct {
	Config Config
	Client restc.Interface
}

func (g *Gitlab) VerifyToken() error {
	return g.Client.Get().SubPath("/api/v4/user").Do(context.Background()).Error()
}
