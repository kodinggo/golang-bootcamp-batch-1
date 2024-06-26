package main

import "fmt"

func main() {
	total := addNumbers(10, 10)
	fmt.Println("TOTAL RESULT: ", total)
}

func addNumbers(a, b int) (total int) {
	total = a + b
	return
}
