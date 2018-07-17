package main

import (
	"fmt"
)

func main() {
	hands := ImportHands("p054_poker.txt")
	one := 0
	for i := 0; i < 2000; i = i + 2 {
		wins := whoWins(hands[i], hands[i+1])
		if wins == "one" {
			one++
		}
	}
	fmt.Println("player one won this many hands: ", one)
}

func whoWins(hand1, hand2 Hand) string {
	s1 := ScoreHand(hand1)
	s2 := ScoreHand(hand2)
	fmt.Printf("player1 has %+v for score: %f\n", hand1, s1)
	fmt.Printf("player2 has %+v for score: %f\n\n", hand2, s2)
	if s1 > s2 {
		return "one"
	} else {
		return "two"
	}
}
