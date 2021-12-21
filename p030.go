package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	myInts := []int{}
	for i := 10; i <= 999999; i++ {
		iStr := strconv.Itoa(i)
		sum := 0
		for _, char := range iStr {
			sum += int(math.Pow(float64(int(char-'0')), 5))
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
