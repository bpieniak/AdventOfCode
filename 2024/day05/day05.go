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
	inputFile, err := os.Open("./day05/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(inputFile io.Reader) (result int) {
	rules, updates := parseRulesAndUpdates(inputFile)

	for _, update := range updates {
		if verifyUpdate(update, rules) {
			result += update[len(update)/2]
		}
	}

	return result
}

func Part2(inputFile io.Reader) (result int) {
	rules, updates := parseRulesAndUpdates(inputFile)

	for _, update := range updates {
		if verifyUpdate(update, rules) {
			continue
		}

		slices.SortFunc(update, func(page1, page2 int) int {
			pagesBefore1 := rules[page1]
			pagesBefore2 := rules[page2]

			if slices.Contains(pagesBefore2, page1) {
				return 1
			}

			if slices.Contains(pagesBefore1, page2) {
				return -1
			}

			return 0
		})

		result += update[len(update)/2]
	}

	return result
}

func verifyUpdate(pages []int, rules map[int][]int) bool {
	pageOrder := make(map[int]int, len(pages))

	for index, page := range pages {
		pageOrder[page] = index
	}

	for currPageIndex, page := range pages {
		pagesRequiredBefore, exists := rules[page]
		if !exists {
			continue
		}

		for _, pageRequiredBefore := range pagesRequiredBefore {
			pageRequiredBeforeIndex, exists := pageOrder[pageRequiredBefore]
			if !exists {
				continue
			}

			if currPageIndex < pageRequiredBeforeIndex {
				return false
			}
		}
	}

	return true
}

func parseRulesAndUpdates(inputFile io.Reader) (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	updates := [][]int{}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		currLn := scanner.Text()
		if len(currLn) == 0 {
			continue
		}

		lineSplit := strings.Split(currLn, "|")

		if len(lineSplit) == 2 {
			// rule
			num1, _ := strconv.Atoi(lineSplit[0])
			num2, _ := strconv.Atoi(lineSplit[1])

			currNums, exists := rules[num2]
			if !exists {
				rules[num2] = []int{num1}
			}
			rules[num2] = append(currNums, num1)
		} else if len(lineSplit) == 1 {
			// update
			pageSplit := strings.Split(currLn, ",")

			page := make([]int, 0, len(pageSplit))

			for _, pageStr := range pageSplit {
				num, _ := strconv.Atoi(pageStr)
				page = append(page, num)
			}

			updates = append(updates, page)
		}
	}
	return rules, updates
}
