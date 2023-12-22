package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkAlmostSymmetric(a [][]int, b [][]int) bool {
	n := len(a[0])
	m := min(len(a), len(b))
	c := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				if !c {
					c = true
				} else {
					return false
				}
			}
		}
	}
	return c
}

func findSingleReflectionLine(pattern [][]int) int {
	result := 0
	for i := range pattern[:len(pattern)-1] {
		top := flipHorizontal(pattern[:i+1])
		bottom := pattern[i+1:]
		if checkSymmetry(top, bottom) {
			if result == 0 {
				result = i + 1
			} else {
				return 0
			}
		}
	}
	return result
}

func findFixedReflectionLine(pattern [][]int) int {
	for i := range pattern[:len(pattern)-1] {
		top := pattern[:i+1]
		topFlip := flipHorizontal(top)
		bottom := pattern[i+1:]
		if checkAlmostSymmetric(topFlip, bottom) {
            return i+1
		}
	}
	return 0
}

func part2(filepath string) int {
	fmt.Println("Part 2")

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
			horizontal := findFixedReflectionLine(pattern)
			vertical := findFixedReflectionLine(transpose(pattern))
			result += vertical + 100*horizontal
			pattern = [][]int{}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
