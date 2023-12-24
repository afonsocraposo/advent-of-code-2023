package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "regexp"
)

type Lens struct {
	label string
	focal int
}

var r = regexp.MustCompile("[a-z]+")

func addLens(boxes map[int][]Lens, lens Lens) {
	box := hash(lens.label)
	lenses, found := boxes[box]
	if !found {
		boxes[box] = []Lens{lens}
	} else {
		for i, l := range lenses {
			if l.label == lens.label {
				boxes[box][i].focal = lens.focal
				return
			}
		}
		boxes[box] = append(lenses, lens)
	}
}
func removeLabel(boxes map[int][]Lens, label string) {
	box := hash(label)
	lenses, found := boxes[box]
	if !found {
		return
	} else {
		index := -1
		for i, l := range lenses {
			if l.label == label {
				index = i
			}
		}
		if index > -1 {
			lenses = append(lenses[:index], lenses[index+1:]...)
			boxes[box] = lenses
		}
	}
}

func part2(filepath string) int {
	fmt.Println("Part 2")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var steps []string
	for scanner.Scan() {
		line := scanner.Text()
		steps = strings.Split(line, ",")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	boxes := map[int][]Lens{}
	for _, step := range steps {
		label := r.FindString(step)
		action := step[len(label)]
		if action == '=' {
			focal, _ := strconv.Atoi(step[len(label)+1:])
			lens := Lens{label, focal}
			addLens(boxes, lens)
		} else {
			removeLabel(boxes, label)
		}
	}

	result := 0
	for box, lenses := range boxes {
		for slot, lens := range lenses {
			result += (box + 1) * (slot+1) * lens.focal
		}
	}

	return result
}
