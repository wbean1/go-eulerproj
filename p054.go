// https://projecteuler.net/problem=54
// notes:
//  1st problem solved
//  took way too long... ~6hr?
//  basically first thing i've ever written in go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Card struct {
	value int
	suit  string
}

type Hand struct {
	cards []Card
}

func (h *Hand) isFlush() bool {
	if (h.cards[0].suit == h.cards[1].suit) && (h.cards[1].suit == h.cards[2].suit) && (h.cards[2].suit == h.cards[3].suit) && (h.cards[3].suit == h.cards[4].suit) {
		return true
	} else {
		return false
	}
}

func (h *Hand) isStraight() bool {
	if (h.cards[0].value == (h.cards[1].value - 1)) && (h.cards[1].value == (h.cards[2].value - 1)) && (h.cards[2].value == (h.cards[3].value - 1)) && (h.cards[3].value == (h.cards[4].value - 1)) {
		return true
	} else {
		return false
	}
}

func howMany(c [5]int, lookingfor int) int {
	res := 0
	for i := 0; i < 5; i++ {
		if c[i] == lookingfor {
			res++
		}
	}
	return res
}

func (h *Hand) isFourKind() (bool, int) {
	c := [5]int{
		h.cards[0].value,
		h.cards[1].value,
		h.cards[2].value,
		h.cards[3].value,
		h.cards[4].value,
	}
	for i := 2; i < 15; i++ {
		if howMany(c, i) == 4 {
			return true, i
		}
	}
	return false, 0
}

func (h *Hand) isThreeKind() (bool, int) {
	c := [5]int{
		h.cards[0].value,
		h.cards[1].value,
		h.cards[2].value,
		h.cards[3].value,
		h.cards[4].value,
	}
	for i := 2; i < 15; i++ {
		if howMany(c, i) == 3 {
			return true, i
		}
	}
	return false, 0
}

func (h *Hand) isTwoPair() (bool, int) {
	c := [5]int{
		h.cards[0].value,
		h.cards[1].value,
		h.cards[2].value,
		h.cards[3].value,
		h.cards[4].value,
	}
	numPairFound := 0
	highestPair := 0
	for i := 2; i < 15; i++ {
		if howMany(c, i) == 2 {
			numPairFound++
			if i > highestPair {
				highestPair = i
			}

		}
	}
	if numPairFound == 2 {
		return true, highestPair
	}
	return false, 0
}

func (h *Hand) isPair() (bool, int) {
	c := [5]int{
		h.cards[0].value,
		h.cards[1].value,
		h.cards[2].value,
		h.cards[3].value,
		h.cards[4].value,
	}
	for i := 2; i < 15; i++ {
		if howMany(c, i) == 2 {
			return true, i
		}
	}
	return false, 0
}

func importHands() []Hand {
	dat, err := ioutil.ReadFile("p054_poker.txt")
	if err != nil {
		log.Fatal(err)
	}
	hands := make([]Hand, 0)
	for i := 0; i < len(dat); i = i + 30 {
		h1 := parseHand(dat[i : i+14])
		h2 := parseHand(dat[i+15 : i+30])
		hands = append(hands, h1, h2)
	}
	return hands
}

func parseHand(h []byte) Hand {
	c := make([]Card, 0)
	c = append(c, parseCard(h[0:2]))
	c = append(c, parseCard(h[3:5]))
	c = append(c, parseCard(h[6:8]))
	c = append(c, parseCard(h[9:11]))
	c = append(c, parseCard(h[12:14]))
	sort.Slice(c, func(i, j int) bool {
		return c[i].value < c[j].value
	})
	hand := Hand{c}
	return hand
}

func parseCard(c []byte) Card {
	card := Card{}
	switch string(c[0]) {
	case "2":
		card.value = 2
	case "3":
		card.value = 3
	case "4":
		card.value = 4
	case "5":
		card.value = 5
	case "6":
		card.value = 6
	case "7":
		card.value = 7
	case "8":
		card.value = 8
	case "9":
		card.value = 9
	case "T":
		card.value = 10
	case "J":
		card.value = 11
	case "Q":
		card.value = 12
	case "K":
		card.value = 13
	case "A":
		card.value = 14
	default:
		panic("unknown card.value")
	}
	switch string(c[1]) {
	case "S":
		card.suit = "S"
	case "C":
		card.suit = "C"
	case "H":
		card.suit = "H"
	case "D":
		card.suit = "D"
	default:
		panic("unknown card.suit")
	}
	return card
}

func whoWins(hand1, hand2 Hand) string {
	s1 := scoreHand(hand1)
	s2 := scoreHand(hand2)
	fmt.Printf("player1 has %+v for score: %f\n", hand1, s1)
	fmt.Printf("player2 has %+v for score: %f\n\n", hand2, s2)
	if s1 > s2 {
		return "one"
	} else {
		return "two"
	}
}

func scoreHand(h Hand) float64 {
	// scores for hands:
	// * highcard: 2 - 14.130
	// * one pair: (10*pair) = 20 - 140
	// * two pair: (1000*pair) = 2000 - 14000
	// * three kind: (10,000*pair) = 20000 - 140000
	// x straight: (100,000*high) = 500000 - 1400000
	// x flush: (1,000,000*high) = 5000000 - 14000000
	// * full-house (10,000,000*pair) = 20,000,000 - 140,000,000
	// * four kind: (100,000,000*pair) = 200,000,000 - 1,400,000,000
	// x straight + flush: (1,000,000,000*high = 5,000,000,000 - 14,000,000,000
	var score float64
	isFourKind, foursRank := h.isFourKind()
	isThreeKind, threesRank := h.isThreeKind()
	isPair, pairRank := h.isPair()
	isTwoPair, twoPairRank := h.isTwoPair()
	if h.isFlush() && h.isStraight() {
		score += (float64(h.cards[4].value) * 1000000000)
		score += (float64(h.cards[3].value) * 0.01)
		score += (float64(h.cards[2].value) * 0.0001)
		score += (float64(h.cards[1].value) * 0.000001)
		score += (float64(h.cards[0].value) * 0.00000001)
	}
	if h.isFlush() {
		score += (float64(h.cards[4].value) * 1000000)
		score += (float64(h.cards[3].value) * 0.01)
		score += (float64(h.cards[2].value) * 0.0001)
		score += (float64(h.cards[1].value) * 0.000001)
		score += (float64(h.cards[0].value) * 0.00000001)
	} else if h.isStraight() {
		score += (float64(h.cards[4].value) * 100000)
		score += (float64(h.cards[3].value) * 0.01)
		score += (float64(h.cards[2].value) * 0.0001)
		score += (float64(h.cards[1].value) * 0.000001)
		score += (float64(h.cards[0].value) * 0.00000001)
	} else if isFourKind {
		score += (float64(foursRank) * 100000000)
	} else if isThreeKind && isPair {
		// Fuller House
		score += (float64(threesRank) * 10000000)
	} else if isThreeKind {
		score += (float64(threesRank) * 10000)
	} else if isTwoPair {
		score += (float64(twoPairRank) * 1000)
	} else if isPair {
		score += (float64(pairRank) * 10)
		score += (float64(h.cards[4].value) * 0.01)
		score += (float64(h.cards[3].value) * 0.0001)
		score += (float64(h.cards[2].value) * 0.000001)
		score += (float64(h.cards[1].value) * 0.00000001)
		score += (float64(h.cards[0].value) * 0.0000000001)
	} else {
		score += (float64(h.cards[4].value) * 1)
		score += (float64(h.cards[3].value) * 0.01)
		score += (float64(h.cards[2].value) * 0.0001)
		score += (float64(h.cards[1].value) * 0.000001)
		score += (float64(h.cards[0].value) * 0.00000001)
	}
	return score
}

func main() {
	hands := importHands()
	one := 0
	for i := 0; i < 2000; i = i + 2 {
		wins := whoWins(hands[i], hands[i+1])
		if wins == "one" {
			one++
		}
	}
	fmt.Println("player one won this many hands: ", one)
}
