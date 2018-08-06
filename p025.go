package main

func main() {
	fibs := []int{1, 1, 2}
	for i := 3; i < 4000000; i = fibs[len(fibs)-1] + fibs[len(fibs)-2] {
		fibs = append(fibs, i)
}
