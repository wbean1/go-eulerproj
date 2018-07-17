package main

import (
	"io/ioutil"
	"log"
	"sort"
)

// ImportHands takes a string filename, opens and parses it
// and returns an array of contained Hands
func ImportHands(file string) []Hand {
	dat, err := ioutil.ReadFile(file)
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

// parseHand is a helper that takes a byte array
// representing 5 Cards, parses them and returns
// them as a Hand
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

// parseCard is a helper that takes a byte array
// representing 1 Card, parses it to make sure
// it matches expected values, and returns
// it as a Card
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
