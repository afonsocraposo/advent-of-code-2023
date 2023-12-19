package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part1(filepath string) int {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	maps := Maps{}
	state := SEEDS

	var seeds []int
    m := Map{}
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case SEEDS:
			seedsStr, found := strings.CutPrefix(line, "seeds: ")
			if found {
				for _, s := range r.FindAllString(seedsStr, -1) {
					sv, err := strconv.Atoi(s)
					if err == nil {
						seeds = append(seeds, sv)
					}
				}
			}
		default:
			matches := r.FindAllString(line, -1)
			if len(matches) == 3 {
				drs, _ := strconv.Atoi(matches[0])
				srs, _ := strconv.Atoi(matches[1])
				rl, _ := strconv.Atoi(matches[2])

				source := EntryRange{start: srs, end: srs + rl - 1}
				destination := EntryRange{start: drs, end: drs + rl - 1}

				mapEntry := MapEntry{source, destination}
				m.entries = append(m.entries, mapEntry)
			}
		}
		if state != HUMIDITY_TO_LOCATION {
			prefix := STATE_TO_PREFIX[state]
			_, found := strings.CutPrefix(line, prefix)
			if found {
				if state != SEEDS {
					maps.maps = append(maps.maps, m)
				}
				state = STATE_TRANSITION[state]
				m = Map{}
			}
		}
	}
	if state == HUMIDITY_TO_LOCATION {
		maps.maps = append(maps.maps, m)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	minLoc := -1
	for _, seed := range seeds {
		location := maps.seedToLocation(seed)
		if minLoc == -1 || location < minLoc {
			minLoc = location
		}
	}

	return minLoc
}
