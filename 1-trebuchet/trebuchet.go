package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var raw []byte
	raw, _ = os.ReadFile("sample.txt")

	var text string = string(raw)

	lines := strings.Split(text, "\n")

	fmt.Println(lines)

	total := 0
	for _, line := range lines {
		for i := range line {
			num := parseNum(line[0 : i+1])
			if num >= 0 {
				total += 10 * num
				break
			}
		}
		last := 0
		for i := range line {
			num := parseNum(line[0 : i+1])
			if num >= 0 {
				last = num
			}
		}
		total += last
	}

	fmt.Printf("Sum: %d\n", total)
}

func parseNum(buffer string) int {
	if len(buffer) >= 1 {
		num := buffer[len(buffer)-1] - '0'
		if num < 10 { // We don't need to check above 0 as it will loop around
			return int(num)
		}
	}
	if len(buffer) >= 3 {
		num := strings.ToLower(string((buffer[len(buffer)-3:])))
		switch num {
		case "one":
			return 1
		case "two":
			return 2
		case "six":
			return 6
		}
	}
	if len(buffer) >= 4 {
		num := strings.ToLower(string((buffer[len(buffer)-4:])))
		switch num {
		case "zero":
			return 0
		case "four":
			return 4
		case "five":
			return 5
		case "nine":
			return 9
		}
	}
	if len(buffer) >= 5 {
		num := strings.ToLower(string((buffer[len(buffer)-5:])))
		switch num {
		case "three":
			return 3
		case "seven":
			return 7
		case "eight":
			return 8
		}
	}
	return -1
}
