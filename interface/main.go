package main

import "fmt"

type Animal interface {
	Info() string
	SetName(name string)
	SetAge(age int)
}

type Cat struct {
	Name string
	Age  int
}

func (c Cat) Info() string {
	return "Name: " + c.Name + ", Age: " + fmt.Sprint(c.Age)
}

func (c *Cat) SetName(name string) {
	c.Name = name
}

func (c *Cat) SetAge(age int) {
	c.Age = age
}

func NewCat() Animal {
	return &Cat{}
}

func main() {
	cat := NewCat()
	cat.SetName("Kitty")
	cat.SetAge(2)

	msg := cat.Info()

	fmt.Println(msg)
}
