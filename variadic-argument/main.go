package main

import "fmt"

func sum(numbers ...int) (int, int) {
	var totalPenjumlahan int
	var totalPengurangan int

	for _, number := range numbers {
		totalPenjumlahan += number
	}

	for _, number := range numbers {
		totalPengurangan -= number
	}

	return totalPenjumlahan, totalPengurangan
}

func main() {
	addition1, subtraction1 := sum(1, 2, 3, 4, 5)
	fmt.Println(addition1, subtraction1)

	numbers := []int{1, 2, 3, 4, 10}
	addition2, subtraction2 := sum(numbers...)
	fmt.Println(addition2, subtraction2)

	addition3, subtraction3 := sum()
	fmt.Println(addition3, subtraction3)
}
