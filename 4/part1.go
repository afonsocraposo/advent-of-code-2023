package main

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strings"
)

func part1(filepath string) int {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

    r := regexp.MustCompile("[0-9]+")
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
        winningStr, numbersStr, _ := strings.Cut(line, "|")
        winning := r.FindAllString(winningStr, -1)[1:]
        numbers := r.FindAllString(numbersStr, -1)
        winningNrs := -1
        for _, n := range(numbers) {
            if slices.Contains(winning, n) {
                winningNrs++
            }
        }
        if winningNrs >= 0 {
            result += 1 << winningNrs
        }
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
