package main

import (
	"fmt"
	"math/big"
)

func main() {
	results := make(map[string]int)
	for a := 2; a < 101; a++ {
		for b := 2; b < 101; b++ {
			bigA := big.NewInt(int64(a))
			bigB := big.NewInt(int64(b))
			c := new(big.Int)
			c.Exp(bigA, bigB, nil)
			results[c.String()]++
		}
	}
	fmt.Printf("answer is: %d", len(results))
}
