package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := parseFile("p018_triangle.txt")
	for i := len(matrix) - 2; i >= 0; i-- {
		for j := 0; j < len(matrix[i]); j++ {
			left := matrix[i+1][j]
			right := matrix[i+1][j+1]
			if left > right {
				matrix[i][j] += left
			} else {
				matrix[i][j] += right
			}
		}
	}
	ans := matrix[0][0]
	fmt.Printf("Answer is: %d", ans)
}

func parseFile(f string) [][]int {

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)

	result := make([][]int, 15)
	for i := 0; i < 15; i++ {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		str := string(line[:])
		str = strings.TrimSpace(str)
		strs := strings.Split(str, " ")
		result[i] = intize(strs)
	}
	return result
}

func intize(s []string) []int {
	i := make([]int, len(s))
	for j, k := range s {
		myInt, _ := strconv.Atoi(k)
		i[j] = myInt
	}
	return i
}
