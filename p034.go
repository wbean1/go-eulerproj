package main

import (
	"fmt"
	"strconv"
)

func main() {
	myInts := []int{}
	for i := 10; i <= 99999999; i++ {
		iStr := strconv.Itoa(i)
		sum := 0
		for _, char := range iStr {
			sum += int(factorial(int(char - '0')))
		}
		if sum == i {
			myInts = append(myInts, i)
			fmt.Printf("found %d\n", i)
		}
	}
	sum := 0
	for _, num := range myInts {
		sum += num
	}
	fmt.Println(sum)
}

func factorial(n int) uint64 {
	answer := uint64(1)
	for i := 1; i <= n; i++ {
		answer *= uint64(i)
	}
	return answer
}
