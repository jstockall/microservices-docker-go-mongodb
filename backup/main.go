package main

import (
	"log"
	"net/http"

	"github.com/mmorejon/cinema/backup/routers"
)

// Entry point for the program
func main() {

	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
