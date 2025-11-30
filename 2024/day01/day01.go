package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("./day01/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(input io.Reader) int {
	leftList, rightList := partLists(input)

	sort.Ints(leftList)
	sort.Ints(rightList)

	var distance int

	for i := range leftList {
		distance += abs(leftList[i] - rightList[i])
	}

	return distance
}

func Part2(input io.Reader) int {
	leftList, rightList := partLists(input)

	rightOccurrences := make(map[int]int, len(rightList))
	for _, num := range rightList {
		occurred, exists := rightOccurrences[num]
		if !exists {
			rightOccurrences[num] = 1
		}

		rightOccurrences[num] = occurred + 1
	}

	var similarityScore int

	for _, num := range leftList {
		similarityScore += num * rightOccurrences[num]
	}

	return similarityScore
}

func partLists(input io.Reader) ([]int, []int) {
	var (
		leftList  []int
		rightList []int
	)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		currLn := scanner.Text()

		currLnSplit := strings.Split(currLn, "   ")

		left, err := strconv.Atoi(currLnSplit[0])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(currLnSplit[1])
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)

	}

	if len(leftList) != len(rightList) {
		panic("invalid list length")
	}

	return leftList, rightList
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
