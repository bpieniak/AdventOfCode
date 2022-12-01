package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	inputScan := helper.GetInputScanner("./day01/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day01/input.txt")
	part2(inputScan)
}

func part1(inputScan *bufio.Scanner) {
	totalCals := getElfCalories(inputScan)

	_, max := helper.MinMax(totalCals)
	fmt.Println(max)
}

func part2(inputScan *bufio.Scanner) {
	totalCals := getElfCalories(inputScan)

	sort.Ints(totalCals)

	top3Sum := 0
	for _, v := range totalCals[len(totalCals)-3:] {
		top3Sum += v
	}

	fmt.Println(top3Sum)
}

func getElfCalories(inputScan *bufio.Scanner) []int {
	var totalCals []int

	currElfCals := 0
	for inputScan.Scan() {
		currLn := inputScan.Text()
		if currLn == "" {
			totalCals = append(totalCals, currElfCals)
			currElfCals = 0
			continue
		}

		cals, _ := strconv.Atoi(currLn)
		currElfCals += cals
	}

	return totalCals
}
