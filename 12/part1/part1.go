package part1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = map[string]int{}

func hash(r string, g []int) string {
	return fmt.Sprintf("%s:%v", r, g)
}

func PossibleArrangements(record string, groups []int) int {
	result, found := cache[hash(record, groups)]
	if found {
		return result
	}
	if len(record) == 0 {
		if len(groups) > 0 {
			return 0
		} else {
			return 1
		}
	}
	c := record[0]
	if c == '?' {
		r := record[1:]
		g := groups
		p1 := PossibleArrangements(r, g)
		p2 := PossibleArrangements("#"+r, g)
		cache[hash(r, g)] = p1
		cache[hash("#"+r, g)] = p2
		return p1 + p2
	} else if c == '#' {
		if len(groups) == 0 {
			return 0
		}
		g := groups[0]
		if g > len(record) {
			return 0
		}
		for i := 1; i < g; i++ {
			if record[i] == '.' {
				return 0
			}
		}
		if len(record) > g {
			if record[g] == '#' {
				return 0
			} else if record[g] == '?' {
				record = record[:g] + "." + record[g+1:]
			}
		}
        p := PossibleArrangements(record[g:], groups[1:])
        cache[hash(record[g:], groups[1:])] = p
        return p
	} else if c == '.' {
        p := PossibleArrangements(record[1:], groups)
        cache[hash(record[1:], groups)] = p
        return p
	}
	return 0
}

func Part1(filepath string) int {
	fmt.Println("Part 1")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		record, groupsStr, _ := strings.Cut(line, " ")
		numbersStr := strings.Split(groupsStr, ",")
		groups := make([]int, len(numbersStr))
		for i, s := range numbersStr {
			groups[i], _ = strconv.Atoi(s)
		}
		arrangements := PossibleArrangements(record, groups)
		result += arrangements
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
