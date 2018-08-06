package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := new(big.Int)
	for i := 123456789; i <= 9876543210; i++ {
		if isPandigital(i) {
			//fmt.Printf("%d is pandigital\n", i)
			if hasProperty(i) {
				fmt.Printf("%d has property\n", i)
				bigI := big.NewInt(int64(i))
				sum.Add(sum, bigI)
			}
		}
	}
	str := sum.Text(10)
	fmt.Printf("sum is %s\n", str)
}

func isPandigital(i int) bool {
	digits := make(map[int]int)
	for j := i; j > 0; j = j / 10 {
		digit := j % 10
		if _, exists := digits[digit]; exists {
			return false
		}
		digits[digit]++
	}
	if i <= 987654321 { // also need to check that there is no 0 in these
		if _, exists := digits[0]; exists {
			return false
		}
	}
	return true
}

func hasProperty(i int) bool {
	digits := []int{}
	for j := i; j > 0; j = j / 10 {
		digits = append([]int{j % 10}, digits...)
	}
	if i <= 987654321 { // need to prepend a zero
		digits = append([]int{0}, digits...)
	}
	if ((digits[1]*100+digits[2]*10+digits[3])%2 == 0) &&
		((digits[2]*100+digits[3]*10+digits[4])%3 == 0) &&
		((digits[3]*100+digits[4]*10+digits[5])%5 == 0) &&
		((digits[4]*100+digits[5]*10+digits[6])%7 == 0) &&
		((digits[5]*100+digits[6]*10+digits[7])%11 == 0) &&
		((digits[6]*100+digits[7]*10+digits[8])%13 == 0) &&
		((digits[7]*100+digits[8]*10+digits[9])%17 == 0) {
		return true
	}
	return false
}
