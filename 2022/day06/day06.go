package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
)

func main() {
	inputScan := helper.GetInputScanner("./day06/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day06/input.txt")
	part2(inputScan)
}

func part1(inputScan *bufio.Scanner) {
	for inputScan.Scan() {
		currLn := inputScan.Text()
		for i := 4; i <= len(currLn); i++ {
			if ok := isUnique(currLn[i-4 : i]); ok {
				fmt.Println(i, currLn[i-4:i])
				return
			}
		}
	}
}

func part2(inputScan *bufio.Scanner) {
	for inputScan.Scan() {
		currLn := inputScan.Text()
		for i := 14; i <= len(currLn); i++ {
			if ok := isUnique(currLn[i-14 : i]); ok {
				fmt.Println(i, currLn[i-14:i])
				return
			}
		}
	}
}

func isUnique(str string) bool {
	var s = make(map[string]struct{})
	for _, c := range str {
		if _, ok := s[string(c)]; ok {
			return false
		}
		s[string(c)] = struct{}{}
	}
	return true
}
