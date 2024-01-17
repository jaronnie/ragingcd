package sshx

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

const (
	Password   = "password"
	PrivateKey = "private_key"
)

type Config struct {
	IP         string
	Port       int
	Username   string
	Type       string
	Password   string
	PrivateKey string
	Timeout    int
}

func NewSshClient(config *Config) (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User:            config.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(config.Timeout) * time.Second,
	}
	if config.Type == PrivateKey {
		privateKey, err := ssh.ParsePrivateKey([]byte(config.PrivateKey))
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		}
	} else if config.Type == Password {
		sshConfig.Auth = []ssh.AuthMethod{
			ssh.Password(config.Password),
		}
	}

	// Connect to host
	addr := fmt.Sprintf("%s:%d", config.IP, config.Port)
	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "dail to addr [%s]", addr)
	}

	return client, nil
}
