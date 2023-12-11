package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	winning []int
	actual  []int
}

func main() {
	raw, _ := os.ReadFile("input.txt")
	cards := string(raw)

	lines := strings.Split(cards, "\n")
	totalCards := len(lines)

	var cardMultipliers map[int]int = make(map[int]int, totalCards)

	for i, line := range lines {
		cardMultipliers[i] += 1
		numbers := strings.Split(line[strings.IndexAny(line, ":")+1:], "|")
		card := Card{
			winning: parseNumberList(numbers[0]),
			actual:  parseNumberList(numbers[1]),
		}

		matches := 0
		for _, num := range card.winning {
			if slices.Contains(card.actual, num) {
				matches++
				if i+matches < totalCards {
					cardMultipliers[i+matches] += cardMultipliers[i]
				}
			}
		}

	}

	sum := 0
	for _, count := range cardMultipliers {
		sum += count
	}
	fmt.Println(sum)
}

func calculateScore(n int) int {
	if n > 0 {
		return int(math.Pow(2, float64(n-1)))
	} else {
		return 0
	}
}

func parseNumberList(s string) []int {
	spaces := regexp.MustCompile(`\s+`)
	stringNums := spaces.Split(strings.TrimSpace(s), -1)
	fmt.Sprintln(len(stringNums))
	nums := make([]int, 0, len(stringNums))
	for _, stringNum := range stringNums {
		num, _ := strconv.Atoi(stringNum)
		nums = append(nums, num)
	}
	slices.Sort(nums)
	return slices.Compact(nums)
}
