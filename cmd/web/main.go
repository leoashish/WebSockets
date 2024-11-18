package main

import (
	"log"
	"net/http"

	"github.com/leoashish/websockets/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listner!!!")
	go handlers.ListenToWsChannel()
	log.Println("Starting the web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
