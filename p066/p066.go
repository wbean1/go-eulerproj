package main

import (
	"fmt"
	"math"
)

func main() {
	answers := make(map[int]int)
	for D := 1; D <= 1000; D++ {
		answers[D] = findMinX(D)
		fmt.Printf("found x=%d for D=%d\n", answers[D], D)
	}
	fmt.Print(answers)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMinX(D int) int {
	if square_check(D) {
		return 0
	}
	for x := 2; x < 1000000; x++ {
		for y := 2; y < 1000000; y++ {
			if floatEquals(float64(x), math.Sqrt(float64(1+D*y*y))) {
				return x
			}
		}
	}
	return 1000000
}

var eps float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}

func square_check(a int) bool {
	var int_root int = int(math.Sqrt(float64(a)))
	return (int_root * int_root) == a
}
