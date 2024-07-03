package main

import (
	"fmt"
)

func main() {
	totalBuffer := 5
	chann := make(chan int, totalBuffer)

	go func() {
		for i := 0; i < totalBuffer; i++ {
			chann <- i
			fmt.Printf("Hello from channel %d\n", i+1)
		}

		close(chann)
	}()

	for i := range chann {
		fmt.Printf("Received from channel %d\n", i)
	}
}
