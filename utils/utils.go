package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	I int
	J int
}

type Range struct {
    Start int
    End int
}

func ReadTextFile(filepath string) []string {
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

func (direction *Direction) Scale(n int) Direction {
	return Direction{direction.I * n, direction.J * n}
}

var CHAR_TO_DIRECTION = map[rune]Direction{
	'U': UP,
	'R': RIGHT,
	'D': DOWN,
	'L': LEFT,
}

func (point *Point) Move(direction Direction) Point {
	return Point{point.I + direction.I, point.J + direction.J}
}

func (p1 *Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.I-p1.I), 2) + math.Pow(float64(p2.J-p1.J), 2))
}

func (point *Point) Inside(m int, n int) bool {
	return point.I >= 0 && point.J >= 0 && point.I < m && point.J < n
}

func (point *Point) ToString() string {
	return fmt.Sprintf("%d:%d", point.I, point.J)
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error parsing int %s: %s\n", s, err)
		return 0
	}
	return i
}

func PrintText(text []string) {
	for _, line := range text {
		fmt.Println(line)
	}
}

func ReplaceInString(s string, i int, r string) string {
	return s[:i] + r + s[i+1:]
}

func Transpose(pattern [][]int) [][]int {
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

func NewMatrix(m int, n int) Matrix {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func (m *Matrix) Get(point Point) int {
	return (*m)[point.I][point.J]
}

func (m *Matrix) Inside(point Point) bool {
	return point.I >= 0 && point.J >= 0 && point.I < len(*m) && point.J < len((*m)[0])
}

func (m1 *Matrix) Copy(m2 Matrix, start Point) {
	for i, row := range m2 {
		for j, v := range row {
			(*m1)[start.I+i][start.J+j] = v
		}
	}
}

func (matrix *Matrix) Size() (m int, n int) {
	return len(*matrix), len((*matrix)[0])
}

func PrintMatrix(matrix Matrix) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}

func WaitForInput() {
	fmt.Scanln()
}
