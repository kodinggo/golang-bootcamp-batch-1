package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Eat() {
	fmt.Printf("%s is eating\n", a.Name)
}

type Dog struct {
	Animal // Menyematkan struct Animal dalam struct Dog
	Breed  string
}

func main() {
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Labrador",
	}

	// Mengakses field struct sendiri
	fmt.Println(dog.Breed)

	// Mengakses field & method dari struct Animal
	fmt.Println(dog.Name)

	dog.Eat()
}
