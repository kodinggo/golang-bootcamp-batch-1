package main

import "fmt"

func main() {
	sayHello("Bagus", "Good Morning", 20)
}

func sayHello(name, message string, age float32) {
	fmt.Printf("Hello %s, %s, your age is %.1f\n", name, message, age)
}
