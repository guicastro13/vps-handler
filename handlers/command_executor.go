package handlers

import (
    "errors"
    "github.com/guicastro13/vps-handler/types"
)


func ExecuteSSHCommand(command string, session *types.SSHSession) (string, error) {
    if session == nil || session.conn == nil {
        return "", errors.New("SSH session not established")
    }

    SSHTotalSession, err := session.conn.NewSession() // Renomeado para SSHTotalSession para evitar conflito de nomes
    if err != nil {
        return "", err
    }
    defer SSHTotalSession.Close()

    output, err := SSHTotalSession.CombinedOutput(command)
    if err != nil {
        return "", err
    }

    return string(output), nil
}