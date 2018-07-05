// https://projecteuler.net/problem=11
// notes:
//   not terribly elegant, but solves the problem
//   completed in <1hr

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	matrixStr := parseFile()

	// lets use integers instead of strings...
	m := make([][]int, 20)
	for y := range matrixStr {
		m[y] = make([]int, 20)
		for x := range matrixStr[y] {
			m[y][x], _ = strconv.Atoi(matrixStr[y][x])
		}
	}
	// fmt.Print(matrixInt)

	answer := 0
	for y := 0; y < 20; y++ {
		for x := 0; x < 20; x++ {

			//look SOUTH (y needs to be 16 or less)
			if y <= 16 {
				answer = max(answer, (m[y][x] * m[y+1][x] * m[y+2][x] * m[y+3][x]))
			}

			//look SE (y & x needs to be 16 or less)
			if (y <= 16) && (x <= 16) {
				answer = max(answer, (m[y][x] * m[y+1][x+1] * m[y+2][x+2] * m[y+3][x+3]))
			}

			//look EAST (x needs to be 16 or less)
			if x <= 16 {
				answer = max(answer, (m[y][x] * m[y][x+1] * m[y][x+2] * m[y][x+3]))
			}

			//look NE (x needs to be 16 or less, y needs to be 3 or more)
			if (y >= 3) && (x <= 16) {
				answer = max(answer, (m[y][x] * m[y-1][x+1] * m[y-2][x+2] * m[y-3][x+3]))
			}

			//look NORTH (y needs to be 3 or more)
			if y >= 3 {
				answer = max(answer, (m[y][x] * m[y-1][x] * m[y-2][x] * m[y-3][x]))
			}

			//look NW (y needs to be 3 or more, x needs to be 3 or more)
			if (y >= 3) && (x >= 3) {
				answer = max(answer, (m[y][x] * m[y-1][x-1] * m[y-2][x-2] * m[y-3][x-3]))
			}

			//look WEST
			if x >= 3 {
				answer = max(answer, (m[y][x] * m[y][x-1] * m[y][x-2] * m[y][x-3]))
			}

			//look SW
			if (x >= 3) && (y <= 16) {
				answer = max(answer, (m[y][x] * m[y+1][x-1] * m[y+2][x-2] * m[y+3][x-3]))
			}
		}
	}
	fmt.Printf("Answer is: %d", answer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseFile() [][]string {
	file, err := os.Open("p011_matrix.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}
