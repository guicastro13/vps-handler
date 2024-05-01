package main

import (
	"log"
	"net/http"
    "commands"
    "middlewares"
    "types"
    "handlers"
)

func main() {
	http.Handle("/machine-info", CorsMiddleware(http.HandlerFunc(HandleMachineInfo)))
	http.Handle("/login", CorsMiddleware(http.HandlerFunc(HandleLogin)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}