package handlers

import (
    "errors"
    "golang.org/x/crypto/ssh"
)

func ExecuteSSHCommand(command string, session *SSHSession) (string, error) {
    if session == nil || session.conn == nil {
        return "", errors.New("SSH session not established")
    }

    sshSession, err := session.conn.NewSession() // Renomeado para sshSession para evitar conflito de nomes
    if err != nil {
        return "", err
    }
    defer sshSession.Close()

    output, err := sshSession.CombinedOutput(command)
    if err != nil {
        return "", err
    }

    return string(output), nil
}