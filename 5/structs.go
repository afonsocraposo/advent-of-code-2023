package main

type EntryRange struct {
	start int
	end   int
}

func (e *EntryRange) includes(input int) bool {
	return input >= e.start && input <= e.end
}

type MapEntry struct {
	source      EntryRange
	destination EntryRange
}

func (m *MapEntry) sourceIncludes(input int) bool {
	return m.source.includes(input)
}

func (m *MapEntry) destIncludes(input int) bool {
	return m.destination.includes(input)
}

func (m *MapEntry) sourceTransform(input int) int {
	d := input - m.source.start
	return m.destination.start + d
}

func (m *MapEntry) destTransform(input int) int {
	d := input - m.destination.start
	return m.source.start + d
}

type Map struct {
	entries []MapEntry
}

func (m *Map) convert(input int) int {
	output := input
	for _, entry := range m.entries {
		if entry.sourceIncludes(input) {
			return entry.sourceTransform(input)
		}
	}
	return output
}

func (m *Map) invConvert(input int) int {
	output := input
	for _, entry := range m.entries {
		if entry.destIncludes(input) {
			return entry.destTransform(input)
		}
	}
	return output
}

type Maps struct {
	maps []Map
}

func (m *Maps) seedToLocation(input int) int {
	for _, m := range m.maps {
        input = m.convert(input)
	}
	return input
}

func (maps *Maps) locationToSeed(input int) int {
	for i := len(maps.maps) - 1; i >= 0; i-- {
		m := maps.maps[i]
        output := m.invConvert(input)
        input = output
	}
	return input
}
