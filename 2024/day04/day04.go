package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("./day04/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(inputFile io.Reader) (result int) {
	input := readInput(inputFile)

	for y := range len(input) {
		for x := range len(input[0]) {
			for _, dir := range []struct {
				horizontal int
				vertical   int
			}{
				{1, 0},   // horizontal down
				{-1, 0},  // horizontal up
				{0, 1},   // vertical right
				{0, -1},  // vertical left
				{1, 1},   // diagonal down right
				{1, -1},  // diagonal down left
				{-1, 1},  // diagonal up right
				{-1, -1}, // diagonal up left
			} {
				//check boundaries
				if !(x+3*dir.vertical >= 0 &&
					x+3*dir.vertical < len(input[0]) &&
					y+3*dir.horizontal >= 0 &&
					y+3*dir.horizontal < len(input)) {
					continue
				}

				if input[y][x] == 'X' &&
					input[y+1*dir.horizontal][x+1*dir.vertical] == 'M' &&
					input[y+2*dir.horizontal][x+2*dir.vertical] == 'A' &&
					input[y+3*dir.horizontal][x+3*dir.vertical] == 'S' {
					result++
				}
			}
		}

	}

	return result
}

func Part2(inputFile io.Reader) (result int) {
	input := readInput(inputFile)

	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[0])-1; x++ {
			if input[y][x] != 'A' {
				continue
			}

			left_up := input[y-1][x-1]
			left_down := input[y+1][x-1]
			right_up := input[y-1][x+1]
			right_down := input[y+1][x+1]

			if (left_up == 'M' && right_down == 'S' || left_up == 'S' && right_down == 'M') &&
				(right_up == 'M' && left_down == 'S' || right_up == 'S' && left_down == 'M') {
				result++
			}
		}
	}

	return result
}

func readInput(inputFile io.Reader) []string {
	var input []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		b := scanner.Text()
		input = append(input, b)
	}
	return input
}
