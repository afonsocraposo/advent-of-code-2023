package main

import (
	"bufio"
	"fmt"
	"os"
)

func load(puzzle [][]int) int {
	result := 0
	for j := 0; j < len(puzzle[0]); j++ {
		for i := 0; i < len(puzzle); i++ {
			c := puzzle[i][j]
			if c == ROCK {
				result += len(puzzle)-i
            }
		}
	}
	return result
}

func tilt(puzzle [][]int) [][]int {
	for j := range puzzle[0] {
		rocks := 0
		block := -1
		for i := range puzzle {
			switch puzzle[i][j] {
			case ROCK:
				puzzle[i][j] = EMPTY
				rocks++
			case BLOCK:
				for rocks > 0 {
					puzzle[block+rocks][j] = ROCK
					rocks--
				}
				block = i
			}
		}
		for rocks > 0 {
			puzzle[block+rocks][j] = ROCK
			rocks--
		}
	}
	return puzzle
}

func part1(filepath string) int {
	fmt.Println("Part 1")

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

	result := load(tilt(puzzle))
	return result
}
