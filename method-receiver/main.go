package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) sayHello() {
	fmt.Println("Hello", p.Name, p.Age)
}

func (p *Person) setName(name string) {
	p.Name = "Mr. " + name
}

func (p *Person) setAge(age int) {
	p.Age = age
}

func main() {
	dataBagus := Person{}

	dataBagus.setName("Bagus")
	dataBagus.setAge(20)

	dataBagus.sayHello()

	var name string = "Bagus"
	changeName(&name)
	fmt.Println(name)
}

func changeName(name *string) {
	*name = "Mr. " + *name
}
