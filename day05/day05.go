package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

type mapping struct {
	destinationStart, sourceStart, length int
}

var mappingOrder = []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

func main() {
	inputScan := helper.GetInputScanner("./day05/input.txt")

	inputScan.Scan()
	currLn := inputScan.Text()
	seeds := parseNums(currLn)

	mapNameRegx := regexp.MustCompile(`\w+-\w+-\w+`)

	mappings := make(map[string][]mapping)

	mappingList := []mapping{}
	currMapping := ""
	for inputScan.Scan() {
		currLn := inputScan.Text()

		if currLn == "" {
			sort.Slice(mappingList, func(i, j int) bool {
				return mappingList[i].sourceStart < mappingList[j].sourceStart
			})
			if len(mappingList) > 0 {
				mappings[currMapping] = mappingList
			}

			mappingList = []mapping{}
		}

		mapName := mapNameRegx.FindString(currLn)
		if mapName != "" {
			currMapping = mapName
		}

		nums := parseNums(currLn)
		if len(nums) != 3 {
			continue
		}

		mappingList = append(mappingList, mapping{nums[0], nums[1], nums[2]})
	}

	part1(seeds, mappings)
	part2(seeds, mappings)
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

func part1(seeds []int, mappings map[string][]mapping) {
	minLocation := math.MaxInt
	for _, seed := range seeds {
		location := locationForSeed(seed, mappings)
		minLocation = min(location, minLocation)
	}

	fmt.Println("part1: ", minLocation)
}

func part2(seeds []int, mappings map[string][]mapping) {
	minLocation := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		fmt.Printf("%d%%\n", 100*i/len(seeds))
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			location := locationForSeed(seed, mappings)
			minLocation = min(location, minLocation)
		}
	}

	fmt.Println("part1: ", minLocation)
}

func locationForSeed(seed int, mappings map[string][]mapping) int {
	currValue := seed
	for _, curr := range mappingOrder {
		currMappings := mappings[curr]
		mapping := getMappingValue(currValue, currMappings)
		currValue = mapping
	}

	return currValue
}

func getMappingValue(currValue int, currMappings []mapping) int {
	for _, mapping := range currMappings {
		if currValue < mapping.sourceStart {
			continue
		}

		if mapping.sourceStart <= currValue && currValue < mapping.sourceStart+mapping.length {
			diff := currValue - mapping.sourceStart
			return mapping.destinationStart + diff
		}
	}

	return currValue
}
