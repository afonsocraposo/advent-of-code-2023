package part2

import (
	. "github.com/afonsocraposo/advent-of-code-2023/7/common"
	"github.com/afonsocraposo/advent-of-code-2023/7/part1"
)

var CARD_TO_SCORE = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
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
		h._handType = computeHandType(hist)
	}
	return h._handType
}

func computeHandType(hist map[rune]int) HandType {

	j, ok := hist['J']
	if !ok {
		return part1.ComputeHandType(hist)
	}

	m := 0
	for c, val := range hist {
		if c != 'J' && val > m {
			m = val
		}
	}

	switch m + j {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if len(hist) == 3 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	}
	return OnePair
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
