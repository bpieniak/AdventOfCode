package main

import (
	"strings"
	"testing"
)

const (
	example = `2333133121414131402`

	part1Solution = 1928
	part2Solution = 2858
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
