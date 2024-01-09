package githosting

import (
	"context"

	"github.com/jaronnie/ragingcd/core/pkg/restc"
)

type Gitea struct {
	Config Config
}

func (g *Gitea) VerifyToken() error {
	restClient, _ := newClient(g.Config)
	return restClient.Get().SubPath("/api/v1/user").Params(restc.QueryParam{
		Name:  "access_token",
		Value: g.Config.Token,
	}).Do(context.Background()).Error()
}
