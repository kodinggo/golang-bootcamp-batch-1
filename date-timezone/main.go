package main

import (
	"fmt"
	"time"
)

func main() {
	timeZone, err := time.LoadLocation("Australia/Perth")
	if err != nil {
		fmt.Println(err)
		return
	}

	now := time.Now().In(timeZone)
	fmt.Println("Time in Perth:", now)
}
