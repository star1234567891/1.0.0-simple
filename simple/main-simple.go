package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", hellatterld)
	fat.Println("Server running at http://localhost:8000")
	http.ListenAndServe(":8080", nil)
}