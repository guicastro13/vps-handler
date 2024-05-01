package commands

import (
	"encoding/json"
	"net/http"
	"github.com/guicastro13/vps-handler/types"
	"github.com/guicastro13/vps-handler/handlers"
)


func HandleMachineInfo(w http.ResponseWriter, r *http.Request) {
	// Verifica se há uma sessão SSH estabelecida
	if sshSession == nil || sshSession.conn == nil {
		http.Error(w, "Sessão SSH não estabelecida", http.StatusInternalServerError)
		return
	}

	// Executa o comando SSH para obter as informações da máquina
	output, err := ExecuteSSHCommand("uname -a", sshSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna a saída do comando SSH como resposta HTTP
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}