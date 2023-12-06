package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Bag struct {
	red   int
	green int
	blue  int
}

var b = Bag{12, 13, 14}

var rexp = regexp.MustCompile(`([0-9]+) red`)
var gexp = regexp.MustCompile(`([0-9]+) green`)
var bexp = regexp.MustCompile(`([0-9]+) blue`)

func validateColor(r *regexp.Regexp, line string, t int) bool {
	counts := r.FindAllStringSubmatch(line, -1)
	for _, c := range counts {
		number, err := strconv.Atoi(c[1])
		if err == nil && number > t {
			return false
		}
	}
	return true
}

func maxColor(r *regexp.Regexp, line string) int {
	counts := r.FindAllStringSubmatch(line, -1)
	m := 0
	for _, c := range counts {
		number, err := strconv.Atoi(c[1])
		if err == nil {
			if number > m {
				m = number
			}
		}
	}
	return m
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

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	fmt.Println("Calibration values:")
	result := 0
	n := 1
	for scanner.Scan() {
		line := scanner.Text()

		if part == 1 {
			valid := validateColor(rexp, line, b.red) &&
				validateColor(gexp, line, b.green) &&
				validateColor(bexp, line, b.blue)

			if valid {
				result += n
			}
		} else {
			maxr := maxColor(rexp, line)
			maxg := maxColor(gexp, line)
			maxb := maxColor(bexp, line)
			power := maxr * maxg * maxb
			result += power
			fmt.Println(power)
		}

		n++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("\nThe result is %d\n", result)
}
