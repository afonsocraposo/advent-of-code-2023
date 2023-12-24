package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hash(s string) int {
    result := 0
    for _, val := range s {
        result += int(val)
        result *= 17
        result = result%256
    }
    return result
}

func part1(filepath string) int {
	fmt.Println("Part 1")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

    var steps []string
	for scanner.Scan() {
		line := scanner.Text()
        steps = strings.Split(line, ",")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	result := 0
    for _, step := range steps {
        result += hash(step)
    }

	return result
}
