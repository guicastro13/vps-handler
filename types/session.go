package types

type SSHSession struct {
    conn *ssh.Client
}

var sshSession *SSHSession

