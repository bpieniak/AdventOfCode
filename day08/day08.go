package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	inputScan := helper.GetInputScanner("./day08/input.txt")
	trees := scanTrees(inputScan)
	part1(trees)
	part2(trees)
}

type Trees struct {
	HeightMap    [][]int
	sizeX, sizeY int
}

func scanTrees(inputScan *bufio.Scanner) Trees {
	mat := make([][]int, 99)
	j := 0
	for inputScan.Scan() {
		for _, char := range inputScan.Text() {
			val, _ := strconv.Atoi(string(char))
			mat[j] = append(mat[j], val)
		}
		j += 1
	}
	sizeX, sizeY := len(mat), len(mat[0])

	return Trees{mat, sizeX, sizeY}
}

func part1(trees Trees) {
	visibleTrees := 2*trees.sizeX + 2*trees.sizeY - 4 // edge
	for i := 1; i < trees.sizeX-1; i++ {
		for j := 1; j < trees.sizeY-1; j++ {
			if trees.isVisible(i, j) {
				visibleTrees += 1
			}
		}
	}
	fmt.Println(visibleTrees)
}

func (t *Trees) isVisible(x int, y int) bool {
	return t.isVisibleLeft(x, y) || t.isVisibleRight(x, y) ||
		t.isVisibleTop(x, y) || t.isVisibleBottom(x, y)
}

func (t *Trees) isVisibleLeft(x int, y int) bool {
	treeHeight := t.HeightMap[x][y]

	for i := 0; i < x; i++ {
		if t.HeightMap[i][y] >= treeHeight {
			return false
		}
	}
	return true
}

func (t *Trees) isVisibleRight(x int, y int) bool {
	treeHeight := t.HeightMap[x][y]

	for i := x + 1; i < t.sizeX; i++ {
		if t.HeightMap[i][y] >= treeHeight {
			return false
		}
	}
	return true
}

func (t *Trees) isVisibleTop(x int, y int) bool {
	treeHeight := t.HeightMap[x][y]

	for j := 0; j < y; j++ {
		if t.HeightMap[x][j] >= treeHeight {
			return false
		}
	}
	return true
}

func (t *Trees) isVisibleBottom(x int, y int) bool {
	treeHeight := t.HeightMap[x][y]

	for j := y + 1; j < t.sizeY; j++ {
		if t.HeightMap[x][j] >= treeHeight {
			return false
		}
	}
	return true
}

func part2(trees Trees) {
	maxScore := 0
	for i := 0; i < trees.sizeX; i++ {
		for j := 0; j < trees.sizeY; j++ {
			score := trees.scenicScore(i, j)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println(maxScore)
	// fmt.Println(trees.scenicScore(98, 98))
}

func (t *Trees) scenicScore(x int, y int) int {
	return t.scenicScoreLeft(x, y) * t.scenicScoreRight(x, y) * t.scenicScoreTop(x, y) * t.scenicScoreBottom(x, y)
}

func (t *Trees) scenicScoreLeft(x int, y int) int {
	treeHeight := t.HeightMap[x][y]
	scenicScore := 0
	for i := x - 1; i >= 0; i-- {
		scenicScore += 1
		if t.HeightMap[i][y] >= treeHeight {
			break
		}
	}
	return scenicScore
}

func (t *Trees) scenicScoreRight(x int, y int) int {
	treeHeight := t.HeightMap[x][y]
	scenicScore := 0
	for i := x + 1; i < t.sizeX; i++ {
		scenicScore += 1
		if t.HeightMap[i][y] >= treeHeight {
			break
		}
	}
	return scenicScore
}

func (t *Trees) scenicScoreTop(x int, y int) int {
	treeHeight := t.HeightMap[x][y]
	scenicScore := 0
	for j := y - 1; j >= 0; j-- {
		scenicScore += 1
		if t.HeightMap[x][j] >= treeHeight {
			break
		}
	}
	return scenicScore
}

func (t *Trees) scenicScoreBottom(x int, y int) int {
	treeHeight := t.HeightMap[x][y]
	scenicScore := 0
	for j := y + 1; j < t.sizeY; j++ {
		scenicScore += 1
		if t.HeightMap[x][j] >= treeHeight {
			break
		}
	}
	return scenicScore
}
