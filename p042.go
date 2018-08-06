package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var letterVal = map[rune]int{
	'A': 1, 'B': 2, 'C': 3, 'D': 4,
	'E': 5, 'F': 6, 'G': 7, 'H': 8,
	'I': 9, 'J': 10, 'K': 11, 'L': 12,
	'M': 13, 'N': 14, 'O': 15, 'P': 16,
	'Q': 17, 'R': 18, 'S': 19, 'T': 20,
	'U': 21, 'V': 22, 'W': 23, 'X': 24,
	'Y': 25, 'Z': 26,
}

func main() {
	var count int
	words := parseFile("p042_words.txt")
	for _, word := range words {
		if isTriangleWord(word) {
			count++
		}
	}
	fmt.Printf("answer is: %d", count)
}

func parseFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	words := []string{}
	for scanner.Scan() {
		words = append(words, strings.Split(strings.Replace(scanner.Text(), "\"", "", -1), ",")...)
	}
	return words
}

func isTriangleWord(s string) bool {
	var wordValue int
	for _, char := range s {
		wordValue += letterVal[char]
	}
	wordValue = (8 * wordValue) + 1
	return int(math.Sqrt(float64(wordValue)))*int(math.Sqrt(float64(wordValue))) == wordValue
}
