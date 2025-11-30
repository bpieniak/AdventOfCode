package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type position struct {
	x, y int
}

type direction int

const (
	up direction = iota
	down
	right
	left
)

func main() {
	inputFile, err := os.Open("./day06/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(inputFile io.Reader) (result int) {
	inputMap := parseMap(inputFile)

	positionLog, _ := traceRoute(inputMap)

	// add 1 for starting position
	return len(positionLog) + 1
}

func Part2(inputFile io.Reader) (result int) {
	inputMap := parseMap(inputFile)

	positionLog, _ := traceRoute(inputMap)

	for position := range positionLog {
		if inputMap[position.y][position.x] == '^' {
			continue
		}

		modifiedMap := deepCopy(inputMap)
		modifiedMap[position.y][position.x] = '#'

		_, err := traceRoute(modifiedMap)
		if err != nil {
			result++
		}
	}

	return result
}

func traceRoute(inputMap [][]byte) (map[position]direction, error) {
	positionLog := make(map[position]direction)

	currX, currY := findStartingPosition(inputMap)
	currDir := up
	for {
		nextX, nextY := findNextPosition(currX, currY, currDir)

		if currDir == up && nextY < 0 ||
			currDir == down && nextY >= len(inputMap) ||
			currDir == left && nextX < 0 ||
			currDir == right && nextX >= len(inputMap[0]) {
			break
		}

		if inputMap[nextY][nextX] == '#' {
			currDir = rotate90Right(currDir)
			continue
		}

		loggedDir, exists := positionLog[position{currX, currY}]
		if exists && loggedDir == currDir {
			return nil, errors.New("loop detected")
		}

		if !exists {
			positionLog[position{currX, currY}] = currDir
		}

		currX, currY = nextX, nextY
	}

	return positionLog, nil
}

func parseMap(inputFile io.Reader) [][]byte {
	var inputMap [][]byte
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputMap = append(inputMap, []byte(scanner.Text()))
	}
	return inputMap
}

func findNextPosition(currX, currY int, currDir direction) (nextX, nextY int) {
	nextX, nextY = currX, currY

	switch currDir {
	case up:
		nextY -= 1
	case down:
		nextY += 1
	case right:
		nextX += 1
	case left:
		nextX -= 1
	}

	return nextX, nextY
}

func rotate90Right(currDir direction) (nextDir direction) {
	switch currDir {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}

	return
}

func findStartingPosition(board [][]byte) (x, y int) {
	for y := range len(board) {
		for x := range len(board[0]) {
			if board[y][x] == '^' {
				return x, y
			}
		}
	}

	panic("can't find starting position")
}

func deepCopy(data [][]byte) [][]byte {
	copyData := make([][]byte, len(data))

	for i, inner := range data {
		copyData[i] = make([]byte, len(inner))
		copy(copyData[i], inner)
	}

	return copyData
}
