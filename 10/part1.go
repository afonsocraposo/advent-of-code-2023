package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(filepath string) int {
	fmt.Println("Part 1")

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

        fmt.Println(parseLine(line))

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
	steps := 1
	pipe, _ := getPipe(maze, pos)
	for !pipe.start {
		pos, comingFrom = nextPos(maze, pos, pipe, comingFrom)
		pipe, _ = getPipe(maze, pos)
		steps++
	}

	return steps / 2
}
