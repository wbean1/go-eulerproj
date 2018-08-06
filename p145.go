package main

import (
	"fmt"
	"strconv"
)

func main() {
	var count int
	for i := 1; i < 1000000000; i++ {
		if i%10 == 0 {
			// skip since reverse will have leading zero
		} else {
			backwardsInt, _ := strconv.Atoi(reverse(strconv.FormatInt(int64(i), 10)))
			if isAllOddDigits(i + backwardsInt) {
				count++
			}
		}
	}
	fmt.Printf("answer is: %d\n", count)
}

func reverse(s string) string {
	orig := []rune(s)
	size := len(orig)
	rev := make([]rune, size)
	for i, r := range orig {
		rev[size-i-1] = r
	}
	return string(rev)
}

func isAllOddDigits(i int) bool {
	for i = i; i > 1; i = i / 10 {
		if i%2 == 0 {
			return false
		}
	}
	return true
}
