package main

import (
	"strings"
	"testing"
)

const (
	example1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	example2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	part1Solution = 161
	part2Solution = 48
)

func TestPart1(t *testing.T) {
	got := Part1(strings.NewReader(example1))

	if got != part1Solution {
		t.Errorf("expected: %d, got: %d", part1Solution, got)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(strings.NewReader(example2))

	if got != part2Solution {
		t.Errorf("expected: %d, got: %d", part2Solution, got)
	}
}
