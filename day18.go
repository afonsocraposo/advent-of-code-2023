package main

import (
	"fmt"
	"strings"
)

func floodFill(puzzle Matrix, start Point) int {
	visited := map[Point]bool{start: true}
	queue := []Point{start}
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		for _, direction := range DIRECTIONS {
			point := start.move(direction)
			if puzzle.inside(point) {
				_, found := visited[point]
				if puzzle.get(point) == 0 && !found {
					visited[point] = true
					queue = append(queue, point)
				}
			}
		}
	}
	return len(visited)
}

func day18(filepath string) {
	text := readTextFile(filepath)

	pos := Point{0, 0}
	border := []Point{pos}
	m := 0
	n := 0
	oi := 0
	oj := 0
	for _, line := range text {
		parts := strings.Split(line, " ")
		d := rune(parts[0][0])
		times := parseInt(parts[1])
		direction := CHAR_TO_DIRECTION[d]
		for i := 0; i < times; i++ {
			pos = pos.move(direction)
			m = max(m, pos.i+1)
			n = max(n, pos.j+1)
			oi = min(oi, pos.i)
			oj = min(oj, pos.j)
			border = append(border, pos)
		}
	}

	m = m - oi
	n = n - oj
	puzzle := newMatrix(m, n)
	for _, point := range border {
		puzzle[point.i-oi][point.j-oj] = 1
	}

	expandedPuzzle := newMatrix(m+2, n+2)
	expandedPuzzle.copy(puzzle, Point{1, 1})

	empty := floodFill(expandedPuzzle, Point{0, 0})

	result := (m+2)*(n+2) - empty
	fmt.Println("Part 1 result:", result)
}
