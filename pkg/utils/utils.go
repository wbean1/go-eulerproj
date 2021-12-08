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

func GetFactors(num int) []int {
	factors := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func GetDivisors(num int) []int {
	factors := []int{}
	for i := 1; i < num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
