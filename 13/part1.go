package main

import (
	"bufio"
	"fmt"
	"os"
)

func getRow(line string) []int {
	row := make([]int, len(line))
	for i, c := range line {
		if c == '#' {
			row[i] = 1
		}
	}
	return row
}

func flipHorizontal(pattern [][]int) [][]int {
	n := len(pattern)
	result := make([][]int, n)
	for i, row := range pattern {
		result[n-1-i] = row
	}
	return result
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
        for _, col := range row {
            if col == 0 {
                fmt.Print(".")
            } else {
                fmt.Print("#")
            }
        }
        fmt.Print("\n")
	}
}

func checkSymmetry(a [][]int, b [][]int) bool {
	n := len(a[0])
	m := min(len(a), len(b))
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func findReflectionLine(pattern [][]int) int {
	for i := range pattern[:len(pattern)-1] {
		top := flipHorizontal(pattern[:i+1])
		bottom := pattern[i+1:]
		if checkSymmetry(top, bottom) {
			return i + 1
		}
	}
	return 0
}

func transpose(pattern [][]int) [][]int {
	m := len(pattern)
	n := len(pattern[0])
	result := make([][]int, n)
	for j := range pattern[0] {
		result[j] = make([]int, m)
		for i := range pattern {
			result[j][i] = pattern[i][j]
		}
	}
	return result
}

func part1(filepath string) int {
	fmt.Println("Part 1")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	pattern := [][]int{}
    result := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			pattern = append(pattern, getRow(line))
		} else {
			horizontal := findReflectionLine(pattern)
			vertical := findReflectionLine(transpose(pattern))
            pattern = [][]int{}
            result += vertical + 100*horizontal
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
