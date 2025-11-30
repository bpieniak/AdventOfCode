package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

type race struct {
	time, distance int
}

func main() {
	input := helper.GetInput("./day06/input.txt")

	var part1races []race
	inputSplit := strings.Split(input, "\n")
	times, distances := parseNums(inputSplit[0]), parseNums(inputSplit[1])
	for i := range times {
		part1races = append(part1races, race{times[i], distances[i]})
	}
	part1(part1races)

	//part2
	time, distance := parseNums(strings.ReplaceAll(inputSplit[0], " ", ""))[0], parseNums(strings.ReplaceAll(inputSplit[1], " ", ""))[0]
	part2 := winRaceWays(race{time, distance})
	fmt.Println(part2)
}

func part1(races []race) {
	waysToBeatRecord := 1
	for _, race := range races {
		waysToBeatRecord *= winRaceWays(race)
	}

	fmt.Println(waysToBeatRecord)
}

func winRaceWays(race race) int {
	var waysToWinRace int
	for i := 1; i < race.time; i++ {
		if calcDistance(i, race.time) > race.distance {
			waysToWinRace++
		}
	}
	return waysToWinRace
}

func calcDistance(holdTime, raceTime int) int {
	return (raceTime - holdTime) * holdTime
}

func parseNums(n string) []int {
	regx := regexp.MustCompile(`\d+`)

	foundNums := regx.FindAllString(n, -1)

	nums := make([]int, 0, len(foundNums))
	for _, num := range foundNums {
		numInt, _ := strconv.Atoi(num)
		nums = append(nums, numInt)
	}

	return nums
}
