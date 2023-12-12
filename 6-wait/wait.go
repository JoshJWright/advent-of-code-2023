package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	time   int
	record int
}

func main() {
	raw, _ := os.ReadFile("input.txt")
	input := string(raw)

	parser := regexp.MustCompile(`Time:\s+(?P<times>.*)\nDistance:\s+(?P<distances>.*)`)
	parsed := parser.FindStringSubmatch(input)
	parsedTimes := parseNumberList(parsed[parser.SubexpIndex("times")])
	parsedDistances := parseNumberList(parsed[parser.SubexpIndex("distances")])

	if len(parsedTimes) != len(parsedDistances) {
		panic("Bad input parsing")
	}

	races := make([]Race, 0, len(parsedTimes))

	for i := range parsedTimes {
		races = append(races, Race{parsedTimes[i], parsedDistances[i]})
	}

	total := 0
	for _, race := range races {
		lower, upper := getOptimalRange(race)
		num := upper - (lower - 1)
		if total == 0 {
			total = num
		} else {
			total *= num
		}
	}
	fmt.Println(total)
}

func getOptimalRange(r Race) (int, int) {

	a := float64(-1)
	b := float64(r.time)
	c := float64(-r.record)

	lower := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a
	upper := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a

	if lower == math.Ceil(lower) {
		lower = lower + 1
	} else {
		lower = math.Ceil(lower)
	}

	if upper == math.Floor(upper) {
		upper = upper - 1
	} else {
		upper = math.Floor(upper)
	}

	return int(lower), int(upper)

}

func parseNumberList(s string) []int {
	spaces := regexp.MustCompile(`\s+`)
	stringNums := spaces.ReplaceAllString(s, "")
	// stringNums := spaces.Split(strings.TrimSpace(s), -1)
	fmt.Sprintln(len(stringNums))
	nums := make([]int, 0, len(stringNums))
	num, _ := strconv.Atoi(stringNums)
	nums = append(nums, num)
	return nums
}
