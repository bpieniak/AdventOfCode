package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"os"
)

type position struct {
	x, y int
}

func main() {
	inputFile, err := os.Open("./day10/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(input io.Reader) (result int) {
	topographicMap := getMap(input)

	for _, pos := range getStartingPostions(topographicMap) {
		result += len(trailheadsScore(topographicMap, pos))
	}

	return result
}

func Part2(input io.Reader) (result int) {
	topographicMap := getMap(input)

	for _, pos := range getStartingPostions(topographicMap) {
		result += distinctTrailheadsScore(topographicMap, pos)
	}

	return result
}

func getStartingPostions(topographicMap [][]byte) []position {
	var startingPositions []position
	for y := range len(topographicMap) {
		for x := range len(topographicMap[0]) {
			if topographicMap[y][x] == '0' {
				startingPositions = append(startingPositions, position{x, y})
			}
		}
	}

	return startingPositions
}

func getMap(input io.Reader) [][]byte {
	var topographicMap [][]byte
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		currLn := scanner.Text()

		topographicMap = append(topographicMap, []byte(currLn))
	}
	return topographicMap
}
func trailheadsScore(topographicMap [][]byte, pos position) map[position]struct{} {
	if topographicMap[pos.y][pos.x] == '9' {
		return map[position]struct{}{
			pos: {},
		}
	}

	var trailsSum = make(map[position]struct{})
	for _, dir := range []struct {
		x, y int
	}{
		{1, 0},  // right
		{-1, 0}, // left
		{0, 1},  // down
		{0, -1}, // up
	} {
		newPos := position{pos.x + dir.x, pos.y + dir.y}
		if !(newPos.x >= 0 && newPos.x < len(topographicMap[0]) &&
			newPos.y >= 0 && newPos.y < len(topographicMap)) {
			continue
		}

		if topographicMap[pos.y][pos.x]+1 == topographicMap[newPos.y][newPos.x] {
			maps.Copy(trailsSum, trailheadsScore(topographicMap, newPos))
		}
	}

	return trailsSum
}

func distinctTrailheadsScore(topographicMap [][]byte, pos position) int {
	if topographicMap[pos.y][pos.x] == '9' {
		return 1
	}

	var trailsSum int
	for _, dir := range []struct {
		x, y int
	}{
		{1, 0},  // right
		{-1, 0}, // left
		{0, 1},  // down
		{0, -1}, // up
	} {
		newPos := position{pos.x + dir.x, pos.y + dir.y}
		if !(newPos.x >= 0 && newPos.x < len(topographicMap[0]) &&
			newPos.y >= 0 && newPos.y < len(topographicMap)) {
			continue
		}

		if topographicMap[pos.y][pos.x]+1 == topographicMap[newPos.y][newPos.x] {
			trailsSum += distinctTrailheadsScore(topographicMap, newPos)
		}
	}

	return trailsSum
}
