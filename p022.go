package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := Input()
	sum := 0
	for i, name := range input {
		sum += scoreWord(name) * (i + 1)
	}
	fmt.Print(sum)
}

func Input() []string {
	return parseFile("/Users/wbean/go-eulerproj/p022_names_fmt.txt")
}

func scoreWord(s string) int {
	score := 0
	for _, letter := range s {
		score += int(letter) - 64
	}
	return score
}

func parseFile(f string) []string {

	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
