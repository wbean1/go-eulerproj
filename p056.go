// https://projecteuler.net/problem=56
// this is wrong and I have no clue why
// 93^99 seems to give the maximum digital sum of 955...
// also was a good practice in go type hell
// (yes, i converted a float -> string -> []string -> []float)

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	mds := float64(0)
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			power := math.Pow(float64(a), float64(b))
			fmt.Printf("%d to the %d is: %.f\n", a, b, power)
			ds := digitalSum(power)
			fmt.Printf("digitalsum is: %.f\n\n", ds)
			mds = max(mds, ds)
		}
	}
	fmt.Printf("Answer is: %.f", mds)
}

func digitalSum(a float64) float64 {
	str := strings.Split(strconv.FormatFloat(a, 'f', 0, 64), "") // this makes splice of single digit strings?!
	sum := float64(0)
	for x := range str {
		num, err := strconv.ParseFloat(str[x], 64)
		if err != nil {
			panic("parse error")
		}
		fmt.Printf("%.f +", int)
		sum = sum + num
	}
	fmt.Printf("= %.f", sum)
	return sum
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
