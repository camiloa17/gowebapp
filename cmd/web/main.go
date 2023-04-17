package main

import (
	"log"
	"net/http"

	"github.com/camiloa17/gowebapp/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application handler
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Listening on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
