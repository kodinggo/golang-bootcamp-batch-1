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
	person := Person{
		Name:     "John",
		Age:      20,
		Email:    "email@example",
		Password: "password",
		Verified: true,
	}

	byteJson, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	println("RESULT JSON: ", string(byteJson))
}
