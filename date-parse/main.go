package main

import (
	"fmt"
	"time"
)

const DateFormatLocal = "02 January 2006"

func main() {
	// parsing format should be same like the time string
	// now := "2024-06-29 13:45:30"
	// format := "2006-01-02 15:04:05"
	now := "June, 29 2024"
	format := "January, 02 2006"

	resultTime, err := time.Parse(format, now)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	fmt.Println("Result parsed time:", resultTime)
}
