package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type moveInstruction struct {
	quantityToMove int
	startingStack  int
	endStack       int
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) getLast() string {
	l := len(s)
	return s[l-1]
}

func main() {
	inputScan := helper.GetInputScanner("./day05/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day05/input.txt")
	part2(inputScan)
}

func part1(inputScan *bufio.Scanner) {
	// get starting stacks
	var startingStacks []string
	for inputScan.Scan() {
		currLn := inputScan.Text()

		if currLn == "" {
			break
		}

		startingStacks = append(startingStacks, currLn)
	}
	stacks := toStacks(startingStacks)

	// get move instuctions
	var instuctions []moveInstruction
	for inputScan.Scan() {
		currLn := inputScan.Text()

		instuction := toMoveInstruction(currLn)
		instuctions = append(instuctions, instuction)
	}

	stacks = moveContainers(stacks, instuctions)
	fmt.Println(getTopOfStacks(stacks))
}

func part2(inputScan *bufio.Scanner) {
	// get starting stacks
	var startingStacks []string
	for inputScan.Scan() {
		currLn := inputScan.Text()

		if currLn == "" {
			break
		}

		startingStacks = append(startingStacks, currLn)
	}
	stacks := toStacks(startingStacks)

	// get move instuctions
	var instuctions []moveInstruction
	for inputScan.Scan() {
		currLn := inputScan.Text()

		instuction := toMoveInstruction(currLn)
		instuctions = append(instuctions, instuction)
	}

	stacks = moveContainers2(stacks, instuctions)
	fmt.Println(getTopOfStacks(stacks))
}

var crateIdentifierPlaces = []int{1, 5, 9, 13, 17, 21, 25, 29, 33}

func toStacks(instuctions []string) map[int]stack {
	var stacks = make(map[int]stack, 9)

	instuctionsRev := reverse(instuctions)
	for _, stackInstruction := range instuctionsRev[1:] {
		for i, place := range crateIdentifierPlaces {
			crateIdentifier := string(stackInstruction[place])
			if crateIdentifier == " " {
				continue
			}
			stacks[i+1] = stacks[i+1].Push(string(stackInstruction[place]))
		}
	}
	return stacks
}

func reverse(s []string) []string {
	var result []string
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func toMoveInstruction(instruction string) moveInstruction {
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(instruction, -1)

	n0, _ := strconv.Atoi(nums[0])
	n1, _ := strconv.Atoi(nums[1])
	n2, _ := strconv.Atoi(nums[2])
	return moveInstruction{n0, n1, n2}
}

func moveContainers(stack map[int]stack, instructions []moveInstruction) map[int]stack {
	var value string
	for _, instr := range instructions {
		for i := 0; i < instr.quantityToMove; i++ {
			stack[instr.startingStack], value = stack[instr.startingStack].Pop()
			stack[instr.endStack] = stack[instr.endStack].Push(value)
		}
	}
	return stack
}

func moveContainers2(stack map[int]stack, instructions []moveInstruction) map[int]stack {
	var value string
	for _, instr := range instructions {
		var movedContainers []string
		for i := 0; i < instr.quantityToMove; i++ {
			stack[instr.startingStack], value = stack[instr.startingStack].Pop()
			movedContainers = append(movedContainers, value)
		}
		for i := len(movedContainers) - 1; i >= 0; i-- {
			stack[instr.endStack] = stack[instr.endStack].Push(movedContainers[i])
		}
	}
	return stack
}

func getTopOfStacks(stacks map[int]stack) []string {
	var tops []string
	for i := 0; i < len(stacks)-1; i++ {
		tops = append(tops, stacks[i+1].getLast())
	}
	return tops
}
