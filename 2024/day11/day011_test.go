package main

import (
	"strings"
	"testing"
)

const (
	example = `125 17`

	part1Solution = 55312
)

func TestPart1(t *testing.T) {
	got := Part1(strings.NewReader(example))

	if got != part1Solution {
		t.Errorf("expected: %d, got: %d", part1Solution, got)
	}
}
