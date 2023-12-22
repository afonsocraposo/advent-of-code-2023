package test

import (
	. "github.com/afonsocraposo/advent-of-code-2023/12/part1"
	"testing"
)

type testSimplify struct {
	record   string
	expected string
}

type testValid struct {
	record   string
	groups   []int
	expected bool
}

type testRecord struct {
	record   string
	groups   []int
	expected int
}

var recordsTests = []testRecord{
	{"???.###", []int{1, 1, 3}, 1},
	{".??..??...?##.", []int{1, 1, 3}, 4},
	{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1},
	{"????.#...#...", []int{4, 1, 1}, 1},
	{"????.######..#####.", []int{1, 6, 5}, 4},
	{"?###????????", []int{3, 2, 1}, 10},
}

func TestPossibleArrangements(t *testing.T) {
	for _, test := range recordsTests {
		if output := PossibleArrangements(test.record, test.groups); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

var simplifyRecordTests = []testSimplify{
	{"#.#.###", "#.#.###"},
	{".##.###", "##.###"},
	{"..#.###", "#.###"},
	{".#....#...###.", "#.#.###"},
	{".#..#..#...##.", "#.#.#.##"},
	{".#..?..?...##.", "#.?.?.##"},
}

func TestSimplifyRecord(t *testing.T) {
	for _, test := range simplifyRecordTests {
		if output := SimplifyRecord(test.record); output != test.expected {
			t.Errorf("Output %s not equal to expected %s, %s", output, test.expected, test.record)
		}
	}
}
