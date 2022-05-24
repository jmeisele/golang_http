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
	fileServer := http.FileServer(http.Dir("./static"))
	
	// Add handlers for endpoints using callbacks
	http.Handle("/", fileServer)
	http.HandleFunc("/second", secondpage)
	http.HandleFunc("/form", formHandler)

	// Start the server
	fmt.Println("Booting server at port", port)
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func secondpage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/second" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Welcome to the second page")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}