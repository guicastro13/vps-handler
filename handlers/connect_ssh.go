package handlers

import (
	"golang.org/x/crypto/ssh"
	"github.com/guicastro13/vps-handler/types"
)

func (c *types.SSHClient) Connect() error {
	clientConfig := &ssh.ClientConfig{
		User: c.Host,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", c.Ip+":22", clientConfig)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}