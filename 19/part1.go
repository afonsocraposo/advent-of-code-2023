package main

import (
	"fmt"

	. "github.com/afonsocraposo/advent-of-code-2023/utils"
)

func part1(filepath string) int {
	fmt.Println("Part 1")

	text := ReadTextFile(filepath)

	workflows := map[string]Workflow{}
	parts := []map[string]int{}

	p := false
	for _, line := range text {
		if line == "" {
			p = true
			continue
		}
		if p {
			part := ParsePart(line)
			parts = append(parts, part)
		} else {
			label, workflow := ParseWorkfowLine(line)
			workflows[label] = workflow
		}
	}

	result := 0
	for _, part := range parts {
		current := "in"
		for current != ACCEPT && current != REJECT {
			w := workflows[current]
			current = w.goTo(part)
		}

		if current == ACCEPT {
			for _, value := range part {
				result += value
			}
		}
	}

	return result
}
