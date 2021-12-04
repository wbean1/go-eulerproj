package main

import "fmt"

func main() {
	factors := []int{}
	remaining := 600851475143
	for i := 1; i <= remaining; i++ {
		if remaining%i == 0 {
			factors = append(factors, i)
			remaining = remaining / i
		}
	}
	fmt.Println("factors are:")
	fmt.Println(factors)
}
