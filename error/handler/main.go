package main

import (
	"fmt"
	"os"
)

// Fungsi divide mengambalikan nilai int & error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	a := 10
	b := 0
	result, err := divide(a, b)
	// Cek apakah variable "err" dari fungsi divide tidak nil (terjadi error)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Result: %d\n", result)
}
