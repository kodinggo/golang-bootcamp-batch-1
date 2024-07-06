package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// format := "January, 02 2006"
	// strTime := now.Format(format)

	formats := []string{
		time.RFC3339, // 2006-01-02T15:04:05Z07:00
		time.RFC1123, // Mon, 02 Jan 2006 15:04:05 MST
		"2006-01-02 15:04:05",
		"02-Jan-2006 15:04",
		"01/02/2006 3:04PM",
		"02 Jan 2006",
	}

	for _, format := range formats {
		strTime := now.Format(format)
		fmt.Println("Time now: ", strTime)
	}
}
