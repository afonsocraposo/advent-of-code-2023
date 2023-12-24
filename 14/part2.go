package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	EMPTY = iota
	BLOCK
	ROCK
)

const CYCLES = 1000000000

var INT_TO_STRING = map[int]string{
	EMPTY: ".",
	BLOCK: "#",
	ROCK:  "O",
}

var CHAR_TO_INT = map[rune]int{
	'.': EMPTY,
	'#': BLOCK,
	'O': ROCK,
}

func printPuzzle(puzzle [][]int) {
	for _, row := range puzzle {
		for _, val := range row {
			fmt.Print(INT_TO_STRING[val])
		}
		fmt.Println()
	}
	fmt.Println()
}

func rotate(puzzle [][]int) [][]int {
	m := len(puzzle)
	n := len(puzzle[0])
	result := make([][]int, n)
	for i := range result {
		row := make([]int, m)
		for j := range row {
			row[j] = puzzle[m-j-1][i]
		}
		result[i] = row
	}
	return result
}

func cycle(puzzle [][]int) [][]int {
	// NORTH
	p := tilt(puzzle)

	// WEST
	p = rotate(p)
	p = tilt(p)

	// SOUTH
	p = rotate(p)
	p = tilt(p)

	// EAST
	p = rotate(p)
	p = tilt(p)

	p = rotate(p)
	return p
}

func encode(puzzle [][]int) string {
	result := ""
	for _, row := range puzzle {
		for _, val := range row {
			result += INT_TO_STRING[val]
		}
		result += ":"
	}
	return result[:len(result)-1]
}

func decode(encoded string) [][]int {
	rows := strings.Split(encoded, ":")
	result := make([][]int, len(rows))
	for i, row := range rows {
		r := make([]int, len(row))
		for j, val := range row {
			r[j] = CHAR_TO_INT[val]
		}
		result[i] = r
	}
	return result
}

func average(m map[string]int) float32 {
	n := float32(len(m))
	result := float32(0)
	for _, val := range m {
		result += float32(val) / n
	}
	return result
}

func part2(filepath string) int {
	fmt.Println("Part 2")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	puzzle := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, c := range line {
			row[i] = CHAR_TO_INT[c]
		}
		puzzle = append(puzzle, row)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	cache := []string{}
	offset := 0
	for i := 1; i <= CYCLES; i++ {
		puzzle = cycle(puzzle)
		e := encode(puzzle)
		if offset = slices.Index(cache, e); offset == -1 {
		    cache = append(cache, e)
		}else {
		    break
		}
	}
	period := len(cache) - offset

	c := (CYCLES - offset) % period
	for i := 0; i < c-1; i++ {
		puzzle = cycle(puzzle)
	}

	result := load(puzzle)

	return result
}
