package main

var letters := map[int]int{
	1: 3,
	2: 3,
	3: 5,
	4: 4,
	5: 4,
	6: 3,
	7: 5,
	8: 5,
	9: 4,
	10: 3,
	11: 6,
	12: 6,
	13: 8,
	14: 8,
	15: 7,
	16: 7,
	17: 9,
	18: 8,
	19: 8,
	20: 6,
	30: 6,
	40: 5,
	50: 5,
	60: 5,
	70: 7,
	80: 6,
	90: 6,
}

var mods := map[int]int{
	
}

func main() {
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += letterCount(i)
	}
	fmt.Printf("sum is %d\n", sum)
}

func letterCount(i int) int {
	if letters[i] {
		return letters[i]
	}
	return 0
}