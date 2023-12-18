package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func part2(filepath string) int {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := regexp.MustCompile("[0-9]+")

	result := 0
	cards := make(map[int]int)
	toCopy := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		winningStr, numbersStr, _ := strings.Cut(line, "|")

		winning := r.FindAllString(winningStr, -1)
		card, _ := strconv.Atoi(winning[0])
		winning = winning[1:]

		numbers := r.FindAllString(numbersStr, -1)

		winningNrs := 0
		for _, n := range numbers {
			if slices.Contains(winning, n) {
				winningNrs++
				toCopy = append(toCopy, card+winningNrs)
			}
		}
		cards[card] = winningNrs
		result++
	}

	i := 0
	for i < len(toCopy) {
		card := toCopy[i]
		matches := cards[card]
		for c := card + 1; c < card+1+matches; c++ {
			toCopy = append(toCopy, c)
		}
		i++
	}
	result += i

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
