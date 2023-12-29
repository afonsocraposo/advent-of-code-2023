package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func floodFillMatrix(puzzle Matrix, start Point) int {
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

func floodFillBorderMap(borderMap map[Point]bool, m int, n int, start Point) int {
	visited := map[Point]bool{start: true}
	queue := []Point{start}
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		for _, direction := range DIRECTIONS {
			point := start.move(direction)
			if point.inside(m, n) {
				_, isBorder := borderMap[point]
				_, didVisit := visited[point]
				if !isBorder && !didVisit {
					visited[point] = true
					queue = append(queue, point)
				}
			}
		}
	}
	return len(visited)
}

func shoelaceFormula(vertices []Point) int {
	s := 0
	v := append(vertices, vertices[0])
	for i := range v[1:] {
		p1 := v[i]
		p2 := v[i+1]
		s += p1.j*p2.i - p1.i*p2.j
	}
	return int(0.5 * math.Abs(float64(s)))
}

func day18ParseHex(hex string) (Direction, int) {
	dir := hex[5]
	var direction Direction
	switch dir {
	case '0':
		direction = RIGHT
	case '1':
		direction = DOWN
	case '2':
		direction = LEFT
	case '3':
		direction = UP
	}

	distance, _ := strconv.ParseInt(hex[:5], 16, 32)

	return direction, int(distance)
}

func getVertices(text []string, part int) ([]Point, int, int, int, int) {
	pos := Point{0, 0}
	border := []Point{pos}
	m := 0
	n := 0
	oi := 0
	oj := 0
	for _, line := range text[:len(text)-1] {
		parts := strings.Split(line, " ")
		var direction Direction
		var times int
		if part == 1 {
			d := rune(parts[0][0])
			times = parseInt(parts[1])
			direction = CHAR_TO_DIRECTION[d]
		} else {
			direction, times = day18ParseHex(parts[2][2:8])
		}
		pos = pos.move(direction.scale(times))
		m = max(m, pos.i+1)
		n = max(n, pos.j+1)
		oi = min(oi, pos.i)
		oj = min(oj, pos.j)
		border = append(border, pos)
	}

	return border, m, n, oi, oj
}

func perimeter(vertices []Point) int {
	v := append(vertices, vertices[0])
	result := 0.0
	for i := range v[1:] {
		p1 := v[i]
		p2 := v[i+1]
		result += p1.distance(p2)
	}
	return int(result)
}

func day18(filepath string) {
	text := readTextFile(filepath)

	for _, part := range []int{1, 2} {
		vertices, _, _, _, _ := getVertices(text, part)
		s := shoelaceFormula(vertices)
        p := perimeter(vertices)
        result := s + (p/2) + 1
		fmt.Printf("Part %d result: %d\n", part, result)
	}
}
