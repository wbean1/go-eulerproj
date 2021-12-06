package main

import (
	"fmt"

	"github.com/wbean1/go-eulerproj/pkg/utils"
)

func main() {
	sum := 0
	sumOfFactors := make(map[int]int)
	for i := 1; i < 10000; i++ {
		sumOfFactors[i] = utils.Sum(utils.GetDivisors(i))
	}
	for key, value := range sumOfFactors {
		if sumOfFactors[value] == key && key != value {
			fmt.Printf("number %d has pair %d\n", key, value)
			sum += key
		}
	}
	fmt.Printf("sum is %d\n", sum)
}
