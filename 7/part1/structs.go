package part1

import (
	. "github.com/afonsocraposo/advent-of-code-2023/7/common"
)

var CARD_TO_SCORE = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type Hand struct {
	cards     string
	bid       int
	rank      int
	_handType HandType
}

func (h *Hand) handType() HandType {
	if h._handType == 0 {
		hist := Histogram(h.cards)
		h._handType = ComputeHandType(hist)
	}
	return h._handType
}

func ComputeHandType(hist map[rune]int) HandType {
	switch len(hist) {
	case 1:
		return FiveOfAKind
	case 2:
		for _, val := range hist {
			if val == 1 || val == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	default:
		pairs := 0
		for _, val := range hist {
			if val == 3 {
				return ThreeOfAKind
			} else if val == 2 {
				pairs++
			}
		}
		if pairs == 2 {
			return TwoPair
		} else if pairs == 1 {
			return OnePair
		}
	}
	return HighCard
}

func (h1 *Hand) compare(h2 *Hand) int {
	if h1.handType() < h2.handType() {
		return 1
	} else if h1.handType() > h2.handType() {
		return -1
	} else {
		for i := 0; i < 5; i++ {
			h1c := rune(h1.cards[i])
			h2c := rune(h2.cards[i])
			if h1c != h2c {
				if CARD_TO_SCORE[h1c] < CARD_TO_SCORE[h2c] {
					return 1
				} else {
					return -1
				}
			}
		}
	}
	return 0
}
