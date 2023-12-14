package main

import (
	"fmt"
	"slices"
	"strings"
	"utils"
)

func main() {
	input := utils.LoadString("input.txt")
	lines := strings.Split(input, "\n")
	steps := lines[0]

	nodes := make(map[string][2]string)

	for _, line := range lines[2:] {
		nodes[line[0:3]] = [2]string{line[7:10], line[12:15]}
	}

	starts := make([]string, 0)

	for k := range nodes {
		if k[2] == 'A' {
			starts = append(starts, k)
		}
	}

	lengths := make([]int, 0, len(starts))

	for _, v := range starts {
		lengths = append(lengths, length(v, steps, nodes))
	}

	fmt.Println(computeLcm(lengths))
}

func computeLcm(lengths []int) int {
	lcm := slices.Clone(lengths)

	iter := 0

	for {
		if slices.IsSorted(lcm) && lcm[0] == lcm[len(lcm)-1] {
			return lcm[0]
		}
		lowestIndex := 0
		for i := 1; i < len(lcm); i++ {
			if lcm[i] < lcm[lowestIndex] {
				lowestIndex = i
			}
		}

		lcm[lowestIndex] += lengths[lowestIndex]
		iter++
		if iter%1000000 == 0 {
			fmt.Println(lcm)
		}

	}
}

func length(start string, steps string, nodes map[string][2]string) int {
	node := start
	stepCounter := 0
	totalSteps := 0
	for node[2] != 'Z' {
		dir := 0
		if steps[stepCounter] == 'R' {
			dir = 1
		}
		node = nodes[node][dir]

		stepCounter++
		totalSteps++
		if stepCounter >= len(steps) {
			stepCounter = 0
		}
	}
	return totalSteps
}
