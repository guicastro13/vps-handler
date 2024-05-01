package types

import (
	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Host     string
	Ip       string
	Password string
	conn     *ssh.Client
}