package main

// Simple webserver application handling 3 routes, one of which is a form.

import (
	"fmt"
	"log"
	"net/http"
)

// Main function that starts the webserver and handles 3 routes.
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// HelloHandler is a function for handling the /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

// FormHandler is a function for handling /form.html route
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}
