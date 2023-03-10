package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the static assets generated by the React app
	http.Handle("/", http.FileServer(http.Dir("./")))

	// Start the server
	log.Println("runserver http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
