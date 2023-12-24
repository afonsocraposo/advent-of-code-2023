package main

import (
	"bufio"
	"fmt"
	"os"
)

func duplicate(puzzle [][]Space) [][]Space {
	result := make([][]Space, len(puzzle))
	for i, row := range puzzle {
		r := make([]Space, len(row))
        for j, s := range row {
            space := Space{map[Direction]bool{}, s.object}
            r[j] = space
        }
		result[i] = r
	}
    return result
}

func energized(puzzle [][]Space, start Beam) int {
	beams := []Beam{start}
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

func part2(filepath string) int {
	fmt.Println("Part 2")

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

	maxResult := 0
	for i := 0; i < len(puzzle); i++ {
		start := Beam{i, -1, RIGHT, false}
		result := energized(duplicate(puzzle), start)
		maxResult = max(result, maxResult)

		start = Beam{i, len(puzzle[0]), LEFT, false}
		result = energized(duplicate(puzzle), start)
		maxResult = max(result, maxResult)

		start = Beam{-1, i, DOWN, false}
		result = energized(duplicate(puzzle), start)
		maxResult = max(result, maxResult)

		start = Beam{len(puzzle), i, UP, false}
		result = energized(duplicate(puzzle), start)
		maxResult = max(result, maxResult)
	}
	return maxResult
}
