package main

import "fmt"

// Constraints interface
type Numeric interface {
	int | int64 | int32 | float64 | float32
}

func increase[T Numeric](n T) T {
	n++
	return n
}

func main() {
	var one int = 1
	var ten int64 = 10
	var pi float64 = 3.14
	fmt.Println(increase(one)) // 2
	fmt.Println(increase(ten)) // 11
	fmt.Println(increase(pi))  // 4.14
}
