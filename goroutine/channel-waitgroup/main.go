package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	var ch = make(chan string)

	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		ch <- "Hello from goroutine 1"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		ch <- "Hello from goroutine 2"
	}()

	go func() {
		defer wg.Done()

		time.Sleep(3 * time.Second)
		ch <- "Hello from goroutine 3"
	}()

	go func() {
		fmt.Println("Waiting for goroutines to finish")
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
