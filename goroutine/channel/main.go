package main

import "fmt"

func sendData(ch chan<- string) {
	// sending string to ch
	ch <- "Hello from channel"
}

func receivedData(ch <-chan string) string {
	// receiving data from ch
	data := <-ch

	return data
}

func main() {
	ch := make(chan string)

	go sendData(ch)

	data := receivedData(ch)
	fmt.Println(data)
}
