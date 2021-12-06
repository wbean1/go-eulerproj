package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(1)
	b := big.NewInt(1)

	// Initialize limit as 10^99, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)

	count := 2
	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		count++
	}
	fibStr := a.String()

	fmt.Printf("first fib seq w/ 1000digits is %s, len(fibStr): %d, which is # %d\n", fibStr, len(fibStr), count)
}
