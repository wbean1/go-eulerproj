package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ans int
	for i := 800; i < 1000; i++ {
		for j := 800; j < 1000; j++ {
			if i*j > ans {
				if isPalindrome(i * j) {
					fmt.Printf("found %d\n", i*j)
					ans = i * j
				}
			}
		}
	}
	fmt.Printf("answer is: %d\n", ans)
}

func isPalindrome(i int) bool {
	str := strconv.FormatInt(int64(i), 10)
	return str == reverse(str)
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
