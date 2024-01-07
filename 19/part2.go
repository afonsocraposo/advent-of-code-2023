package main

import (
	"fmt"
	"maps"

	. "github.com/afonsocraposo/advent-of-code-2023/utils"
)

func processPart(nodes map[string]Node, part map[string]Range, current string) int {
	if current == ACCEPT {
		result := 1
		for _, r := range part {
			result *= r.End - r.Start + 1
		}
		return result
	} else if current == REJECT {
		return 0
	}
	n := nodes[current]

	p1 := map[string]Range{}
	maps.Copy(p1, part)
	p2 := map[string]Range{}
	maps.Copy(p2, part)

	pr1 := part[n.label]
	pr2 := part[n.label]
	if n.Condition.greater {
		pr1.Start = max(pr1.Start, n.value + 1)
		pr2.End = min(pr2.End, n.value)
	} else {
		pr1.End = min(pr1.End, n.value - 1)
		pr2.Start = max(pr2.Start, n.value)
	}
	p1[n.label] = pr1
	p2[n.label] = pr2

	return processPart(nodes, p1, n.dest) + processPart(nodes, p2, n.altDest)
}

func part2(filepath string) int {
	fmt.Println("Part 2")
	text := ReadTextFile(filepath)

	nodes := map[string]Node{}

	for _, line := range text {
		if line == "" {
			break
		}
		n := ParseNodesLine(line)
		maps.Copy(nodes, n)
	}

	part := map[string]Range{
		"x": {Start: 1, End: 4000},
		"m": {Start: 1, End: 4000},
		"a": {Start: 1, End: 4000},
		"s": {Start: 1, End: 4000},
	}

	current := "in"
	result := processPart(nodes, part, current)
	return result
}
