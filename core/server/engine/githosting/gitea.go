package githosting

import (
	"context"

	"github.com/jaronnie/ragingcd/core/pkg/restc"
)

type Gitea struct {
	Config Config
	Client restc.Interface
}

func (g *Gitea) VerifyToken() error {
	return g.Client.Get().SubPath("/api/v1/user").Do(context.Background()).Error()
}
