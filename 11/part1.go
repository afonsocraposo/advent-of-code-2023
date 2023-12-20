package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func printUniverse(universe []string) {
	for _, line := range universe {
		fmt.Println(line)
	}
}

func expandUniverse(universe []string, rowsExpand []int, colsExpand []int) []string {
    offset := 0
	for _, r := range rowsExpand {
        row := r+offset
		universe = append(universe[:row+1], universe[row:]...)
        offset++
	}
	for i, line := range universe {
        offset := 0
		for _, c := range colsExpand {
            col := c + offset
			line = line[:col+1] + line[col:]
            offset++
		}
		universe[i] = line
	}
	return universe
}

func part1(filepath string) int {
	fmt.Println("Part 1")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	universe := make([]string, 0)
	rowsExpand := []int{}
	colsExpand := []int{}
	excludeCols := map[int]bool{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		universe = append(universe, line)

		excludeRow := false
		for j, char := range line {
			if char == '#' {
				excludeRow = true
				excludeCols[j] = true
			}
		}
		if !excludeRow {
			rowsExpand = append(rowsExpand, i)
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for j := range universe[0] {
		_, found := excludeCols[j]
		if !found {
			colsExpand = append(colsExpand, j)
		}
	}

	universe = expandUniverse(universe, rowsExpand, colsExpand)

	galaxies := [][]int{}
	for i, line := range universe {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, []int{i, j})
			}
		}
	}

    result := 0
    for i, g1 := range galaxies[:len(galaxies)-1] {
		for _, g2 := range galaxies[i+1:] {
            d := math.Abs(float64(g1[0] - g2[0])) + math.Abs(float64(g1[1] - g2[1]))
            result += int(d)
		}
	}

	return result
}
