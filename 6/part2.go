package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2(filepath string) int {
	fmt.Println("Part 2")

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

    time, _ := strconv.Atoi(strings.Join(times, ""))
    record, _ := strconv.Atoi(strings.Join(records, ""))
    race := Race{time, record}
    return race.waysToBeatRecord()
}
