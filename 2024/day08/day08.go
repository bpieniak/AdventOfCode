package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type position struct {
	x, y int
}

func main() {
	inputFile, err := os.Open("./day08/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(input io.Reader) (result int) {
	inputMap := parseMap(input)

	antennasMap := findAntennas(inputMap)
	uniqueAntinodes := make(map[position]struct{})

	for _, positions := range antennasMap {
		for i := range positions {
			for j := i + 1; j < len(positions); j++ {
				xVec, yVec := positions[i].x-positions[j].x, positions[i].y-positions[j].y

				antinode1 := position{
					x: positions[i].x + xVec,
					y: positions[i].y + yVec,
				}

				antinode2 := position{
					x: positions[j].x - xVec,
					y: positions[j].y - yVec,
				}

				if withinBorders(antinode1, len(inputMap[0]), len(inputMap)) {
					uniqueAntinodes[antinode1] = struct{}{}
				}

				if withinBorders(antinode2, len(inputMap[0]), len(inputMap)) {
					uniqueAntinodes[antinode2] = struct{}{}
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

func Part2(input io.Reader) (result int) {
	inputMap := parseMap(input)

	antennasMap := findAntennas(inputMap)
	uniqueAntinodes := make(map[position]struct{})

	for _, positions := range antennasMap {
		for i := range positions {
			for j := i + 1; j < len(positions); j++ {
				xVec, yVec := positions[i].x-positions[j].x, positions[i].y-positions[j].y

				antinode1 := positions[i]
				for withinBorders(antinode1, len(inputMap[0]), len(inputMap)) {
					uniqueAntinodes[antinode1] = struct{}{}

					antinode1 = position{
						x: antinode1.x + xVec,
						y: antinode1.y + yVec,
					}
				}

				antinode2 := positions[j]
				for withinBorders(antinode2, len(inputMap[0]), len(inputMap)) {
					uniqueAntinodes[antinode2] = struct{}{}

					antinode2 = position{
						x: antinode2.x - xVec,
						y: antinode2.y - yVec,
					}
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

func withinBorders(pos position, maxX, maxY int) bool {
	return pos.x >= 0 && pos.x < maxX && pos.y >= 0 && pos.y < maxY
}

func findAntennas(inputMap [][]byte) map[byte][]position {
	antennasMap := make(map[byte][]position)

	for y := range len(inputMap) {
		for x := range len(inputMap[0]) {
			field := inputMap[y][x]

			if field >= '0' && field <= '9' ||
				field >= 'A' && field <= 'Z' ||
				field >= 'a' && field <= 'z' {
				antennas, exists := antennasMap[field]
				if !exists {
					antennasMap[field] = []position{{x: x, y: y}}
					continue
				}

				antennasMap[field] = append(antennas, position{x: x, y: y})
			}
		}
	}

	return antennasMap
}

func parseMap(inputFile io.Reader) [][]byte {
	var inputMap [][]byte
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputMap = append(inputMap, []byte(scanner.Text()))
	}
	return inputMap
}
