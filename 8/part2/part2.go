package part2

import (
	"bufio"
	"fmt"
	. "github.com/afonsocraposo/advent-of-code-2023/8/common"
	"os"
)

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func Part2(filepath string) int {
	fmt.Println("Part 2")

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

	var start []*Node
	nodesMap := make(map[string]*Node, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matches := R.FindAllString(line, -1)
		if len(matches) == 3 {
			nodeLabel := matches[0]
			leftLabel := matches[1]
			rightLabel := matches[2]

			n := GetNode(&nodesMap, nodeLabel)
			l := GetNode(&nodesMap, leftLabel)
			r := GetNode(&nodesMap, rightLabel)

			n.Label = nodeLabel
			n.Left = l
			n.Right = r
			n.End = nodeLabel[2] == END_RUNE

			if nodeLabel[2] == START_RUNE {
				start = append(start, n)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	nodesMap = nil

	nodes := start
    steps := make([]int, len(nodes))
	for index, n := range nodes {
		i := 0
		result := 0
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
        steps[index] = result
	}

    return LCM(steps[0], steps[1], steps[2:]...)
}
