package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type seedRange EntryRange

func containsSeed(seeds []seedRange, seed int) bool {
    for _, sr := range(seeds) {
        if seed >= sr.start && seed <= sr.end {
            return true
        }
    }
    return false
}

func part2(filepath string) int {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	maps := Maps{}
	state := SEEDS

	var seeds []seedRange
    m := Map{}
	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case SEEDS:
			seedsStr, found := strings.CutPrefix(line, "seeds: ")
			if found {
				seedStart := 0
				for i, s := range r.FindAllString(seedsStr, -1) {
					sv, err := strconv.Atoi(s)
					if err == nil {
						if i%2 == 1 {
							seeds = append(seeds, seedRange{seedStart, seedStart + sv - 1})
						} else {
							seedStart = sv
						}
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

    maxLoc := -1
    humToLoc := maps.maps[len(maps.maps)-1]
    for _, mapEntry := range(humToLoc.entries) {
        loc := mapEntry.destination.end
        if loc > maxLoc {
            maxLoc = loc
        }
    }

    for loc := 0; loc <= maxLoc; loc++ {
        seed := maps.locationToSeed(loc)
        if containsSeed(seeds, seed) {
            return loc
        }
    }

	return -1
}
