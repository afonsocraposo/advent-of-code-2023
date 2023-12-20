package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func posStr(pos []int) string {
	return fmt.Sprintf("%d:%d", pos[0], pos[1])
}

func derivative(array []int) []int {
	d := make([]int, len(array)-1)
	for i, val := range array[1:] {
		d[i] = val - array[i]
	}
	return d
}

func scanLine(line string) int {
	points := 0
	inside := false
	for _, val := range line {
		if val == '|' || val == 'L' || val == 'J' {
			inside = !inside
		}else{
            if inside && val == '.' {
                points++
            }
        }
	}
	return points
}

func getStartPipe(directions []Direction) rune {
	if slices.Contains(directions, NORTH) {
		if slices.Contains(directions, WEST) {
			return 'J'
		}
		if slices.Contains(directions, EAST) {
			return 'L'
		}
		if slices.Contains(directions, SOUTH) {
			return '|'
		}
	}
	if slices.Contains(directions, EAST) {
		if slices.Contains(directions, WEST) {
			return '-'
		}
		if slices.Contains(directions, SOUTH) {
			return 'F'
		}
	}
	return '7'
}

func part2(filepath string) int {
	fmt.Println("Part 2")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	maze := make([]string, 0)
	start := []int{0, 0}
	n := 0
	for scanner.Scan() {
		line := scanner.Text()

		maze = append(maze, line)
		if index := strings.Index(line, "S"); index != -1 {
			start[0] = n
			start[1] = index
		}
		n++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	d := firstStep(maze, start)
	pos, comingFrom := moveTo(start, d[0])
	pipe, _ := getPipe(maze, pos)
	border := []string{posStr(start)}
	for !pipe.start {
		border = append(border, posStr(pos))
		pos, comingFrom = nextPos(maze, pos, pipe, comingFrom)
		pipe, _ = getPipe(maze, pos)
	}

	startPipe := getStartPipe(d)

	for i, line := range maze {
		newLine := []rune(line)
		for j := range newLine {
			if start[0] == i && start[1] == j {
				newLine[j] = startPipe
			} else {
				pos := posStr([]int{i, j})
				if !slices.Contains(border, pos) {
					newLine[j] = '.'
				}
			}
		}
		maze[i] = string(newLine)
		fmt.Println(parseLine(maze[i]))
	}

	area := 0
	for _, line := range maze {
		area += scanLine(line)
	}

	return area
}
