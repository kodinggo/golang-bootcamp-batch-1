package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write status code
	w.WriteHeader(http.StatusOK)

	// Write JSON response
	w.Write([]byte(`{"message": "Hello, World!"}`))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
