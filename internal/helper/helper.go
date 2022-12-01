package helper

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetInput(path string) string {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body)
}

func GetInputScanner(path string) *bufio.Scanner {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
