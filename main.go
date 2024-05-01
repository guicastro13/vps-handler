package main

import (
	"log"
	"net/http"
    "github.com/guicastro13/vps-handler/commands"
    "github.com/guicastro13/vps-handler/middlewares"
    "github.com/guicastro13/vps-handler/types"
    "github.com/guicastro13/vps-handler/handlers"
)

func main() {
	http.Handle("/machine-info", CorsMiddleware(http.HandlerFunc(HandleMachineInfo)))
	http.Handle("/login", CorsMiddleware(http.HandlerFunc(HandleLogin)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}