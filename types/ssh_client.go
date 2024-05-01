package types

type SSHClient struct {
	Host     string
	Ip       string
	Password string
	conn     *ssh.Client
}