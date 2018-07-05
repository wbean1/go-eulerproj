// https://projecteuler.net/problem=81
// notes:
//  2nd problem solved
//  completed in ~1hr
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
	matrix := parseFile()
	fmt.Print(matrix)
	// ok, so the plan is to start in the top left and compute the
	// minimal-path-sum (mps) from the top left for each position by looking only
	// at a 2x2 matrix at a time
	//given:
	// [ 20, 30 ]
	// [ 40, 50 ]
	//return:
	// [ 20, 50  ]
	// [ 60, 100 ]
	// should be able to do this recursively, but that might get tricky,
	// so probably will just iterate to be lazy
	// the answer will then be the value of the lower right element of the result
	result := make([][]int, 80)
	for i := range result {
		result[i] = make([]int, 80)
		for j := range result[i] {
			cursorValue, _ := strconv.Atoi(matrix[i][j])
			if (i == 0) && (j == 0) {
				result[i][j] = cursorValue
			} else if i == 0 {
				result[i][j] = cursorValue + result[i][j-1]
			} else if j == 0 {
				result[i][j] = cursorValue + result[i-1][j]
			} else {
				downSum := cursorValue + result[i-1][j]
				rightSum := cursorValue + result[i][j-1]
				result[i][j] = min(downSum, rightSum)
			}
		}
	}
	fmt.Print(result)
	fmt.Printf("answer is: %d", result[79][79])

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parseFile() [][]string {
	file, err := os.Open("p081_matrix.txt")
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
