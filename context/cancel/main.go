package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx) // Inisialisasi context cancel

	go func(ctx context.Context) { // Menjalankan goroutine dengan context
		for {
			if isContextDone(ctx) { // Mengecek apakah fungsi cancel telah dipanggil, jika iya maka stop proses
				break
			}

			fmt.Printf("Doing some work\n")
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Canceling the parent context after 5 seconds")
		cancel() // Fungsi cancel dipanggil setelah 5 detik
	})

	var s string
	fmt.Scan(&s)
}

func isContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done(): // Check apakah context cancel telah ditrigger
		fmt.Printf("task is done, err: %v\n", ctx.Err())
		return true
	default:
		return false
	}
}
