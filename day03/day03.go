package main

import (
	"fmt"
	"slices"
	"strconv"
	"unicode"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

var symbols = []rune{'-', '@', '*', '/', '&', '#', '%', '+', '=', '$'}

type numberCordinates struct {
	line, start, end int
}

func main() {
	inputScan := helper.GetInputScanner("./day03/input.txt")
	var input []string
	for inputScan.Scan() {
		input = append(input, inputScan.Text())
	}

	var gearRationsSum int
	foundNumbersSet := make(map[numberCordinates]struct{})
	for y, line := range input {
		for x, c := range line {
			if !isSymbol(c) {
				continue
			}

			// part 1
			foundNumbers := findAdjecentNumbers(input, x, y)
			for num := range foundNumbers {
				foundNumbersSet[num] = struct{}{}
			}

			// part 2
			if c == '*' && len(foundNumbers) == 2 {
				gearRation := 1
				for num := range foundNumbers {
					i, _ := strconv.Atoi(input[num.line][num.start : num.end+1])
					gearRation *= i
				}
				gearRationsSum += gearRation
			}
		}
	}

	var sum int
	for num := range foundNumbersSet {
		i, _ := strconv.Atoi(input[num.line][num.start : num.end+1])
		sum += i
	}

	fmt.Println("part1", sum)
	fmt.Println("part2", gearRationsSum)
}

func isSymbol(r rune) bool {
	return slices.Contains(symbols, r)
}

func findAdjecentNumbers(input []string, col, row int) map[numberCordinates]struct{} {
	xLen, yLen := len(input[0]), len(input)

	cords := make(map[numberCordinates]struct{})
	for _, offset := range []struct{ x, y int }{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{-1, 1},
		{-1, -1},
		{1, 1},
		{1, -1},
	} {
		checkX, checkY := col+offset.x, row+offset.y
		if checkX < 0 || checkX > xLen-1 || checkY < 0 || checkY > yLen-1 {
			continue
		}

		if unicode.IsNumber(rune(input[checkY][checkX])) {
			cords[findFullNumber(input[checkY], checkX, checkY)] = struct{}{}
		}
	}

	return cords
}

func findFullNumber(line string, x, y int) numberCordinates {
	var start int
	for start = x; start > 0 && unicode.IsNumber(rune(line[start-1])); start-- {
	}

	var end int
	for end = x; end < len(line)-1 && unicode.IsNumber(rune(line[end+1])); end++ {
	}

	return numberCordinates{
		line:  y,
		start: start,
		end:   end,
	}
}
