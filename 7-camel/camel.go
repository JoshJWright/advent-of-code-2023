package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"utils"
)

type Hand struct {
	bid         int
	handType    int
	gradedCards [5]int
}

func main() {
	input := utils.LoadString("input.txt")

	lines := strings.Split(input, "\n")

	hands := make([]Hand, 0, len(lines))

	for _, line := range lines {
		data := strings.Split(line, " ")
		bid, _ := strconv.Atoi(data[1])
		hands = append(hands, Hand{
			bid:         bid,
			handType:    getType(data[0]),
			gradedCards: gradeCards([]rune(data[0]))},
		)

	}

	slices.SortFunc(hands, func(a Hand, b Hand) int {
		if a.handType != b.handType {
			return a.handType - b.handType
		}
		for i := 0; i < 5; i++ {
			if a.gradedCards[i] != b.gradedCards[i] {
				return a.gradedCards[i] - b.gradedCards[i]
			}
		}
		return 0
	})

	total := 0
	for i, v := range hands {
		total += (i + 1) * v.bid
	}
	fmt.Println(total)
}

func gradeCards(cards []rune) [5]int {
	allCards := []rune("J23456789TQKA")
	graded := [5]int{}
	for i := range graded {
		graded[i] = slices.Index(allCards, cards[i])
	}
	return graded
}

func getType(hand string) int {
	types := [][5]int{{1, 1, 1, 1, 1}, {2, 1, 1, 1}, {2, 2, 1}, {3, 1, 1}, {3, 2}, {4, 1}, {5}}

	cardMap := make([]int, 0, 5)
	cards := []rune(hand)
	slices.Sort(cards)
	var last rune
	jokers := 0
	for _, c := range cards {
		if c == 'J' {
			jokers++
			continue
		}
		if c == last {
			cardMap[len(cardMap)-1]++
		} else {
			cardMap = append(cardMap, 1)
			last = c
		}
	}

	slices.Sort(cardMap)
	slices.Reverse(cardMap)
	handType := [5]int(cardMap[0:cap(cardMap)])
	handType[0] += jokers
	return slices.Index(types, handType)
}
