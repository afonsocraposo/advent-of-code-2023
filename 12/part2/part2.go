package part2

import (
	"bufio"
	"fmt"
	. "github.com/afonsocraposo/advent-of-code-2023/12/part1"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string, N int) (string, []int) {
	r, groupsStr, _ := strings.Cut(line, " ")

	numbersStr := strings.Split(groupsStr, ",")
	l := len(numbersStr)
	groups := make([]int, N*l)
	n := 0
	record := ""
	for n < N {
		if n > 0 {
			record += "?" + r
		} else {
			record = r
		}
		for i, s := range numbersStr {
			groups[i+l*n], _ = strconv.Atoi(s)
		}
		n++
	}
	return record, groups
}

func Part2(filepath string) int {
	fmt.Println("Part 2")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		N := 5
		record, groups := parseLine(line, N)
		arrangements := PossibleArrangements(record, groups)
		result += arrangements
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
