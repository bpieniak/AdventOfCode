package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bpieniak/AdventOfCode2023/internal/helper"
)

var cards = []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var cardsJoker = []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

type hand struct {
	cards string
	bid   int
}

type handType int

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func main() {
	input := helper.GetInputScanner("./day07/input.txt")

	var hands []hand
	for input.Scan() {
		currLn := input.Text()
		currLnSplit := strings.Split(currLn, " ")

		bid, _ := strconv.Atoi(currLnSplit[1])

		hands = append(hands, hand{currLnSplit[0], bid})
	}

	part1(hands)
	part2(hands)
}

func part1(hands []hand) {
	sort.Slice(hands, func(i, j int) bool {
		return !isWin(hands[i].cards, hands[j].cards, getTypePart1, cards)
	})

	var score int
	for i, hand := range hands {
		score += (i + 1) * hand.bid
	}

	fmt.Println(score)
}

func part2(hands []hand) {
	sort.Slice(hands, func(i, j int) bool {
		return !isWin(hands[i].cards, hands[j].cards, getTypeWithJoker, cardsJoker)
	})

	var score int
	for i, hand := range hands {
		score += (i + 1) * hand.bid
	}

	fmt.Println(score)
}

func isWin(hand1, hand2 string, getType func(string) handType, cardDeck []byte) bool {
	type1, type2 := getType(hand1), getType(hand2)
	if type1 == type2 {
		for i := 0; i < len(hand1); i++ {
			if hand1[i] == hand2[i] {
				continue
			}

			return cardScore(cardDeck, hand1[i]) > cardScore(cardDeck, hand2[i])
		}
	}

	return type1 > type2
}

func cardScore(deck []byte, card byte) int {
	for k, v := range deck {
		if card == v {
			return len(deck) - k
		}
	}
	return -1
}

type occurences struct {
	card  string
	count int
}

func getTypePart1(hand string) handType {
	occMap := make(map[string]int, len(hand))
	for _, card := range hand {
		occMap[string(card)] = occMap[string(card)] + 1
	}

	var occ []occurences
	for card, count := range occMap {
		occ = append(occ, occurences{card, count})
	}

	sort.Slice(occ, func(i, j int) bool {
		return occ[i].count > occ[j].count
	})

	return calcType(occ)
}

func getTypeWithJoker(hand string) handType {
	occMap := make(map[string]int, len(hand))
	for _, card := range hand {
		occMap[string(card)] = occMap[string(card)] + 1
	}

	jokers, exists := occMap["J"]
	delete(occMap, "J") // delete to not being count in occurances as it will be replaced

	var occ []occurences
	for card, count := range occMap {
		occ = append(occ, occurences{card, count})
	}

	sort.Slice(occ, func(i, j int) bool {
		return occ[i].count > occ[j].count
	})

	if exists {
		if jokers == 5 {
			return FiveOfAKind
		}

		// replace joker with most common card
		for i := 0; i < len(occ); i++ {
			if occ[i].card != "J" {
				occ[i].count += jokers
				break
			}
		}
	}

	return calcType(occ)
}

func calcType(occ []occurences) handType {
	if occ[0].count == 5 {
		return FiveOfAKind
	} else if occ[0].count == 4 {
		return FourOfAKind
	} else if occ[0].count == 3 && occ[1].count == 2 {
		return FullHouse
	} else if occ[0].count == 3 {
		return ThreeOfAKind
	} else if occ[0].count == 2 && occ[1].count == 2 {
		return TwoPair
	} else if occ[0].count == 2 {
		return OnePair
	}

	return HighCard
}
