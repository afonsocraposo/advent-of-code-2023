package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d *Direction) rotateLeft() Direction {
	switch *d {
	case UP:
		return LEFT
	case LEFT:
		return DOWN
	case DOWN:
		return RIGHT
	}
	return UP
}
func (d *Direction) rotateRight() Direction {
	switch *d {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	}
	return UP
}

type Node struct {
	i          int
	j          int
	comingFrom Direction
	cost       int
}

func (n *Node) moveTo(puzzle [][]int, direction Direction) (Node, bool) {

	node := Node{}
	switch direction {
	case UP:
		node.i = n.i - 1
		if node.i < 0 {
			return node, false
		}
		node.j = n.j
		node.comingFrom = DOWN
	case RIGHT:
		node.j = n.j + 1
		if node.j >= len(puzzle[0]) {
			return node, false
		}
		node.i = n.i
		node.comingFrom = LEFT
	case DOWN:
		node.i = n.i + 1
		if node.i >= len(puzzle) {
			return node, false
		}
		node.j = n.j
		node.comingFrom = UP
	case LEFT:
		node.j = n.j - 1
		if node.j < 0 {
			return node, false
		}
		node.i = n.i
		node.comingFrom = RIGHT
	default:
		return node, false
	}

	node.cost = n.cost + puzzle[node.i][node.j]

	return node, true
}

func label(n Node) string {
    return fmt.Sprintf("%d:%d:%d", n.i, n.j, n.comingFrom)
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
			steps := 0
			node := current
			for steps < 3 {
				node, canMove = node.moveTo(puzzle, direction)
				if canMove {
					l := label(node)
					c, found := costs[l]
					if !found || node.cost < c {
						costs[l] = node.cost
                        queue = append(queue, node)
					}
				}
				steps++
			}
		}
	}

	return minCost
}
