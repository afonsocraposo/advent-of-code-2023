package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2(filepath string) int {
	fmt.Println("Part 2")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// line := scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	result := 0

	return result
}
