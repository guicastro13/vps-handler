package main

import (
	"encoding/json"
	"net/http"
	"log"
)


type LoginRequest struct {
	Host     string `json:"host"`
	Ip       string `json:"ip"`
	Password string `json:"password"`
}


func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var reqData LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := &SSHClient{
		Host:     reqData.Host,
		Ip:       reqData.Ip,
		Password: reqData.Password,
	}

	if err := client.Connect(); err != nil {
		log.Printf("Erro ao conectar via SSH: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("false"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("true"))

	sshSession = &SSHSession{
		conn: client.conn,
	}
}