// https://projecteuler.net/problem=56
// fixed this using math/big
// also added goroutine usage for pro points

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

var sums = make(chan int)

func main() {
	var mds int
	for a := 90; a < 100; a++ {
		for b := 90; b < 100; b++ {
			go func(a, b int) {
				sums <- digitalSum(a, b)
			}(a, b)
		}
	}
	for i := 0; i < 100; i++ {
		mds = max(mds, <-sums)
	}
	fmt.Printf("Answer is: %d", mds)
}

func digitalSum(a, b int) int {
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	power := new(big.Int)
	power.Exp(bigA, bigB, big.NewInt(0))
	str := power.Text(10)
	var sum int
	for _, x := range str {
		num, _ := strconv.Atoi(string(x))
		sum += num
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
