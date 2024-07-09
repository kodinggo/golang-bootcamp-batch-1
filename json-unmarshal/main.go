package main

import (
	"encoding/json"
	"time"
)

type Ibu struct {
	Address string
}

type Person struct {
	Name      string    `json:"name"`
	Age       int       `json:"age,omitempty"`
	Email     string    `json:"email"`
	Verified  bool      `json:"-"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	DataIbu   Ibu       `json:"dataIbu"`
}

func main() {
	jsonString := `{"age":20,"name":"John"}`

	person := Person{}

	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		panic(err)
	}

	person.Name = person.Name + " " + " Doe"

	println("Name: ", person.Name)
	println("Age: ", person.Age)
}
