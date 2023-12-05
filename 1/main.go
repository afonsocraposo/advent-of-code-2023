package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"regexp"
)

func parseValue(value string) string {
	if len(value) == 1 {
		return value
	}

	switch value {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	panic("Invalid value to parse")
}

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        panic("Provide an input file")
    }
    filepath := args[0]

    var r *regexp.Regexp
    if len(args) > 1 && args[1] == "2" {
		r = regexp.MustCompile(`([1-9]|one|two|three|four|five|six|seven|eight|nine)`)
    }else {
		r = regexp.MustCompile(`([1-9])`)
    }

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

    fmt.Println("Calibration values:")
	result := 0
	for scanner.Scan() {
		line := scanner.Text()

        matches := []string{}
        for i := 0; i < len(line); i++ {
            for _, match := range r.FindAllString(line[i:], -1) {
                matches = append(matches, match)
            }
        }
		left := matches[0]
		right := matches[len(matches)-1]

        numberStr := parseValue(left) + parseValue(right)
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
        result += number

        fmt.Println(line, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nThe result is %d\n", result)
}
