package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var r = regexp.MustCompile("[0-9]+")

func part1(filepath string) int {
	fmt.Println("Part 1")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	times := r.FindAllString(line, -1)

	scanner.Scan()
	line = scanner.Text()
	records := r.FindAllString(line, -1)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if len(times) != len(records) {
		msg := fmt.Sprintf("times and records should have the same length: %s %s", times, records)
		panic(msg)
	}

    result := 1
	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		record, _ := strconv.Atoi(records[i])
        race := Race{time, record}
        result *= race.waysToBeatRecord()
	}

	return result
}
