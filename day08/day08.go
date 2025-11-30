package main

import (
	"fmt"
	"regexp"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

type node struct {
	l, r string
}

func main() {
	input := helper.GetInputScanner("./day08/input.txt")

	input.Scan()
	instructions := input.Text()

	nodeMap := make(map[string]node)
	for input.Scan() {
		currLn := input.Text()
		if currLn == "" {
			continue
		}
		nodes := parseNodes(currLn)

		nodeMap[nodes[0]] = node{nodes[1], nodes[2]}
	}

	fmt.Println("part1", findSteps(instructions, nodeMap, "AAA", func(s string) bool {
		return s == "ZZZ"
	}))
	part2(instructions, nodeMap)
}

func part2(instructions string, nodes map[string]node) {
	var endingWithA []string
	for node := range nodes {
		if node[2] == 'A' {
			endingWithA = append(endingWithA, node)
		}
	}

	var steps []int
	for _, starting := range endingWithA {
		steps = append(steps, findSteps(instructions, nodes, starting, func(s string) bool {
			return s[2] == 'Z'
		}))
	}

	fmt.Println("part2", LCM(steps...))
}

func findSteps(instructions string, nodes map[string]node, startingNode string, isEnding func(string) bool) int {
	currNode := startingNode

	var steps int
	for steps = 0; ; steps++ {
		currInstruction := instructions[steps%len(instructions)]
		if currInstruction == 'L' {
			currNode = nodes[currNode].l
		} else if currInstruction == 'R' {
			currNode = nodes[currNode].r
		}

		if isEnding(currNode) {
			break
		}
	}
	return steps + 1
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(integers ...int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func parseNodes(n string) []string {
	regx := regexp.MustCompile(`\w{3}`)

	return regx.FindAllString(n, -1)
}
