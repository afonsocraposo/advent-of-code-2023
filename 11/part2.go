package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const EXPAND = 999999

func part2(filepath string) int {
	fmt.Println("Part 2")

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
	galaxies := [][]int{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		universe = append(universe, line)

		excludeRow := false
		for j, char := range line {
			if char == '#' {
				excludeRow = true
				excludeCols[j] = true
                galaxies = append(galaxies, []int{i, j})
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

	for index, galaxy := range galaxies {
        times := 0
        for _, i := range rowsExpand {
            if galaxy[0] > i {
                times++
            } else {
                break
            }
        }
        galaxy[0] += times*EXPAND
        times = 0
        for _, j := range colsExpand {
            if galaxy[1] > j {
                times++
            } else {
                break
            }
        }
        galaxy[1] += times*EXPAND
        galaxies[index] = galaxy
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
