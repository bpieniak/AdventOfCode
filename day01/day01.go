package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

func main() {
	inputScan := helper.GetInputScanner("./day01/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day01/input.txt")
	part2(inputScan)
}

func part1(inputScan *bufio.Scanner) {
	var sum int
	for inputScan.Scan() {
		currLn := inputScan.Text()
		sum += getFirstAndLastCombined(currLn)
	}

	fmt.Println(sum)
}

func part2(inputScan *bufio.Scanner) {
	reg := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)

	var sum int
	for inputScan.Scan() {
		currLn := inputScan.Text()
		// use of this weird replacing func and double replacing to account for
		// situations when two spelled numbers overlap, eg. nineoneight
		currLn = reg.ReplaceAllStringFunc(currLn, spelledToNum)
		currLn = reg.ReplaceAllStringFunc(currLn, spelledToNum)
		sum += getFirstAndLastCombined(currLn)
	}
}

func getFirstAndLastCombined(text string) int {
	var first, last string
	for _, c := range text {
		if c < 48 || c > 57 { // 0-9 in ASCII
			continue
		}

		if first == "" {
			first = string(c)
		}

		last = string(c)
	}

	i, _ := strconv.Atoi(fmt.Sprint(first, last))
	fmt.Println(i)

	return i
}

func spelledToNum(str string) string {
	return spellMap[str]
}

var spellMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "th3e",
	"four":  "fo4ur",
	"five":  "fi5ve",
	"six":   "s6x",
	"seven": "se7en",
	"eight": "ei8ht",
	"nine":  "ni9ne",
}
