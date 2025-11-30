package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("./day11/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(inputFile io.Reader) (result int) {
	inputMap := parseInput(inputFile)

	for range 25 {
		blink(inputMap)
	}

	for _, count := range inputMap {
		result += count
	}

	return result
}

func Part2(inputFile io.Reader) (result int) {
	inputMap := parseInput(inputFile)

	for range 75 {
		blink(inputMap)
	}

	for _, count := range inputMap {
		result += count
	}

	return result
}

func parseInput(inputFile io.Reader) map[int]int {
	input, _ := io.ReadAll(inputFile)

	var inputInts []int
	inputSplit := strings.Split(string(input), " ")
	for _, str := range inputSplit {
		i, _ := strconv.Atoi(str)
		inputInts = append(inputInts, i)
	}

	inputMap := make(map[int]int, len(input))
	for _, i := range inputInts {
		inputMap[i] = inputMap[i] + 1
	}
	return inputMap
}

func blink(input map[int]int) {
	newNums := make(map[int]int)
	for stone, count := range input {
		delete(input, stone)

		if stone == 0 {
			newNums[1] += count
			continue
		}

		strStone := strconv.Itoa(stone)
		if len(strStone)%2 == 0 {
			split1, split2 := strStone[:len(strStone)/2], strStone[len(strStone)/2:]
			split1Num, _ := strconv.Atoi(split1)
			split2Num, _ := strconv.Atoi(split2)

			newNums[split1Num] += count
			newNums[split2Num] += count
			continue
		}

		newNums[stone*2024] += count
	}

	for k, v := range newNums {
		input[k] += v
	}
}
