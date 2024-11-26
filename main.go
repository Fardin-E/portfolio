package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve files from the ./static directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
