package main

import (
	"Go-Streaming-Server/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.FetchNewsRequestHandler)

	log.Println("Starting server on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
