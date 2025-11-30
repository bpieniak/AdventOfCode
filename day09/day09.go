package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

func main() {
	inputScan := helper.GetInputScanner("./day09/input.txt")

	var input [][]int
	for inputScan.Scan() {
		currLn := inputScan.Text()
		nums := parseNums(currLn)
		input = append(input, nums)
	}

	var sumPart1, sumPart2 int
	for _, line := range input {
		first, last := predictNext(line)
		sumPart1 += last
		sumPart2 += first
	}

	fmt.Println(sumPart1, sumPart2)
}

func predictNext(nums []int) (int, int) {
	var firstValues, lastValues []int
	var currLine = nums
	for {
		if isAllZeros(currLine...) {
			break
		}
		lastValues = append(lastValues, currLine[len(currLine)-1])
		firstValues = append(firstValues, currLine[0])

		currLine = getDifferences(currLine)
	}

	currExtrapolatedLast := 0
	for i := len(lastValues) - 1; i >= 0; i-- {
		currExtrapolatedLast = lastValues[i] + currExtrapolatedLast
	}

	currExtrapolatedFirst := 0
	for i := len(firstValues) - 1; i >= 0; i-- {
		currExtrapolatedFirst = firstValues[i] - currExtrapolatedFirst
	}

	return currExtrapolatedFirst, currExtrapolatedLast
}

func getDifferences(nums []int) []int {
	var differences []int
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		differences = append(differences, diff)
	}
	return differences
}

func isAllZeros(nums ...int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func parseNums(n string) []int {
	regx := regexp.MustCompile(`-{0,1}\d+`)

	foundNums := regx.FindAllString(n, -1)

	nums := make([]int, 0, len(foundNums))
	for _, num := range foundNums {
		numInt, _ := strconv.Atoi(num)
		nums = append(nums, numInt)
	}

	return nums
}
