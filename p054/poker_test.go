package main

import (
	"math"
	"testing"
)

type pokerHandCompareCase struct {
	description string
	winnerHand  string
	loserHand   string
}

var compareTestCases = []pokerHandCompareCase{
	{
		description: "A high beats K high",
		winnerHand:  "AH 2D 3H 4C 9D",
		loserHand:   "KH 3D 4H 5C TD",
	},
	{
		description: "Kickers matter",
		winnerHand:  "AH 3D 4C 5C AS",
		loserHand:   "AD 2H 4H 5D AC",
	},
	{
		description: "Full houses, higher overs",
		winnerHand:  "4H 4D 4C 2C 2S",
		loserHand:   "3D 3H 3C 9D 9C",
	},
	{
		description: "Pair beats high card",
		winnerHand:  "4H 4D AC 2C 3S",
		loserHand:   "AD KH QC TD 9C",
	},
	{
		description: "Three kind beats pair",
		winnerHand:  "4H 4D 4C 2C AS",
		loserHand:   "4H 4D AC 2C 3S",
	},
	{
		description: "Straight beats Three kind",
		winnerHand:  "4H 5D 3C 2C AS",
		loserHand:   "4H 4D 4C 2C AS",
	},
	{
		description: "Flush beats Straight",
		winnerHand:  "KD 9D 4D 2D 8D",
		loserHand:   "4H 5D 3C 2C AS",
	},
	{
		description: "FullHouse beans Flush",
		winnerHand:  "4H 4D 4C 2C 2S",
		loserHand:   "KD 9D 4D 2D 8D",
	},
	{
		description: "FourKind beats FullHouse",
		winnerHand:  "4H 4D 4C 4C 2S",
		loserHand:   "4H 4D 4C 2C 2S",
	},
	{
		description: "StraightFlush beats FourKind",
		winnerHand:  "4H 5H 6H 7H 8H",
		loserHand:   "4H 4D 4C 4C 2S",
	},
	{
		description: "Chicken Dinner",
		winnerHand:  "TH JH QH KH AH",
		loserHand:   "4H 5H 6H 7H 8H",
	},
	{
		description: "Compare straight flushes",
		winnerHand:  "4H 5H 6H 7H 8H",
		loserHand:   "5H 4H 3H 2H AH",
	},
}

func TestCompareScoreHand(t *testing.T) {
	for _, testCase := range compareTestCases {
		winner := parseHand([]byte(testCase.winnerHand))
		loser := parseHand([]byte(testCase.loserHand))
		winnerScore := ScoreHand(winner)
		loserScore := ScoreHand(loser)
		if !(winnerScore > loserScore) {
			t.Fatalf("FAIL: %s(winnerHand: %s, loserHand: %s)",
				testCase.description, testCase.winnerHand, testCase.loserHand)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

type pokerTestCase struct {
	description string
	input       string
	isFlush     bool
	isStraight  bool
	isFourKind  bool
	isThreeKind bool
	isPair      bool
	value       float64
}

var testCases = []pokerTestCase{
	{
		description: "K-high flush (not straight)",
		input:       "KD TD 9D 8D 7D",
		isFlush:     true,
		isStraight:  false,
		isFourKind:  false,
		isThreeKind: false,
		isPair:      false,
		value:       13000000.10090807,
	},
	{
		description: "K-high straight (not flush)",
		input:       "KD QD JD TD 9H",
		isFlush:     false,
		isStraight:  true,
		isFourKind:  false,
		isThreeKind: false,
		isPair:      false,
		value:       1300000,
	},
	{
		description: "Four Ks",
		input:       "KH KD KC KS 2D",
		isFlush:     false,
		isStraight:  false,
		isFourKind:  true,
		isThreeKind: false,
		isPair:      false,
		value:       1300000000,
	},
	{
		description: "Full House, kings over twos",
		input:       "KH KD KC 2S 2D",
		isFlush:     false,
		isStraight:  false,
		isFourKind:  false,
		isThreeKind: true,
		isPair:      true,
		value:       130000000,
	},
	{
		description: "Three Ks",
		input:       "KH KD KC 3S 2D",
		isFlush:     false,
		isStraight:  false,
		isFourKind:  false,
		isThreeKind: true,
		isPair:      false,
		value:       130000,
	},
	{
		description: "Two Ks",
		input:       "KH KD 4C 3S 2D",
		isFlush:     false,
		isStraight:  false,
		isFourKind:  false,
		isThreeKind: false,
		isPair:      true,
		value:       130.1313040302,
	},
	{
		description: "King high",
		input:       "KH 5D 4C 3S 2D",
		isFlush:     false,
		isStraight:  false,
		isFourKind:  false,
		isThreeKind: false,
		isPair:      false,
		value:       13.05040302,
	},
	{
		description: "Five high straight",
		input:       "AH 2D 3C 4S 5D",
		isFlush:     false,
		isStraight:  true,
		isFourKind:  false,
		isThreeKind: false,
		isPair:      false,
		value:       500000,
	},
}

func TestIsFlush(t *testing.T) {
	for _, testCase := range testCases {
		h := parseHand([]byte(testCase.input))
		if is, _ := h.isFlush(); is != testCase.isFlush {
			t.Fatalf("FAIL: %s(%s)\nExpected: %t\nActual: %t",
				testCase.description, testCase.input, testCase.isFlush, is)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func TestIsStraight(t *testing.T) {
	for _, testCase := range testCases {
		h := parseHand([]byte(testCase.input))
		if is, _ := h.isStraight(); is != testCase.isStraight {
			t.Fatalf("FAIL: %s(%s)\nExpected: %t\nActual: %t",
				testCase.description, testCase.input, testCase.isStraight, is)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func TestIsFourKind(t *testing.T) {
	for _, testCase := range testCases {
		h := parseHand([]byte(testCase.input))
		if is, _ := h.isFourKind(); is != testCase.isFourKind {
			t.Fatalf("FAIL: %s(%s)\nExpected: %t\nActual: %t",
				testCase.description, testCase.input, testCase.isFourKind, is)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func TestIsThreeKind(t *testing.T) {
	for _, testCase := range testCases {
		h := parseHand([]byte(testCase.input))
		if is, _ := h.isThreeKind(); is != testCase.isThreeKind {
			t.Fatalf("FAIL: %s(%s)\nExpected: %t\nActual: %t",
				testCase.description, testCase.input, testCase.isThreeKind, is)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func TestIsPairKind(t *testing.T) {
	for _, testCase := range testCases {
		h := parseHand([]byte(testCase.input))
		if is, _ := h.isPair(); is != testCase.isPair {
			t.Fatalf("FAIL: %s(%s)\nExpected: %t\nActual: %t",
				testCase.description, testCase.input, testCase.isPair, is)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func TestScoreHand(t *testing.T) {
	for _, testCase := range testCases {
		h := parseHand([]byte(testCase.input))
		if score := ScoreHand(h); !floatEquals(score, testCase.value) {
			t.Fatalf("FAIL: %s(%s)\nExpected: %.8f\nActual: %.8f",
				testCase.description, testCase.input, testCase.value, score)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

var eps float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}
