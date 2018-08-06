package main

import "fmt"

func main() {
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			for c := 1; c < 1000; c++ {
				if a+b+c == 1000 {
					if isTriplet(a, b, c) {
						fmt.Printf("answer is: %d\n", a*b*c)
					}
				}
			}
		}
	}
}

func isTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}
