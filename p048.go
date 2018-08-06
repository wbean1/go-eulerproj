package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := new(big.Int)
	for i := 1; i <= 1000; i++ {
		bigI := big.NewInt(int64(i))
		sum.Add(sum, bigI.Exp(bigI, bigI, nil))
	}
	fmt.Printf("answer is: %s", sum.String())
}
