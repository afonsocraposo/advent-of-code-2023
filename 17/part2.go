package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part2(filepath string) int {
	fmt.Println("Part 2")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	puzzle := [][]int{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for j, c := range line {
			cost, _ := strconv.Atoi(string(c))
			row[j] = cost
		}
		puzzle = append(puzzle, row)
		i++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	m := len(puzzle)
	n := len(puzzle[0])

	costs := map[string]int{}
	queue := []Node{
		{0, 0, UP, 0},
		{0, 0, LEFT, 0},
	}
	var current Node
	var canMove bool
	minCost := 99999999
	for len(queue) > 0 {
		// pop
		current = queue[0]
		queue = queue[1:]
		if current.i == m-1 && current.j == n-1 {
			minCost = min(current.cost, minCost)
		}

		for _, direction := range []Direction{
			current.comingFrom.rotateLeft(),
			current.comingFrom.rotateRight(),
		} {
			steps := 1
			node := current
			for steps <= 10 {
				node, canMove = node.moveTo(puzzle, direction)
				if canMove {
					l := label(node)
					c, found := costs[l]
					if !found || node.cost < c {
						if steps >= 4 {
							costs[l] = node.cost
							queue = append(queue, node)
						}
					}
				}
				steps++
			}
		}
	}

	return minCost
}
