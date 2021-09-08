package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jtheo/milestone1-code/storage"
)

// Run is the starting point to server the webserver
// It requires a variable ids of type storage.Storage
func Run(ids storage.Storage) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", ids.SecretHandler)
	mux.HandleFunc("/healthcheck", healthCheckHandler)

	addr := "localhost:8080"
	log.Println("Starting on", addr)
	log.Panic(http.ListenAndServe(addr, mux))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
