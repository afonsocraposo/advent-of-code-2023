package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	. "github.com/afonsocraposo/advent-of-code-2023/utils"
)

func floodFillMatrix(puzzle Matrix, start Point) int {
	visited := map[Point]bool{start: true}
	queue := []Point{start}
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		for _, direction := range DIRECTIONS {
			point := start.Move(direction)
			if puzzle.Inside(point) {
				_, found := visited[point]
				if puzzle.Get(point) == 0 && !found {
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
			point := start.Move(direction)
			if point.Inside(m, n) {
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
		s += p1.J*p2.I - p1.I*p2.J
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
	pos := Point{I: 0, J: 0}
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
			times = ParseInt(parts[1])
			direction = CHAR_TO_DIRECTION[d]
		} else {
			direction, times = day18ParseHex(parts[2][2:8])
		}
		pos = pos.Move(direction.Scale(times))
		m = max(m, pos.I+1)
		n = max(n, pos.J+1)
		oi = min(oi, pos.I)
		oj = min(oj, pos.J)
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
		result += p1.Distance(p2)
	}
	return int(result)
}

func day18(filepath string) {
	text := ReadTextFile(filepath)

	for _, part := range []int{1, 2} {
		vertices, _, _, _, _ := getVertices(text, part)
		s := shoelaceFormula(vertices)
		p := perimeter(vertices)
		result := s + (p / 2) + 1
		fmt.Printf("Part %d result: %d\n", part, result)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Provide a day number")
	}
	day := args[0]
	fmt.Printf("Day %s\n", day)

	filepath := fmt.Sprintf("example/day%s.txt", day)
	if len(args) > 1 {
		filepath = args[1]
	}
	fmt.Printf("File: %s\n\n", filepath)

	day18(filepath)
}
