package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("./day02/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(input io.Reader) (result int) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		currLn := scanner.Text()

		rules := stringToInts(currLn)

		if isSafe(rules) {
			result++
		}
	}
	return result
}

func Part2(input io.Reader) (result int) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		currLn := scanner.Text()

		rule := stringToInts(currLn)

		if isSafe(rule) || isSafeWithDamp(rule) {
			result++
		}

	}

	return result
}

func isSafe(rule []int) bool {
	isRuleDecreasing := isDecreasing(rule[0], rule[1])

	for levelIndex := 1; levelIndex < len(rule); levelIndex++ {
		// The levels are either all increasing or all decreasing.
		if isRuleDecreasing != isDecreasing(rule[levelIndex-1], rule[levelIndex]) {
			return false
		}

		// Any two adjacent levels differ by at least one and at most three.
		absDiff := abs(rule[levelIndex-1] - rule[levelIndex])
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return true
}

func isSafeWithDamp(rule []int) bool {
	for levelToDump := 0; levelToDump < len(rule); levelToDump++ {
		newRule := removeIndex(slices.Clone(rule), levelToDump)

		if isSafe(newRule) {
			return true
		}
	}

	return false
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func isDecreasing(a, b int) bool {
	return b > a
}

func stringToInts(line string) []int {
	split := strings.Split(line, " ")

	ints := make([]int, 0, len(split))
	for _, numStr := range split {
		numInt, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		ints = append(ints, numInt)
	}

	return ints
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
