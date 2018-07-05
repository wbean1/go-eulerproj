// https://projecteuler.net/problem=82
// notes:
//	 took about 1.5hr after p081...
//   screwed up orientation of csv.ReadAll vs result matrixes
//   looping for corrections is a terrible hack
//	 ... i'm not very proud of this, but it does give the right answer.

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
	// different boundary conditions from p081, but similar approach

	//    0     x      79
	//   |--------------|
	// 0-
	//  |
	//  |
	// y|
	//  |
	//  |
	//79-

	// initialize the result matrix...
	result := make([][]int, 80)
	for y := range result {
		result[y] = make([]int, 80)
	}
	resultArr := make([]int, 80) // this will hold the last column that is all we care about...

	// i'm going to first iterate down columns from left to right...
	// in this loop i can only calculate the right & down moves, will have
	// to loop other direction and correct for up moves?
	for x := range result {
		for y := range result[x] {
			cursorValue, _ := strconv.Atoi(matrix[y][x])
			if x == 0 {
				result[y][x] = cursorValue
			} else if y == 0 {
				// assume moved right; cant compute moved up yet...
				rightSum := cursorValue + result[y][x-1]
				result[y][x] = rightSum
			} else {
				// assume moved right or down; can't compute moved up yet...
				rightSum := cursorValue + result[y][x-1]
				downSum := cursorValue + result[y-1][x]
				result[y][x] = min(rightSum, downSum)
			}
		}
	}

	// I have to run this 80 times... since corrections influence other corrections...
	// this is awful, but gives me the correct answer...
	for i := 0; i < 80; i++ {
		// up move correction loops!
		for x := range result {
			for y := 79; y >= 0; y-- {
				cursorValue, _ := strconv.Atoi(matrix[y][x])
				if x == 0 {
					result[y][x] = cursorValue
				} else if y == 79 {
					rightSum := cursorValue + result[y][x-1]
					downSum := cursorValue + result[y-1][x]
					result[y][x] = min(rightSum, downSum)
				} else if y == 0 {
					rightSum := cursorValue + result[y][x-1]
					upSum := cursorValue + result[y+1][x]
					result[y][x] = min(rightSum, upSum)
				} else {
					rightSum := cursorValue + result[y][x-1]
					upSum := cursorValue + result[y+1][x]
					downSum := cursorValue + result[y-1][x]
					result[y][x] = min(rightSum, min(upSum, downSum))
				}
				if x == 79 {
					resultArr[y] = result[y][x]
				}
			}
		}
	}

	fmt.Print(result)
	fmt.Print(resultArr)
	finish := minIntSlice(resultArr)
	fmt.Printf("the answer is: %d", finish)
}

func minIntSlice(v []int) int {
	m := v[0] // lol, terrible
	for e := range v {
		if v[e] < m {
			m = v[e]
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parseFile() [][]string {
	file, err := os.Open("p082_matrix.txt")
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
