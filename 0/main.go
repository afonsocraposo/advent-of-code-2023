package main

import (
	"fmt"
	"os"
)

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
