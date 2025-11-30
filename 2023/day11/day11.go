package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

type coordinate struct {
	x, y int
}

func main() {
	input := helper.GetInput("./day11/input.txt")
	inputMat := toMatrix(input)
	emptyRows := getEmptyRows(inputMat)
	emptyCols := getEmptyColumns(inputMat)

	coords := findGalaxies(inputMat)

	var part1Sum int
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			part1Sum += calcDistance2(coords[i], coords[j], emptyRows, emptyCols, 2)
		}
	}

	var part2Sum int
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			part2Sum += calcDistance2(coords[i], coords[j], emptyRows, emptyCols, 1_000_000)
		}
	}

	fmt.Println("part1:", part1Sum)
	fmt.Println("part2:", part2Sum)
}

func getEmptyRows(input [][]string) []int {
	var emptyRows []int

	for i, row := range input {
		empty := true
		for _, char := range row {
			if char != "." {
				empty = false
				break
			}
		}

		if empty {
			emptyRows = append(emptyRows, i)
		}
	}

	return emptyRows
}

func getEmptyColumns(input [][]string) []int {
	var emptyCols []int

	for col := 0; col < len(input[0]); col++ {
		empty := true
		for _, row := range input {
			if row[col] != "." {
				empty = false
				break
			}
		}

		if empty {
			emptyCols = append(emptyCols, col)
		}
	}

	return emptyCols
}

func findGalaxies(input [][]string) []coordinate {
	coords := []coordinate{}
	for i := range input {
		for j := range input[i] {
			if input[i][j] == "#" {
				coords = append(coords, coordinate{j, i})
			}
		}
	}

	return coords
}

func calcDistance2(c1, c2 coordinate, emptyRows, emptyCols []int, emptyDistance int) int {
	inbetweenEmptyRows := getInbetweenCount(emptyRows, c1.y, c2.y)
	inbetweenEmptyCols := getInbetweenCount(emptyCols, c1.x, c2.x)

	xDistance := int(math.Abs(float64(c2.x-c1.x))) - inbetweenEmptyCols + inbetweenEmptyCols*emptyDistance
	yDistance := int(math.Abs(float64(c2.y-c1.y))) - inbetweenEmptyRows + inbetweenEmptyRows*emptyDistance

	return xDistance + yDistance
}

func getInbetweenCount(s []int, v1, v2 int) int {
	var (
		minV = min(v1, v2)
		maxV = max(v1, v2)
	)

	var inbetween int
	for _, v := range s {
		if v >= maxV {
			break
		}

		if v > minV {
			inbetween++
		}
	}

	return inbetween
}

func toMatrix(input string) [][]string {
	var mat [][]string
	for _, line := range strings.Split(input, "\n") {
		var lineSlice []string
		for _, c := range line {
			lineSlice = append(lineSlice, string(c))
		}

		mat = append(mat, lineSlice)
	}

	return mat
}
