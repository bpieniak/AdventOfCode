package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
)

func main() {
	inputScan := helper.GetInputScanner("./day03/input.txt")
	part1(inputScan)

	inputScan = helper.GetInputScanner("./day03/input.txt")
	part2(inputScan)
}

func part1(inputScan *bufio.Scanner) {
	prioritieSum := 0
	for inputScan.Scan() {
		currLn := inputScan.Text()

		midElem := len(currLn) / 2
		items1, items2 := currLn[:midElem], currLn[midElem:]

		elem := findCommonElem(items1, items2)
		prioritie := getPriority(elem)

		prioritieSum += prioritie
	}
	fmt.Println(prioritieSum)
}

func findCommonElem(items1 string, items2 string) string {
	if len(items1) != len(items2) {
		panic("")
	}

	var elemMap = make(map[string]struct{}, len(items1))

	for i := 0; i < len(items1); i++ {
		elemMap[string(items1[i])] = struct{}{}
	}

	for i := 0; i < len(items2); i++ {
		item := string(items2[i])
		if _, ok := elemMap[item]; ok {
			return item
		}
	}

	return ""
}

func getPriority(char string) int {
	asciiVal := int(char[0])
	if asciiVal >= 96 {
		return int(char[0]) - 96
	} else {
		return int(char[0]) - 64 + 26
	}
}

func part2(inputScan *bufio.Scanner) {
	prioritieSum := 0
	for inputScan.Scan() {
		Ln1 := inputScan.Text()
		inputScan.Scan()
		Ln2 := inputScan.Text()
		inputScan.Scan()
		Ln3 := inputScan.Text()

		elem := findCommonElem3(Ln1, Ln2, Ln3)
		prioritie := getPriority(elem)

		prioritieSum += prioritie
	}
	fmt.Println(prioritieSum)
}

func findCommonElem3(items1 string, items2 string, items3 string) string {

	var elemMap1 = make(map[string]struct{}, len(items1))
	var elemMap2 = make(map[string]struct{}, len(items1))

	for i := 0; i < len(items1); i++ {
		elemMap1[string(items1[i])] = struct{}{}
	}

	for i := 0; i < len(items2); i++ {
		elemMap2[string(items2[i])] = struct{}{}
	}

	for i := 0; i < len(items3); i++ {
		item := string(items3[i])
		_, ok1 := elemMap1[item]
		_, ok2 := elemMap2[item]
		if ok1 && ok2 {
			return item
		}
	}

	return ""
}
