package main

import (
	"fmt"
	"io"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Membaca body request
	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(string(byteBody))
	io.WriteString(w, "success")
}

func main() {
	http.HandleFunc("/form", formHandler)
	http.ListenAndServe(":8080", nil)
}
