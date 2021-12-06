package main

import "fmt"

var letters = map[int]int{
	1:  3,
	2:  3,
	3:  5,
	4:  4,
	5:  4,
	6:  3,
	7:  5,
	8:  5,
	9:  4,
	10: 3,
	11: 6,
	12: 6,
	13: 8,
	14: 8,
	15: 7,
	16: 7,
	17: 9,
	18: 8,
	19: 8,
	20: 6,
	30: 6,
	40: 5,
	50: 5,
	60: 5,
	70: 7,
	80: 6,
	90: 6,
}

func main() {
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += letterCountInt(i)
	}
	fmt.Printf("sum is %d\n", sum)
}

func letterCountInt(i int) int {
	if letters[i] != 0 {
		return letters[i]
	} else if i == 1000 {
		return letterCountInt(1) + letterCountStr("thousand")
	} else if i%100 == 00 {
		return letterCountInt(i/100) + letterCountStr("hundred")
	} else if i > 99 && i < 1000 {
		return letterCountInt(i/100) + letterCountStr("hundred") + letterCountStr("and") + letterCountInt(i%100)
	} else {
		return letterCountInt((i/10)*10) + letterCountInt(i%((i/10)*10))
	}
	return 0
}

func letterCountStr(s string) int {
	if s == "thousand" {
		return 8
	}
	if s == "hundred" {
		return 7
	}
	if s == "and" {
		return 3
	}
	return 0
}
