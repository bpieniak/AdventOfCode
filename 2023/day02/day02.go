package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

func main() {
	inputScan := helper.GetInputScanner("./day02/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day02/input.txt")
	part2(inputScan)
}

type game struct {
	id   int
	bags []Bag
}

type Bag struct {
	blue, green, red int
}

// Game 1: 4 blue, 16 green, 2 red; 5 red, 11 blue, 16 green; 9 green, 11 blue; 10 blue, 6 green, 4 red
func parseGame(str string) game {
	split := strings.Split(str, ":")
	gameIDPart, gameInfoPart := split[0], split[1]

	gameIDSplit := strings.Split(gameIDPart, " ")
	gameID, _ := strconv.Atoi(gameIDSplit[1])

	game := game{
		id: gameID,
	}

	games := strings.Split(gameInfoPart, ";")
	for _, gameStr := range games {
		bagsStr := strings.Split(gameStr, ",")

		bag := Bag{}
		for _, bagStr := range bagsStr {
			bagStr = strings.TrimPrefix(bagStr, " ")

			bagSplit := strings.Split(bagStr, " ")
			count, _ := strconv.Atoi(bagSplit[0])
			color := bagSplit[1]

			switch color {
			case "blue":
				bag.blue = count
			case "red":
				bag.red = count
			case "green":
				bag.green = count
			}
		}

		game.bags = append(game.bags, bag)
	}
	return game
}

func getMaxValues(game game) (maxRed, maxGreen, maxBlue int) {
	for _, bag := range game.bags {
		maxRed = max(maxRed, bag.red)
		maxGreen = max(maxGreen, bag.green)
		maxBlue = max(maxBlue, bag.blue)
	}

	return maxRed, maxGreen, maxBlue
}

const (
	maxRed   = 12
	maxGreen = 13
	maxblue  = 14
)

func part1(inputScan *bufio.Scanner) {
	var idsSum int
	for inputScan.Scan() {
		currLn := inputScan.Text()
		game := parseGame(currLn)
		r, g, b := getMaxValues(game)

		if r > maxRed || g > maxGreen || b > maxblue {
			continue
		}

		idsSum += game.id
	}
	fmt.Println(idsSum)
}

func part2(inputScan *bufio.Scanner) {
	var powersSum int
	for inputScan.Scan() {
		currLn := inputScan.Text()
		game := parseGame(currLn)
		r, g, b := getMaxValues(game)

		powersSum += r * g * b
	}
	fmt.Println(powersSum)
}
