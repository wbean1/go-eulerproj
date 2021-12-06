package utils

import (
	"log"
	"strconv"
)

func SumOfDigitsInString(str string) int {
	sum := 0
	for i, _ := range str {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	return sum
}
