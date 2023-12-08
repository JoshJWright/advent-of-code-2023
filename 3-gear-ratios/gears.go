package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Engine struct {
	parts   []PartNumber
	symbols []Symbol
}

type PartNumber struct {
	num, line, start, length int
	valid                    bool
}

type Symbol struct {
	line, index int
	isGear      bool
	gearRatio   int
}

func main() {
	raw, _ := os.ReadFile("input.txt")
	input := string(raw)

	engine := &Engine{}

	for line, data := range strings.Split(input, "\n") {
		engine.parseLine(line, data)
	}

	engine.validateParts()

	var sum int
	for _, symbol := range engine.symbols {
		if symbol.isGear {
			sum += symbol.gearRatio
		}
	}

	fmt.Println(sum)
}

func (engine *Engine) parseLine(line int, s string) {
	partNumberMatcher := regexp.MustCompile(`\d+`)
	for _, partNumberMatch := range partNumberMatcher.FindAllStringIndex(s, -1) {
		partNumber, _ := strconv.Atoi(s[partNumberMatch[0]:partNumberMatch[1]])

		engine.parts = append(engine.parts, PartNumber{
			num:    partNumber,
			start:  partNumberMatch[0],
			length: partNumberMatch[1] - partNumberMatch[0],
			line:   line,
			valid:  false,
		})
	}

	symbolMatcher := regexp.MustCompile(`[^0-9\.]`)
	for _, symbolMatch := range symbolMatcher.FindAllStringIndex(s, -1) {
		engine.symbols = append(engine.symbols, Symbol{line: line, index: symbolMatch[0]})
	}
}

func (engine *Engine) validateParts() {
	for sym := range engine.symbols {
		symbol := &engine.symbols[sym]
		adjacent := 0

		for p := range engine.parts {
			part := &engine.parts[p]

			isAdjacent := symbol.line >= part.line-1 &&
				symbol.line <= part.line+1 &&
				symbol.index >= part.start-1 &&
				symbol.index <= part.start+part.length

			if isAdjacent {
				adjacent++
				switch adjacent {
				case 1:
					symbol.gearRatio = part.num
				case 2:
					symbol.gearRatio = symbol.gearRatio * part.num
					symbol.isGear = true
				case 3:
					symbol.isGear = false
				}
			}
			part.valid = part.valid || isAdjacent

		}
	}
}
