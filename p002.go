package main

import "fmt"

func main() {
	fibs := []int{1, 2}
	sum := 2
	for i := 3; i < 4000000; i = fibs[len(fibs)-1] + fibs[len(fibs)-2] {
		fibs = append(fibs, i)
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Print(fibs)
	fmt.Printf("sum is %d", sum)
}
