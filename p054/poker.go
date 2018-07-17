package main

type Card struct {
	value int
	suit  string
}

type Hand struct {
	cards []Card
}

// ScoreHand provides a numerical score that can be used to
// compare/rank a Hand value.  The score takes into account
// things that factor into a Hand's rank, like if it contains
// a pair, the additional cards also matter because another
// hand could also contain the same pair (and high card, etc...)

// This function highly relies on the Hand being sorted in value order

// scores for hands:
// * highcard: 2 - 14.130
// * one pair: (10*pair) = 20 - 140
// * two pair: (1000*pair) = 2000 - 14000
// * three kind: (10,000*pair) = 20000 - 140000
// * straight: (100,000*high) = 500000 - 1400000
// * flush: (1,000,000*high) = 5000000 - 14000000
// * full-house (10,000,000*pair) = 20,000,000 - 140,000,000
// * four kind: (100,000,000*pair) = 200,000,000 - 1,400,000,000
// * straight + flush: (1,000,000,000*high = 5,000,000,000 - 14,000,000,000
func ScoreHand(h Hand) float64 {
	var score float64
	isFlush, flushRank := h.isFlush()
	isStraight, straightRank := h.isStraight()
	isFourKind, foursRank := h.isFourKind()
	isThreeKind, threesRank := h.isThreeKind()
	isPair, pairRank := h.isPair()
	isTwoPair, twoPairRank := h.isTwoPair()
	if isFlush && isStraight {
		score += (float64(straightRank) * 1000000000)
	} else if isFlush {
		score += (float64(flushRank) * 1000000)
		score += (float64(h.cards[3].value) * 0.01)
		score += (float64(h.cards[2].value) * 0.0001)
		score += (float64(h.cards[1].value) * 0.000001)
		score += (float64(h.cards[0].value) * 0.00000001)
	} else if isStraight {
		score += (float64(straightRank) * 100000)
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

func (h *Hand) isFlush() (bool, int) {
	if (h.cards[0].suit == h.cards[1].suit) && (h.cards[1].suit == h.cards[2].suit) && (h.cards[2].suit == h.cards[3].suit) && (h.cards[3].suit == h.cards[4].suit) {
		return true, h.cards[4].value
	} else {
		return false, 0
	}
}

func (h *Hand) isStraight() (bool, int) {
	if (h.cards[0].value == (h.cards[1].value - 1)) && (h.cards[1].value == (h.cards[2].value - 1)) && (h.cards[2].value == (h.cards[3].value - 1)) && (h.cards[3].value == (h.cards[4].value - 1)) {
		return true, h.cards[4].value
	} else if (h.cards[4].value == 14) && (h.cards[0].value == 2) && (h.cards[1].value == 3) && (h.cards[2].value == 4) && (h.cards[3].value == 5) {
		return true, h.cards[3].value
	} else {
		return false, 0
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
