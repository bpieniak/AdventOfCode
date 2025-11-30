package main

import (
	"fmt"
	"strings"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

type Pipe struct {
	direction1 string
	direction2 string
}

var pipes = map[string]Pipe{
	"|": {"north", "south"},
	"-": {"east", "west"},
	"L": {"north", "east"},
	"J": {"north", "west"},
	"7": {"south", "west"},
	"F": {"south", "east"},
}

type coordinate struct {
	x, y int
}

var directions = map[string]coordinate{
	"north": {0, -1},
	"south": {0, 1},
	"east":  {1, 0},
	"west":  {-1, 0},
}

var opposite = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

const MAX_STEPS = 1000000

func main() {
	input := helper.GetInput("./day10/input.txt")

	inputMat := toMatrix(input)
	start := startingPoint(inputMat)

	// check valid pipes from start
	var validDirections []string
	for dir, coord := range directions {
		if start.y+coord.y < 0 || start.y+coord.y > len(inputMat) || start.x+coord.x < 0 || start.x+coord.x > len(inputMat[0]) {
			continue
		}

		pipe := inputMat[start.y+coord.y][start.x+coord.x]
		if pipe == "." {
			continue
		}

		if pipes[pipe].direction1 == opposite[dir] || pipes[pipe].direction2 == opposite[dir] {
			validDirections = append(validDirections, dir)
		}
	}

	var (
		step          int
		currDirection = validDirections[0]
		currPosition  = start
	)

	for step = 0; step < MAX_STEPS; step++ {
		coord := directions[currDirection]
		currPosition = &coordinate{currPosition.x + coord.x, currPosition.y + coord.y}
		currPipe := inputMat[currPosition.y][currPosition.x]
		if currPipe == "S" {
			break
		}

		if pipes[currPipe].direction1 == opposite[currDirection] {
			currDirection = pipes[currPipe].direction2
		} else {
			currDirection = pipes[currPipe].direction1
		}
	}

	fmt.Println("part1", step/2+1)
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

func startingPoint(input [][]string) *coordinate {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == "S" {
				return &coordinate{j, i}
			}
		}
	}

	return nil
}
