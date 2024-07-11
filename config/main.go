package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // load env
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println(os.Getenv("DB_USER"))
	fmt.Println(os.Getenv("DB_PASSWORD"))
	fmt.Println(os.Getenv("REDIS"))
}
