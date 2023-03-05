package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fprintf, err := fmt.Fprintf(w, "Hello world from GfG")
		if err != nil {
			return
		}
		fmt.Println(fprintf)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fprintf, err := fmt.Fprintf(w, "Hi")
		if err != nil {
			return
		}
		fmt.Println(fprintf)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fprintf, err := fmt.Fprintf(w, "hello this is the first thing I tried calling the http services")
		if err != nil {
			return
		}
		fmt.Println(fprintf)
	})

	port := ":5001"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))

}