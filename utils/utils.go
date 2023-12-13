package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LoadString(fileName string) string {
	raw, _ := os.ReadFile("input.txt")
	return string(raw)
}

func ParseNumberList(s string) []int {
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
