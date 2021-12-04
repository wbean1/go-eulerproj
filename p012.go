package main

import "fmt"

func main() {
	answer := 0
	previous := 0
	for i := 1; 1 > 0; i++ {
		tNumber := i + previous
		factors := getFactors(tNumber)
		if len(factors) > 500 {
			fmt.Printf("len of factors: %d\n", len(factors))
			answer = tNumber
			break
		}
		previous = tNumber
	}
	fmt.Printf("answer is: %d\n", answer)
}

func getFactors(num int) []int {
	factors := []int{}
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
