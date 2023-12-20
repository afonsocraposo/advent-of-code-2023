package main

import (
	"errors"
)

type Pipe struct {
	pipe       rune
	directions []Direction
	start      bool
}

func (pipe *Pipe) goTo(comingFrom Direction) Direction {
	if pipe.directions[0] == comingFrom {
		return pipe.directions[1]
	} else {
		return pipe.directions[0]
	}
}

func (pipe *Pipe) validComingFrom(comingFrom Direction) bool {
	return pipe.directions[0] == comingFrom || pipe.directions[1] == comingFrom
}

var pipes = map[rune]Pipe{
	'|': {'|', []Direction{NORTH, SOUTH}, false},
	'-': {'-', []Direction{EAST, WEST}, false},
	'L': {'L', []Direction{NORTH, EAST}, false},
	'J': {'J', []Direction{NORTH, WEST}, false},
	'7': {'7', []Direction{SOUTH, WEST}, false},
	'F': {'F', []Direction{SOUTH, EAST}, false},
	'S': {'S', nil, true},
}

var charToBox = map[rune]rune{
	'|': '│',
	'-': '─',
	'L': '└',
	'J': '┘',
	'7': '┐',
	'F': '┌',
	'.': '·',
}

func parseLine(line string) string {
	parsedLine := ""
	for _, val := range line {
		char, ok := charToBox[val]
		if !ok {
			char = val
		}
		parsedLine += string(char)
	}
	return parsedLine
}

const NORTH_PIPES = "|LJ'"
const EAST_PIPES = "-LF"
const SOUTH_PIPES = "|7F"
const WEST_PIPES = "-J7"

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func getPipe(maze []string, pos []int) (Pipe, error) {
	pipeType := rune(maze[pos[0]][pos[1]])
	pipe, valid := pipes[pipeType]
	if !valid {
		return Pipe{}, errors.New("not a pipe")
	}
	return pipe, nil
}

func getPos(maze []string, pos []int, direction Direction) []int {
	switch direction {
	case NORTH:
		return []int{pos[0] - 1, pos[1]}
	case EAST:
		return []int{pos[0], pos[1] + 1}
	case SOUTH:
		return []int{pos[0] + 1, pos[1]}
	case WEST:
		return []int{pos[0], pos[1] - 1}
	}
	panic("Invalid parameters")
}

func firstStep(maze []string, start []int) []Direction {
    d := []Direction{}
	if start[0] > 0 {
		// coming from SOUTH
		nextPos := []int{start[0] - 1, start[1]}
		pipe, err := getPipe(maze, nextPos)
		if err == nil {
			if pipe.validComingFrom(SOUTH) {
				d = append(d, NORTH)
			}
		}
	}
	if start[0] < len(maze)-1 {
		// coming from NORTH
		nextPos := []int{start[0] + 1, start[1]}
		pipe, err := getPipe(maze, nextPos)
		if err == nil {
			if pipe.validComingFrom(NORTH) {
				d = append(d, SOUTH)
			}
		}
	}
	if start[1] > 0 {
		// coming from EAST
		nextPos := []int{start[0], start[1] - 1}
		pipe, err := getPipe(maze, nextPos)
		if err == nil {
			if pipe.validComingFrom(EAST) {
				d = append(d, WEST)
			}
		}
	}
	if start[1] < len(maze[0])-1 {
		// coming from WEST
		nextPos := []int{start[0], start[1] + 1}
		pipe, err := getPipe(maze, nextPos)
		if err == nil {
			if pipe.validComingFrom(WEST) {
				d = append(d, EAST)
			}
		}
	}
    return d
}

func moveTo(pos []int, d Direction) ([]int, Direction) {
	switch d {
	case NORTH:
		return []int{pos[0] - 1, pos[1]}, SOUTH
	case EAST:
		return []int{pos[0], pos[1] + 1}, WEST
	case SOUTH:
		return []int{pos[0] + 1, pos[1]}, NORTH
	case WEST:
		return []int{pos[0], pos[1] - 1}, EAST
	}
    panic("Invalid direction")
}

func nextPos(maze []string, pos []int, pipe Pipe, comingFrom Direction) ([]int, Direction) {
	nextDirection := pipe.goTo(comingFrom)
    return moveTo(pos, nextDirection)
}
