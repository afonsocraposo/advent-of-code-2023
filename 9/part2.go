package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2(filepath string) int {
	fmt.Println("Part 2")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbersStr := strings.Split(line, " ")
		numbers := make([]int, len(numbersStr))
		for i, ns := range numbersStr {
			number, err := strconv.Atoi(ns)
			if err != nil {
				panic(err)
			}
			numbers[i] = number
		}
		extrapolation := numbers[0]
		factor := -1
		for !allZeros(numbers) {
			numbers = derivative(numbers)
			extrapolation += factor * numbers[0]
			factor *= -1
		}
		result += extrapolation
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
