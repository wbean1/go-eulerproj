package main

import (
	"fmt"
	"strconv"
)

func main() {
	lychrels := []int{}
	for i := 1; i < 10000; i++ {
		next, _ := strconv.Atoi(reverse(strconv.FormatInt(int64(i), 10)))
		next += i
		if !isPalindromeYet(next, 49) {
			lychrels = append(lychrels, i)
		}
	}
	fmt.Print(lychrels)
	fmt.Printf("answer is %d\n", len(lychrels))
}

func isPalindromeYet(number, iterations int) bool {
	str := strconv.FormatInt(int64(number), 10)
	if str == reverse(str) {
		return true
	}
	if iterations == 1 {
		return false
	}
	next, _ := strconv.Atoi(reverse(str))
	next += number
	return isPalindromeYet(next, iterations-1)
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
