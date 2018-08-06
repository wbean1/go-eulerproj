package main

import "fmt"

func main() {
	var longestStart int
	var longestLen int
	for i := 1; i < 1000000; i++ {
		chain := makeChain(i)
		if len(chain) > longestLen {
			longestLen = len(chain)
			longestStart = i
		}
	}
	fmt.Printf("longest start is %d with len %d", longestStart, longestLen)
}

func makeChain(i int) []int {
	if i > 1 {
		if i%2 == 0 {
			return append(makeChain(i/2), i)
		} else {
			return append(makeChain(3*i+1), i)
		}
	}
	return []int{i} // i = 1
}
