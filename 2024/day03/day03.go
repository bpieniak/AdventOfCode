package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("./day03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(input io.Reader) (result int) {
	inputBytes, _ := io.ReadAll(input)

	result = doMultiplications(inputBytes)

	return result
}

func Part2(input io.Reader) (result int) {
	inputBytes, _ := io.ReadAll(input)

	result = doMultiplications2(inputBytes)

	return result
}

func doMultiplications(inputBytes []byte) (result int) {
	reMultiplication := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	reDigit := regexp.MustCompile(`\d{1,3}`)

	multiplications := reMultiplication.FindAll(inputBytes, -1)

	for _, mul := range multiplications {
		digits := reDigit.FindAll(mul, -1)

		num1, _ := strconv.Atoi(string(digits[0]))
		num2, _ := strconv.Atoi(string(digits[1]))

		result += num1 * num2
	}
	return result
}

func doMultiplications2(inputBytes []byte) (result int) {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	reDigit := regexp.MustCompile(`\d{1,3}`)

	matches := re.FindAll(inputBytes, -1)

	var doFlag = true
	for _, match := range matches {
		if bytes.Equal(match, []byte("do()")) {
			doFlag = true
		} else if bytes.Equal(match, []byte("don't()")) {
			doFlag = false
		} else {
			if !doFlag {
				continue
			}

			digits := reDigit.FindAll(match, -1)

			num1, _ := strconv.Atoi(string(digits[0]))
			num2, _ := strconv.Atoi(string(digits[1]))

			result += num1 * num2
		}
	}
	return result
}
