package part1

import (
	"bufio"
	"fmt"
	. "github.com/afonsocraposo/advent-of-code-2023/8/common"
	"os"
)

func getNode(nodes *map[string]*Node, nodeLabel string) *Node {
	n, ok := (*nodes)[nodeLabel]
	if !ok {
		n := &Node{}
		(*nodes)[nodeLabel] = n
		return n
	}
	return n
}

func Part1(filepath string) int {
	fmt.Println("Part 1")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var instructions string
	if scanner.Scan() {
		instructions = scanner.Text()
	}

    var start *Node
	nodes := make(map[string]*Node, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matches := R.FindAllString(line, -1)
		if len(matches) == 3 {
			nodeLabel := matches[0]
			leftLabel := matches[1]
			rightLabel := matches[2]

			n := getNode(&nodes, nodeLabel)
			l := getNode(&nodes, leftLabel)
			r := getNode(&nodes, rightLabel)

			n.Left = l
			n.Right = r
			n.End = nodeLabel == END_LABEL

            if nodeLabel == START_LABEL {
                start = n
            }
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

    nodes = nil

	result := 0
    n := start
	i := 0
	for !n.End {
		if i >= len(instructions) {
			i = 0
		}
		if instructions[i] == 'L' {
			n = n.Left
		} else {
			n = n.Right
		}
		i++
		result++
	}

	return result
}
