package main

type Race struct {
	time     int
	record int
}

func (race *Race) distance(p int) int {
	return race.time*p - p*p
}

func (race *Race) waysToBeatRecord() int {
	result := 0
	for p := 0; p <= race.time; p++ {
        if race.distance(p) > race.record {
            result ++
        }
	}
    return result
}
