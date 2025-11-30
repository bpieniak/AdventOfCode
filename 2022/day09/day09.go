package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type instruction struct {
	dir   string
	steps int
}

type position struct {
	x, y int
}

func (p *position) move(instr instruction) {
	switch instr.dir {
	case "R":
		p.x += 1
	case "L":
		p.x -= 1
	case "U":
		p.y += 1
	case "D":
		p.y -= 1
	}
}

func (p *position) adjustPosition(head *position) {
	diffX := float64(head.x - p.x)
	diffY := float64(head.y - p.y)

	var dir string
	if math.Abs(diffX) > 1 && math.Abs(diffY) == 0 {
		dir = getDirHorizontal(diffX)
		p.move(instruction{dir, 1})
	} else if math.Abs(diffY) > 1 && math.Abs(diffX) == 0 {
		dir = getDirVertical(diffY)
		p.move(instruction{dir, 1})
	} else if (math.Abs(diffX) > 1 && math.Abs(diffY) > 0) || (math.Abs(diffY) > 1 && math.Abs(diffX) > 0) { // diagonal
		dir = getDirHorizontal(diffX)
		p.move(instruction{dir, 1})
		dir = getDirVertical(diffY)
		p.move(instruction{dir, 1})
	}
}

func getDirHorizontal(diffX float64) string {
	if diffX > 0 {
		return "R"
	} else {
		return "L"
	}
}

func getDirVertical(diffY float64) string {
	if diffY > 0 {
		return "U"
	} else {
		return "D"
	}
}

func main() {
	inputScan := helper.GetInputScanner("./day09/input.txt")
	instructions := scanInput(inputScan)

	part1(instructions)
	part2(instructions)
}

func scanInput(inputScan *bufio.Scanner) []instruction {
	var instructions []instruction
	for inputScan.Scan() {
		currLn := inputScan.Text()
		instr := strings.Split(currLn, " ")
		steps, _ := strconv.Atoi(instr[1])

		instructions = append(instructions, instruction{instr[0], steps})
	}

	return instructions
}

func part1(instructions []instruction) {
	headPos := position{}
	tailPos := position{}

	tailPosHistory := make(map[position]struct{})
	for _, instr := range instructions {
		for i := 0; i < instr.steps; i++ {
			headPos.move(instr)
			tailPos.adjustPosition(&headPos)

			tailPosHistory[tailPos] = struct{}{}
		}
	}

	fmt.Println(len(tailPosHistory))
}

func part2(instructions []instruction) {
	headPos := position{}
	tailsPos := make([]position, 9)

	tailPosHistory := make(map[position]struct{})
	for _, instr := range instructions {
		for i := 0; i < instr.steps; i++ {
			headPos.move(instr)
			for i := 0; i < len(tailsPos); i++ {
				if i == 0 {
					tailsPos[0].adjustPosition(&headPos)
				} else {
					tailsPos[i].adjustPosition(&tailsPos[i-1])
				}
			}

			tailPosHistory[tailsPos[8]] = struct{}{}
		}
	}

	fmt.Println(len(tailPosHistory))
}
