package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

func main() {
	inputScan := helper.GetInputScanner("./day04/input.txt")

	// [card_num]matching_len
	cards := make(map[int]int)

	for inputScan.Scan() {
		currLn := inputScan.Text()

		lineSplit := strings.Split(currLn, ":")
		cardNum := parseNums(lineSplit[0])[0]
		cardInfo := lineSplit[1]

		cardInfoSplit := strings.Split(cardInfo, "|")
		winningNumbers, gotNumbers := parseNums(cardInfoSplit[0]), parseNums(cardInfoSplit[1])

		cards[cardNum] = len(getMatchingNumbers(winningNumbers, gotNumbers))

	}

	part1(cards)
	part2(cards)
}

func parseNums(n string) []int {
	regx := regexp.MustCompile(`\d{1,3}`)

	foundNums := regx.FindAllString(n, -1)

	nums := make([]int, 0, len(foundNums))
	for _, num := range foundNums {
		numInt, _ := strconv.Atoi(num)
		nums = append(nums, numInt)
	}

	return nums
}

func getMatchingNumbers(winningNumbers, gotNumbers []int) []int {
	winningSet := make(map[int]struct{})
	for _, num := range winningNumbers {
		winningSet[num] = struct{}{}
	}

	var winningNums []int
	for _, num := range gotNumbers {
		if _, exists := winningSet[num]; exists {
			winningNums = append(winningNums, num)
		}
	}

	return winningNums
}

func part1(cards map[int]int) {
	var sum int
	for _, matching := range cards {
		sum += getScore(matching)
	}
	fmt.Println(sum)
}

func part2(cards map[int]int) {
	gotCards := 0
	for cardNum, matching := range cards {
		gotCards += getCards(cards, getRange(cardNum+1, matching))
	}
	fmt.Println(gotCards)
}

func getCards(cards map[int]int, won []int) int {
	gotCards := 1
	for _, cardNum := range won {
		gotCards += getCards(cards, getRange(cardNum+1, cards[cardNum]))
	}
	return gotCards
}

func getRange(start, len int) []int {
	numRange := make([]int, 0, len)
	for i := start; i < start+len; i++ {
		numRange = append(numRange, i)
	}
	return numRange
}

func getScore(matches int) int {
	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches-1)))
}
