package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputScan := helper.GetInputScanner("./day10/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day10/input.txt")
	part2(inputScan)
}

var instructionCost = map[string]int{
	"addx": 2,
	"noop": 1,
}

var cyclesToCheck = []int{20, 60, 100, 140, 180, 220}

func part1(inputScan *bufio.Scanner) {

	checkOnCycle := 20

	register := 1
	cycle := 0
	signalStrengths := 0
	for inputScan.Scan() {
		currLn := inputScan.Text()
		instruction := strings.Split(currLn, " ")

		for i := 0; i < instructionCost[instruction[0]]; i++ {
			cycle += 1
			if cycle == checkOnCycle {
				signalStrengths += cycle * register
				checkOnCycle += 40
			}
		}

		if instruction[0] == "addx" {
			val, _ := strconv.Atoi(instruction[1])
			register += val
		}
	}

	fmt.Println(signalStrengths)
}

func part2(inputScan *bufio.Scanner) {
	register := 1
	cycle := 0
	for inputScan.Scan() {
		currLn := inputScan.Text()
		instruction := strings.Split(currLn, " ")

		for i := 0; i < instructionCost[instruction[0]]; i++ {
			rowCycle := cycle % 40
			if rowCycle == register-1 || rowCycle == register || rowCycle == register+1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			cycle += 1
			if cycle%40 == 0 {
				fmt.Println()
			}
		}

		if instruction[0] == "addx" {
			val, _ := strconv.Atoi(instruction[1])
			register += val
		}
	}
}
