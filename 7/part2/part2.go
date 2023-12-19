package part2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rankHands(hands *[]Hand) {
	for i := 0; i < len(*hands); i++ {
		h1 := &(*hands)[i]
		for j := i; j < len(*hands); j++ {
			h2 := &(*hands)[j]
			if h1.compare(h2) > 0 {
				h2.rank++
			} else {
				h1.rank++
			}
		}
	}
}

func Part2(filepath string) int {
	fmt.Println("Part 2")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	hands := make([]Hand, 0)
	for scanner.Scan() {
		line := scanner.Text()
		cards, bidStr, _ := strings.Cut(line, " ")
		bid, _ := strconv.Atoi(bidStr)

		h := Hand{}
		h.cards = cards
		h.bid = bid
		hands = append(hands, h)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rankHands(&hands)

	result := 0
	for _, hand := range hands {
		result += hand.rank * hand.bid
	}

	return result
}
