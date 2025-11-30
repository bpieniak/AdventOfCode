package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputScan := helper.GetInputScanner("./day04/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day04/input.txt")
	part2(inputScan)
}

func part1(inputScan *bufio.Scanner) {
	var overlapCount int
	for inputScan.Scan() {
		currLn := inputScan.Text()
		assignments := strings.Split(currLn, ",")
		if overlaps := isFullyOverlaped(assignments[0], assignments[1]); overlaps {
			overlapCount += 1
		}
	}

	fmt.Println(overlapCount)
}

func isFullyOverlaped(range1str string, range2str string) bool {
	range1 := strings.Split(range1str, "-")
	lowRange1, _ := strconv.Atoi(range1[0])
	highRange1, _ := strconv.Atoi(range1[1])

	range2 := strings.Split(range2str, "-")
	lowRange2, _ := strconv.Atoi(range2[0])
	highRange2, _ := strconv.Atoi(range2[1])

	if lowRange1 >= lowRange2 && highRange1 <= highRange2 {
		return true
	} else if lowRange2 >= lowRange1 && highRange2 <= highRange1 {
		return true
	}

	return false
}

func part2(inputScan *bufio.Scanner) {
	var overlapCount int
	for inputScan.Scan() {
		currLn := inputScan.Text()
		assignments := strings.Split(currLn, ",")
		if overlaps := isOverlaped(assignments[0], assignments[1]); overlaps {
			overlapCount += 1
		}
	}

	fmt.Println(overlapCount)
}

func isOverlaped(range1str string, range2str string) bool {
	range1 := strings.Split(range1str, "-")
	lowRange1, _ := strconv.Atoi(range1[0])
	highRange1, _ := strconv.Atoi(range1[1])

	range2 := strings.Split(range2str, "-")
	lowRange2, _ := strconv.Atoi(range2[0])
	highRange2, _ := strconv.Atoi(range2[1])

	return isInside(lowRange1, lowRange2, highRange2) || isInside(lowRange2, lowRange1, highRange1) ||
		isInside(highRange1, lowRange2, highRange2) || isInside(highRange2, lowRange1, highRange1)
}

func isInside(checked int, low int, high int) bool {
	return low <= checked && checked <= high
}
