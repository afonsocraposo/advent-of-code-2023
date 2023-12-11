package main

import (
	"bufio"
	"os"
	"strconv"
)

func part1(filepath string) int {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

    var pn []PartNumber
    s := make(map[int]bool)

	n := 1
	for scanner.Scan() {
		line := scanner.Text()

        // get all the numbers in the line
        r := numbers.FindAllStringIndex(line, -1)
        for _, m := range(r) {
            numberStr := line[m[0]:m[1]]
            number, err := strconv.Atoi(numberStr)
            if err == nil {
                pn = append(pn, PartNumber{number, n, m})
            }
        }

        // get all the symbols in the line and create a hash map
        r = symbols.FindAllStringIndex(line, -1)
        for _, m := range(r) {
            h := hash(n, m[0])
            s[h] = true
        }

		n++
	}

    result := 0
    // iterate all part numbers
    for _, p := range(pn) {
        // iterate all adjacent positions of part number
        for _, pos := range(getAdjPos(p)) {
            // check if position's hash is in symbols hash map
            _, ok := s[pos]
            if ok {
                result += p.number
                break
            }
        }
    }

	if err := scanner.Err(); err != nil {
		panic(err)
	}

    return result
}


