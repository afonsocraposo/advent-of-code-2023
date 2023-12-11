package main

import (
	"fmt"
	"os"
	"regexp"
)

type PartNumber struct {
    number int;
    row int;
    columns []int;
}

var numbers = regexp.MustCompile(`[0-9]+`)
var symbols = regexp.MustCompile(`[^0-9.]`)
var gears = regexp.MustCompile(`\*`)

// compute hash for position (row, col)
func hash(row int, col int) int {
   return 6581*row + 7919*col
}

// get positions adjacent to the part number
func getAdjPos(p PartNumber) []int {
    size := 2*(p.columns[1]-p.columns[0] + 2) + 2
    pos := make([]int, size)

    pos[0] = hash(p.row, p.columns[0]-1)
    pos[1] = hash(p.row, p.columns[1])

    n := 2
    for i := p.columns[0]-1; i <= p.columns[1]; i++ {
        pos[n] = hash(p.row - 1, i)
        pos[n+1] = hash(p.row + 1, i)
        n += 2
    }

    return pos
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Provide an input file")
	}
	filepath := args[0]
	var part int
	if len(args) > 1 && args[1] == "2" {
		part = 2
	} else {
		part = 1
	}

    var result int
    if part == 1 {
        result = part1(filepath)
    } else {
        result = part2(filepath)
    }
	fmt.Printf("\nThe result is %d\n", result)
}
