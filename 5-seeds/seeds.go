package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type RangedMap struct {
	destStart   int
	sourceStart int
	length      int
}

func main() {
	raw, _ := os.ReadFile("input.txt")
	input := string(raw)
	almanacPattern := regexp.MustCompile(`seeds:(?P<seeds>[\d \n]*)\n+seed-to-soil map:\n(?P<ss>[\d \n]*)\n+soil-to-fertilizer map:\n(?P<sf>[\d \n]*)\n+fertilizer-to-water map:\n(?P<fw>[\d \n]*)\n+water-to-light map:\n(?P<wl>[\d \n]*)\n+light-to-temperature map:\n(?P<lt>[\d \n]*)\n+temperature-to-humidity map:\n(?P<th>[\d \n]*)\n+humidity-to-location map:\n(?P<hl>[\d \n]*)`)
	almanacText := almanacPattern.FindStringSubmatch(input)

	seeds := parseNumberList(almanacText[almanacPattern.SubexpIndex("seeds")])

	minLocations := make([]int, 0, len(seeds)/2)

	for i := 0; i < len(seeds); i += 2 {
		minLocations = append(minLocations, getLowestLocation(seeds[i], seeds[i+1], almanacText, almanacPattern))
		fmt.Println(slices.Min(minLocations))
	}

	fmt.Println(slices.Min(minLocations))
}

func getLowestLocation(startSeed int, numSeeds int, text []string, pattern *regexp.Regexp) int {

	locations := make([]int, 0, numSeeds)
	for i := 0; i < numSeeds; i++ {
		locations = append(locations, startSeed+i)
	}
	transformInput(locations, parseRangedMaps(text[pattern.SubexpIndex("ss")]))
	transformInput(locations, parseRangedMaps(text[pattern.SubexpIndex("sf")]))
	transformInput(locations, parseRangedMaps(text[pattern.SubexpIndex("fw")]))
	transformInput(locations, parseRangedMaps(text[pattern.SubexpIndex("wl")]))
	transformInput(locations, parseRangedMaps(text[pattern.SubexpIndex("lt")]))
	transformInput(locations, parseRangedMaps(text[pattern.SubexpIndex("th")]))
	transformInput(locations, parseRangedMaps(text[pattern.SubexpIndex("hl")]))
	return slices.Min(locations)
}

func transformInput(input []int, mapping []RangedMap) {
	for i := range input {
		for _, rm := range mapping {
			if input[i] >= rm.sourceStart && rm.sourceStart+rm.length > input[i] {
				input[i] = input[i] + rm.destStart - rm.sourceStart
				break
			}
		}
	}
}

func parseRangedMaps(s string) []RangedMap {
	inputs := strings.Split(s, "\n")
	parsed := make([]RangedMap, 0, len(inputs))
	for _, input := range inputs {
		if input != "" {
			parsed = append(parsed, parseRangedMap(input))
		}
	}
	return parsed
}

func parseRangedMap(s string) RangedMap {
	nums := parseNumberList(s)
	if len(nums) != 3 {
		fmt.Println(s, nums)
		panic("WTF")
	}

	return RangedMap{destStart: nums[0], sourceStart: nums[1], length: nums[2]}
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
	return nums
}
