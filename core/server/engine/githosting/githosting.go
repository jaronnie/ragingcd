package githosting

import "github.com/jaronnie/ragingcd/core/pkg/restc"

const (
	GITHUB = "github"
	GITLAB = "gitlab"
	GITEA  = "gitea"
)

type Config struct {
	Type     string
	Url      string
	Username string
	Token    string
}

type Interface interface {
	VerifyToken() error
}

func New(config Config) Interface {
	switch config.Type {
	case GITHUB:
		return &Github{config}
	case GITLAB:
		return &Gitlab{config}
	case GITEA:
		return &Gitea{config}
	default:
		return &Github{config}
	}
}

func newClient(config Config) (restc.Interface, error) {
	restClient, _ := restc.New(restc.WithUrl(config.Url))
	return restClient, nil
}
