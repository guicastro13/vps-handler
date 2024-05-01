package types

import (
	"golang.org/x/crypto/ssh"
)

type SSHSession struct {
    conn *ssh.Client
}

var SSHTotalSession *SSHSession
