package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Setup our server config
	port := "8081"
	host := "localhost"

	// Add handlers for endpoints using callbacks
	http.HandleFunc("/", homepage)
	http.HandleFunc("/second", secondpage)

	// Start the server
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
}
func secondpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the second page")
}