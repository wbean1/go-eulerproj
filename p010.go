package main

import "fmt"

func main() {
	primes := []int{}
	for i := 2; i < 2000000; i++ {
		for divisor := 2; divisor <= i; divisor++ {
			if i == divisor {
				// i is prime
				primes = append(primes, i)
			}
			if i%divisor == 0 {
				// i not prime
				break
			}
		}
	}
	fmt.Println("primes:")
	fmt.Println(primes)
	sum := sum(primes)
	fmt.Printf("sum of primes: %d\n", sum)
}

func sum(list []int) int {
	i := 0
	for _, num := range list {
		i += num
	}
	return i
}
