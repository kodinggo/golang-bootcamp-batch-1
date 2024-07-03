package main

import (
	"fmt"
	"time"
)

// func main() {
// 	var wg sync.WaitGroup

// 	var mu sync.Mutex

// 	var counter int
// 	// Melakukan operasi increment dg 1000 goroutines
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()

// 			mu.Lock()

// 			counter++

// 			mu.Unlock()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("Final counter:", counter) // Final counter tidak 1000
// }

func main() {
	ch := make(chan int, 1)

	// var counter int
	// Melakukan operasi increment dg 1000 goroutines
	for i := 0; i < 2; i++ {
		go func() {
			ch <- 1
			time.Sleep(time.Second)
			defer close(ch)
		}()
	}

	for i := range ch {
		fmt.Println(i)
	}

	// fmt.Println("Final counter:", counter) // Final counter tidak 1000
}
