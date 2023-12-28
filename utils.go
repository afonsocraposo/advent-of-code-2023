package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	i int
	j int
}

func readTextFile(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

type Direction Point

var (
	UP    = Direction{-1, 0}
	RIGHT = Direction{0, 1}
	DOWN  = Direction{1, 0}
	LEFT  = Direction{0, -1}
)

var DIRECTIONS = []Direction{UP, RIGHT, DOWN, LEFT}

var CHAR_TO_DIRECTION = map[rune]Direction{
	'U': UP,
	'R': RIGHT,
	'D': DOWN,
	'L': LEFT,
}

func (point *Point) move(direction Direction) Point {
	return Point{point.i + direction.i, point.j + direction.j}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error parsing int %s: %s\n", s, err)
		return 0
	}
	return i
}

func printText(text []string) {
	for _, line := range text {
		fmt.Println(line)
	}
}

func replaceInString(s string, i int, r string) string {
	return s[:i] + r + s[i+1:]
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

type Matrix [][]int

func newMatrix(m int, n int) Matrix {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func (m *Matrix) get(point Point) int {
	return (*m)[point.i][point.j]
}

func (m *Matrix) inside(point Point) bool {
	return point.i >= 0 && point.j >= 0 && point.i < len(*m) && point.j < len((*m)[0])
}

func (m1 *Matrix) copy(m2 Matrix, start Point) {
	for i, row := range m2 {
		for j, v := range row {
			(*m1)[start.i+i][start.j+j] = v
		}
	}
}

func printMatrix(matrix Matrix) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}
