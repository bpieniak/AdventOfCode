package main

import (
	"strings"
	"testing"
)

const (
	example1 = `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`
	example1Solution = 2

	example2 = `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`
	example2Solution = 4

	example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	part1Solution = 36
	part2Solution = 81
)

func TestPart1Example1(t *testing.T) {
	got := Part1(strings.NewReader(example1))

	if got != example1Solution {
		t.Errorf("expected: %d, got: %d", example1Solution, got)
	}
}

func TestPart1Example2(t *testing.T) {
	got := Part1(strings.NewReader(example2))

	if got != example2Solution {
		t.Errorf("expected: %d, got: %d", example2Solution, got)
	}
}

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
