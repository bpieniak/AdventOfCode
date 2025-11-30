package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	inputFile, err := os.Open("./day09/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", Part1(inputFile))

	inputFile.Seek(0, io.SeekStart)
	fmt.Println("Part2: ", Part2(inputFile))
}

func Part1(input io.Reader) (result int) {
	diskMap, _ := io.ReadAll(input)

	blocksRepresentation := toBlockRepresentation(diskMap)

	left, right := 0, len(blocksRepresentation)-1
	for left < right {
		if blocksRepresentation[left] != -1 {
			left++
			continue
		}

		if blocksRepresentation[right] == -1 {
			right--
			continue
		}

		swap(blocksRepresentation, left, right)
	}

	result = calculateChecksum(blocksRepresentation)

	return result
}

func calculateChecksum(blocksRepresentation []int) (checksum int) {
	for i, block := range blocksRepresentation {
		if block == -1 {
			continue
		}

		checksum += i * block
	}
	return checksum
}

func toBlockRepresentation(diskMap []byte) []int {
	var blocksRepresentation []int
	var blockID = 0

	for i, block := range diskMap {
		length, _ := strconv.Atoi(string(block))
		representation := -1
		if i%2 == 0 {
			representation = blockID
			blockID++
		}

		for range length {
			blocksRepresentation = append(blocksRepresentation, representation)
		}
	}
	return blocksRepresentation
}

type block struct {
	id       int
	length   int
	wasMoved bool
}

func Part2(input io.Reader) (result int) {
	diskMap, _ := io.ReadAll(input)

	var blocks []block
	var blockID = 0

	for i, cell := range diskMap {
		length, _ := strconv.Atoi(string(cell))
		representation := -1
		if i%2 == 0 {
			representation = blockID
			blockID++
		}

		blocks = append(blocks, block{
			id:     representation,
			length: length,
		})
	}

	left, right := 0, len(blocks)-1
	for right > 0 {
		if right <= left {
			right--
			left = 0
			continue
		}

		if blocks[right].id == -1 || blocks[right].wasMoved {
			right--
			continue
		}

		if !(blocks[left].id == -1 && blocks[left].length >= blocks[right].length) {
			left++
			continue
		}

		sizeDiff := blocks[left].length - blocks[right].length

		blocks[right].wasMoved = true
		blocks[left].length = blocks[right].length

		swap(blocks, left, right)
		if sizeDiff > 0 {
			blocks = slices.Insert(blocks, left+1, block{
				id:     -1,
				length: sizeDiff,
			})
		}
		left = 0
	}

	return calculateChecksum(toInts(blocks))
}

func toInts(blocks []block) []int {
	var representation []int
	for _, block := range blocks {
		for range block.length {
			representation = append(representation, block.id)
		}
	}
	return representation
}

func swap[E any](slice []E, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
