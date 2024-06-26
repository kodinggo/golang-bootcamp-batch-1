package main

import "fmt"

func greet(name string, argument func(string)) {
	argument(name)
}

func calculation(numbers []int, callback func([]int)) {
	callback(numbers)
}

func find(numbers []int, isFive func(int) bool) {
	for _, number := range numbers {
		if isFive(number) {
			fmt.Println("Betul, Numbernya adalah ", number)
			return
		}
	}

	fmt.Println("Tidak ada number yang sesuai")
}

func main() {
	greetings := func() {
		fmt.Println("Hello World")
	}

	greetings()

	greet("Bagus", func(name string) {
		fmt.Println("Hello", name)
	})

	calculation([]int{1, 2, 3, 4, 5}, func(numbers []int) {
		var total int
		for _, number := range numbers {
			total += number
		}

		fmt.Println("CALCULATION 1", total)
	})

	calculation([]int{1, 2, 3, 4, 5}, func(numbers []int) {
		var total int
		for _, number := range numbers {
			total -= number
		}

		fmt.Println("CALCULATION 1", total)
	})

	checkfunc := func(number int) bool {
		return number == 5
	}

	find([]int{1, 2, 3, 4, 5}, checkfunc)
}
