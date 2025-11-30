package main

import (
	"fmt"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

type pattern []string

func main() {
	input := helper.GetInputScanner("./day13/input.txt")

	var patterns []pattern
	currPattern := pattern{}
	for input.Scan() {
		currLn := input.Text()
		if currLn == "" {
			patterns = append(patterns, currPattern)
			currPattern = pattern{}
			continue
		}
		currPattern = append(currPattern, currLn)
	}
	patterns = append(patterns, currPattern)

	fmt.Println("part1:", reflectionSum(patterns, 0))
	fmt.Println("part2:", reflectionSum(patterns, 1))
}

func reflectionSum(patterns []pattern, smudges int) int {
	var sum int
	for _, pattern := range patterns {
		cols := findVerticalReflectionWithSmudges(pattern, smudges)
		if cols != -1 {
			sum += cols
			continue
		}

		rows := findHorizontalReflectionWithSmudges(pattern, smudges)
		if rows != -1 {
			sum += rows * 100
			continue
		}
	}

	return sum
}

func findVerticalReflectionWithSmudges(pattern pattern, expectedSmudges int) int {
	for i := 0; i < len(pattern[0])-1; i++ {
		if errors := getVerticalReflectionErrors(i, pattern); errors == expectedSmudges {
			return i + 1
		}
	}

	return -1
}

func getVerticalReflectionErrors(vLine int, pattern pattern) int {
	var errors int
	for i := 0; vLine-i >= 0 && vLine+1+i < len(pattern[0]); i++ {
		for _, patternLine := range pattern {
			if patternLine[vLine-i] != patternLine[vLine+1+i] {
				errors++
			}
		}
	}

	return errors
}

func findHorizontalReflectionWithSmudges(pattern pattern, expectedSmudges int) int {
	for i := 0; i < len(pattern)-1; i++ {
		if errors := getHorizontalReflectionErrors(i, pattern); errors == expectedSmudges {
			return i + 1
		}
	}

	return -1
}

func getHorizontalReflectionErrors(hLine int, pattern pattern) int {
	var errors int
	for i := 0; hLine-i >= 0 && hLine+1+i < len(pattern); i++ {
		for c := 0; c < len(pattern[0]); c++ {
			if pattern[hLine-i][c] != pattern[hLine+1+i][c] {
				errors++
			}
		}
	}

	return errors
}
