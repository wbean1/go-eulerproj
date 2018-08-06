package main

import "fmt"

func main() {
	d := Difference(100)
	fmt.Printf("difference is: %d\n", d)
}

// SquareOfSums computes (1 + 2 + ... + n)^2
func SquareOfSums(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum * sum
}

// SumOfSquares computes 1^2 + 2^2 + ... + n^2
func SumOfSquares(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum += i * i
	}
	return sum
}

// Difference computes SquareOfSums(n) - SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
