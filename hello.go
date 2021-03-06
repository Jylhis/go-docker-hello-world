package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Server starting")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)

	fmt.Println("Shutting down...")
}
