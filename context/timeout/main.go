package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	initialCtx := context.Background()
	ctx, cancel := context.WithTimeout(initialCtx, 5*time.Second) // Inisialisasi context dengan timeout
	defer cancel()

	go func(ctx context.Context) { // Menjalankan goroutine dengan context
		for {
			if isContextDone(ctx) { // Mengecek apakah fungsi cancel telah dipanggil, jika iya maka stop proses
				break
			}
			fmt.Printf("Doing some work\n")
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

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
