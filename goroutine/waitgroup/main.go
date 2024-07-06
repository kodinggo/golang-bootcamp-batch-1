package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Hello from goroutine 1")
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello from goroutine 2")
		wg.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Hello from goroutine 3")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("All goroutines done")
}
