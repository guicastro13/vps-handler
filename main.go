package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/ssh"
)

type LoginRequest struct {
	Host     string `json:"host"`
	Ip       string `json:"ip"`
	Password string `json:"password"`
}

type SSHClient struct {
	ID       string
	Host     string
	Ip       string
	Password string
	conn     *ssh.Client
}

var activeConnections map[string]*SSHClient

func init() {
	activeConnections = make(map[string]*SSHClient)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var reqData LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := &SSHClient{
		ID:       uuid.New().String(),
		Host:     reqData.Host,
		Ip:       reqData.Ip,
		Password: reqData.Password,
	}

	if err := client.Start(); err != nil {
		log.Printf("Erro ao conectar via SSH: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("false"))
		return
	}

	jsonResponse, err := json.Marshal(map[string]string{"id": client.ID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	activeConnections[client.ID] = client
}

func HandleMachineInfo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	client, ok := activeConnections[id]
	if !ok {
		http.Error(w, "Conexão não encontrada", http.StatusNotFound)
		return
	}

	if client.conn == nil {
		http.Error(w, "Conexão não estabelecida", http.StatusInternalServerError)
		return
	}

	output, err := ExecuteSSHCommand("uname -a", client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func ExecuteSSHCommand(command string, client *SSHClient) (string, error) {
	if client.conn == nil {
		return "", errors.New("conexão SSH não estabelecida")
	}

	session, err := client.conn.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (c *SSHClient) Start() error {
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

func (c *SSHClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func HandleListConnections(w http.ResponseWriter, r *http.Request) {
	var connectionIDs []string
	for id := range activeConnections {
		connectionIDs = append(connectionIDs, id)
	}

	response, err := json.Marshal(connectionIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/machine-info", CorsMiddleware(http.HandlerFunc(HandleMachineInfo)))
	http.Handle("/login", CorsMiddleware(http.HandlerFunc(HandleLogin)))
	http.Handle("/list-connections", CorsMiddleware(http.HandlerFunc(HandleListConnections)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
