package main

import (
	"aoc/internal/helper"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type fifo []uint64

func (f fifo) add(v uint64) fifo {
	return append(f, v)
}

func (f fifo) get() (fifo, uint64) {
	return f[1:], f[0]
}

func (f fifo) isEmpty() bool {
	return len(f) == 0
}

type operation struct {
	op  string
	val uint64
}

type monkey struct {
	items      fifo
	op         operation
	testDiv    uint64
	trueThrow  int
	falseThrow int
}

func main() {
	inputScan := helper.GetInputScanner("./day11/input.txt")
	monkeys := scanMonkeys(inputScan)

	// part1(monkeys)
	part2(monkeys)
}

func scanMonkeys(inputScan *bufio.Scanner) []monkey {
	var input []string
	for inputScan.Scan() {
		input = append(input, inputScan.Text())
	}

	var monkeys []monkey
	for i := 0; i < len(input); i += 7 {
		monkeys = append(monkeys, parseMonkey(input[i:i+7]))
	}
	return monkeys
}

func parseMonkey(desc []string) monkey {
	var startingItems fifo
	for _, itm := range strings.Split(desc[1][18:], ", ") {
		i, _ := strconv.Atoi(itm)
		startingItems = startingItems.add(uint64(i))
	}

	oper := desc[2][23:24]
	valStr := desc[2][25:]
	val, _ := strconv.Atoi(valStr)
	op := operation{oper, uint64(val)}

	testStr := desc[3][21:]
	test, _ := strconv.Atoi(testStr)

	trueStr := desc[4][29:]
	trueV, _ := strconv.Atoi(trueStr)

	falseStr := desc[5][30:]
	falseV, _ := strconv.Atoi(falseStr)
	m := monkey{startingItems, op, uint64(test), trueV, falseV}
	return m

}

func part1(monkeys []monkey) {
	var timesInspected = make(map[int]int, len(monkeys))

	roundsToDo := 20
	for i := 0; i < roundsToDo; i++ {
		doRound(monkeys, timesInspected, func(v uint64) uint64 { return v / 3 })
	}
	monkeyBusiness(timesInspected)
}

func part2(monkeys []monkey) {
	dividersLCM := uint64(1)
	for _, monkey := range monkeys {
		dividersLCM = leastCommonMultiple(dividersLCM, monkey.testDiv)
	}

	roundsToDo := 10000
	var timesInspected = make(map[int]int, len(monkeys))
	for i := 0; i < roundsToDo; i++ {
		doRound(monkeys, timesInspected, func(v uint64) uint64 { return v % dividersLCM })
	}
	monkeyBusiness(timesInspected)
}

func leastCommonMultiple(a, b uint64) uint64 {
	return a * b / greatestCommonDivisor(a, b)
}

func greatestCommonDivisor(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func doRound(monkeys []monkey, timesInspected map[int]int, relive func(v uint64) uint64) {
	for i, _ := range monkeys {
		timesInspected[i] += len(monkeys[i].items)

		for !monkeys[i].items.isEmpty() {
			rest, item := monkeys[i].items.get()
			monkeys[i].items = rest

			item = inspect(item, monkeys[i].op)
			item = relive(item)

			if item%monkeys[i].testDiv == 0 {
				monkeys[monkeys[i].trueThrow].items = monkeys[monkeys[i].trueThrow].items.add(item)
			} else {
				monkeys[monkeys[i].falseThrow].items = monkeys[monkeys[i].falseThrow].items.add(item)
			}
		}
	}
	// fmt.Println(monkeys)
}

func inspect(item uint64, op operation) uint64 {
	val := op.val
	if val == 0 {
		val = item
	}

	if op.op == "*" {
		return item * val
	} else if op.op == "+" {
		return item + val
	}
	fmt.Println("XD")
	return 0
}

func monkeyBusiness(m map[int]int) {
	var v []int

	for _, value := range m {
		v = append(v, value)
	}
	fmt.Println(m)
	sort.Sort(sort.Reverse(sort.IntSlice(v)))
	fmt.Println(v[0] * v[1])
}
