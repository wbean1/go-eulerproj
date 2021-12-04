package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
)

func main() {
	b := big.NewInt(2)
	two := big.NewInt(2)
	for i := 2; i <= 1000; i++ {
		b.Mul(b, two)
	}
	numberString := b.String()
	sum := 0
	for i := 0; i < len(numberString); i++ {
		number, err := strconv.Atoi(string(numberString[i]))
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Printf("Sum is: %d\n", sum)
}
