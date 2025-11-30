package main

import (
	"strings"
	"testing"
)

const (
	example = `3   4
4   3
2   5
1   3
3   9
3   3`

	part1Solution = 11
	part2Solution = 31
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
