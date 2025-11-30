package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	value   int
	numbers []int
}

func main() {
	inputFile, err := os.Open("./day07/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(inputFile io.Reader) (result int) {
	operators := []string{"+", "*"}

	equations := parseEquations(inputFile)

	return findCalibrationResult(equations, operators, result)
}

func Part2(inputFile io.Reader) (result int) {
	operators := []string{"+", "*", "||"}

	equations := parseEquations(inputFile)

	return findCalibrationResult(equations, operators, result)
}

func findCalibrationResult(equations []equation, operators []string, result int) int {
	for _, equation := range equations {
		operatorsPermutations := generateOperatorPermutations(operators, len(equation.numbers)-1)

		for _, operators := range operatorsPermutations {
			calculatedValue := calculateEquation(equation.numbers, operators)
			if calculatedValue == equation.value {
				result += calculatedValue
				break
			}
		}
	}
	return result
}

func calculateEquation(numbers []int, operators []string) (result int) {
	for i, number := range numbers {
		if i == 0 {
			result = number
			continue
		}

		result = calculate(result, number, operators[i-1])
	}

	return result
}

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "*":
		return num1 * num2
	case "+":
		return num1 + num2
	case "||":
		numConcat := strconv.Itoa(num1) + strconv.Itoa(num2)
		result, _ := strconv.Atoi(numConcat)

		return result
	}

	panic(operator)
}

func parseEquations(inputFile io.Reader) []equation {
	var equations []equation

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		currLn := scanner.Text()

		currLnSplit := strings.Split(currLn, ":")

		value, _ := strconv.Atoi(currLnSplit[0])

		numbers := strings.Split(currLnSplit[1], " ")
		numbersInt := make([]int, 0, len(numbers))

		for _, number := range numbers[1:] {
			numberInt, _ := strconv.Atoi(number)
			numbersInt = append(numbersInt, numberInt)
		}

		equations = append(equations, equation{
			value:   value,
			numbers: numbersInt,
		})
	}

	return equations
}

func generateOperatorPermutations(values []string, length int) [][]string {
	if length == 0 {
		return [][]string{{}}
	}

	subPermutations := generateOperatorPermutations(values, length-1)
	var permutations [][]string

	for _, value := range values {
		for _, perm := range subPermutations {
			newPerm := append([]string{value}, perm...)
			permutations = append(permutations, newPerm)
		}
	}

	return permutations
}
