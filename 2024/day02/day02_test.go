package main

import (
	"strings"
	"testing"
)

const (
	example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	part1Solution = 2
	part2Solution = 4
)

func TestPart1(t *testing.T) {
	got := Part1(strings.NewReader(example))

	if got != part1Solution {
		t.Errorf("expected: %d, got: %d", part1Solution, got)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(strings.NewReader(example))

	if got != part2Solution {
		t.Errorf("expected: %d, got: %d", part2Solution, got)
	}
}
