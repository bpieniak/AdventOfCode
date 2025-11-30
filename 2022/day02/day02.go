package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	inputScan := helper.GetInputScanner("./day02/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day02/input.txt")
	part2(inputScan)
}

const (
	winScore   = 6
	drawScore  = 3
	looseScore = 0
)

var pickScore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var pick = map[string]string{
	"X": "A", // rock
	"Y": "B", // paper
	"Z": "C", // scissors
}

func part1(inputScan *bufio.Scanner) {
	totalScore := 0
	for inputScan.Scan() {
		currLn := inputScan.Text()
		currLnSplit := strings.Split(currLn, " ")

		opponentPick, myPick := currLnSplit[0], currLnSplit[1]
		score := gameScore(opponentPick, myPick) + pickScore[myPick]
		totalScore += score
	}

	fmt.Println(totalScore)
}

func gameScore(opp string, you string) int {
	you = pick[you]

	if opp == you {
		return drawScore
	} else if opp == "A" && you == "B" {
		return winScore
	} else if opp == "A" && you == "C" {
		return looseScore
	} else if opp == "B" && you == "A" {
		return looseScore
	} else if opp == "B" && you == "C" {
		return winScore
	} else if opp == "C" && you == "A" {
		return winScore
	} else if opp == "C" && you == "B" {
		return looseScore
	}
	return -1
}

var gameScore2 = map[string]int{
	"X": looseScore,
	"Y": drawScore,
	"Z": winScore,
}

func whatToPick(opp string, gameResult string) string {
	if gameResult == "Y" { // draw
		return opp
	} else if gameResult == "X" { // lost
		if opp == "A" {
			return "C"
		} else if opp == "B" {
			return "A"
		} else if opp == "C" {
			return "B"
		}
	} else if gameResult == "Z" { // win
		if opp == "A" {
			return "B"
		} else if opp == "B" {
			return "C"
		} else if opp == "C" {
			return "A"
		}
	}
	return ""
}

var pickScore2 = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func part2(inputScan *bufio.Scanner) {
	totalScore := 0
	for inputScan.Scan() {
		currLn := inputScan.Text()
		currLnSplit := strings.Split(currLn, " ")

		opponentPick, gameResult := currLnSplit[0], currLnSplit[1]
		myPick := whatToPick(opponentPick, gameResult)

		score := gameScore2[gameResult] + pickScore2[myPick]
		totalScore += score
	}

	fmt.Println(totalScore)
}
