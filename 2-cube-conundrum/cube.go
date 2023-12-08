package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	id   int
	info []CubeRecord
}

type CubeRecord struct {
	red   int
	green int
	blue  int
}

func main() {
	rawInput, _ := os.ReadFile("input.txt")
	input := string(rawInput)
	games := parseGames(input)

	// limit := CubeRecord{red: 12, green: 13, blue: 14}
	var sum int
	for _, game := range games {
		// if game.isWithin(limit) {
		// 	sum += game.id
		// }

		sum += game.computePower()
	}
	fmt.Println(sum)
}

func (game Game) isWithin(limit CubeRecord) bool {
	for _, record := range game.info {
		if record.red > limit.red || record.green > limit.green || record.blue > limit.blue {
			return false
		}
	}
	return true
}

func (game Game) computePower() int {
	minimums := CubeRecord{}
	for _, record := range game.info {
		minimums = CubeRecord{
			red:   int(math.Max(float64(minimums.red), float64(record.red))),
			green: int(math.Max(float64(minimums.green), float64(record.green))),
			blue:  int(math.Max(float64(minimums.blue), float64(record.blue))),
		}
	}
	return minimums.red * minimums.green * minimums.blue
}

func parseGames(input string) []Game {
	gameData := strings.Split(input, "\n")

	games := make([]Game, 0, len(gameData))

	var gameId *regexp.Regexp = regexp.MustCompile(`Game (\d+): `)

	for _, data := range gameData {
		idMatch := gameId.FindStringSubmatch(data)
		id, _ := strconv.Atoi(idMatch[1])
		cubeSets := strings.Split(data[len(idMatch[0]):], `;`)

		game := Game{id: id, info: make([]CubeRecord, len(cubeSets))}

		var cubes *regexp.Regexp = regexp.MustCompile(`(\d+) (\w+)`)
		for i, cubeSet := range cubeSets {
			parsedCubes := cubes.FindAllStringSubmatch(cubeSet, -1)
			for _, colour := range parsedCubes {
				num, _ := strconv.Atoi(colour[1])
				switch strings.ToLower(colour[2]) {
				case "red":
					game.info[i].red = game.info[i].red + num
				case "blue":
					game.info[i].blue = game.info[i].blue + num
				case "green":
					game.info[i].green = game.info[i].green + num
				}
			}
		}
		games = append(games, game)
	}
	return games
}
