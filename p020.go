package main

import (
	"fmt"
	"math/big"

	"github.com/wbean1/go-eulerproj/pkg/utils"
)

func main() {
	var factorial big.Int
	factorial.MulRange(1, 100)
	factorialStr := factorial.String()
	sumOfFactorialDigits := utils.SumOfDigitsInString(factorialStr)
	fmt.Printf("factorial is %s and sum of digits is: %d\n", factorialStr, sumOfFactorialDigits)
}
