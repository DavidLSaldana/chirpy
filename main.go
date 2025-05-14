package main

import (
	"log"
	"net/http"
)

func main() {
	filepathRoot := "/"
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	mux.Handle(filepathRoot, http.FileServer(http.Dir(".")))
	mux.Handle(filepathRoot+"assets", http.FileServer(http.Dir("./assets")))

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)

	log.Fatal(server.ListenAndServe())

}
