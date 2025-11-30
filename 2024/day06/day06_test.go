package main

import (
	"strings"
	"testing"
)

const (
	example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	part1Solution = 41
	part2Solution = 6
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
