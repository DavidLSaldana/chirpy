package main

import (
	"log"
	"net/http"
)

func main() {
	filepathRoot := "/"
	filepathApp := "/app/"
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	//need to update the root dir line to /app/ instead of /
	mux.Handle(filepathRoot, http.FileServer(http.Dir(".")))

	mux.Handle(filepathRoot+"assets", http.FileServer(http.Dir("./assets")))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK")) // not sure if this is okay to do this way
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)

	log.Fatal(server.ListenAndServe())

}
