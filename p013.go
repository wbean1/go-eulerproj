package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	strs := parseFile("p013_input.txt")
	sum := new(big.Int)
	for _, str := range strs {
		bigN := new(big.Int)
		bigN, _ = bigN.SetString(str, 10)
		sum.Add(sum, bigN)
	}
	fmt.Printf("sum is: %s", sum.String())
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
		words = append(words, scanner.Text())
	}
	return words
}
