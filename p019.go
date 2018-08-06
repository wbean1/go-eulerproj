package main

import (
	"fmt"
	"time"
)

func main() {
	var num int
	days, _ := time.ParseDuration("168h") // represents a week
	for curr := time.Date(1901, time.January, 6, 12, 0, 0, 0, time.UTC); curr.Before(time.Date(2000, time.December, 31, 12, 0, 0, 0, time.UTC)); curr = curr.Add(days) {
		if curr.Day() == 1 {
			if curr.Weekday() == 0 {
				num++
			}
		}
	}
	fmt.Printf("I found %d Sundays which were first-of-month", num)
}
