package main

import (
	"bufio"
	"fmt"
	"os"
)

var CHAR_TO_OBJECT = map[rune]Object{
	'|':  VSPLITTER,
	'-':  HSPLITTER,
	'/':  FMIRROR,
	'\\': BMIRROR,
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Object int

const (
	EMPTY Object = iota
	FMIRROR
	BMIRROR
	HSPLITTER
	VSPLITTER
)

type Space struct {
	beams  map[Direction]bool
	object Object
}

type Beam struct {
	i         int
	j         int
	direction Direction
	stop      bool
}

func (b *Beam) move(puzzle [][]Space) *Beam {
	switch b.direction {
	case UP:
		b.i--
	case RIGHT:
		b.j++
	case DOWN:
		b.i++
	case LEFT:
		b.j--
	}
	if b.i >= len(puzzle) || b.i < 0 || b.j >= len(puzzle[0]) || b.j < 0 {
		b.stop = true
		return nil
	}
	s := puzzle[b.i][b.j]

    _, found := s.beams[b.direction]
    if found {
        b.stop = true
        return nil
    }
    s.beams[b.direction] = true

	if s.object == EMPTY {
		return nil
	}

	d1, d2 := s.nextDirections(b.direction)

	if d1 == -1 {
		b.stop = true
		return nil
	}
	b.direction = d1

	if d2 == -1 {
		return nil
	}

	newBeam := Beam{
		b.i,
		b.j,
		d2,
		b.stop,
	}
	return &newBeam
}

func (s *Space) nextDirections(input Direction) (Direction, Direction) {
	switch input {
	case UP:
		switch s.object {
		case FMIRROR:
			return RIGHT, -1
		case BMIRROR:
			return LEFT, -1
		case HSPLITTER:
			return LEFT, RIGHT
		}
	case RIGHT:
		switch s.object {
		case FMIRROR:
			return UP, -1
		case BMIRROR:
			return DOWN, -1
		case VSPLITTER:
			return UP, DOWN
		}
	case DOWN:
		switch s.object {
		case FMIRROR:
			return LEFT, -1
		case BMIRROR:
			return RIGHT, -1
		case HSPLITTER:
			return LEFT, RIGHT
		}
	case LEFT:
		switch s.object {
		case FMIRROR:
			return DOWN, -1
		case BMIRROR:
			return UP, -1
		case VSPLITTER:
			return UP, DOWN
		}
	}
	return input, -1
}

func printEnergy(puzzle [][]Space) {
	for _, row := range puzzle {
		for _, space := range row {
			if len(space.beams) > 0 {
                fmt.Print("#")
			}else{
                fmt.Print(".")
            }
		}
        fmt.Println()
	}
}

func part1(filepath string) int {
	fmt.Println("Part 1")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	puzzle := [][]Space{}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Space, len(line))
		for i, c := range line {
			row[i] = Space{
				map[Direction]bool{},
				CHAR_TO_OBJECT[c],
			}
		}
		puzzle = append(puzzle, row)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	beams := []Beam{{0, -1, RIGHT, false}}
	for len(beams) > 0 {
		beam := beams[0]
		for !beam.stop {
			newBeam := beam.move(puzzle)
			if newBeam != nil {
				beams = append(beams, *newBeam)
			}
		}
		beams = beams[1:]
	}

	result := 0
	for _, row := range puzzle {
		for _, space := range row {
			if len(space.beams) > 0 {
				result++
            }
		}
	}
	return result
}
