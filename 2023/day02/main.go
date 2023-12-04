package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	var sum int

gameLoop:
	for _, line := range lines {
		game, rounds, _ := strings.Cut(line, ": ")
		for _, round := range strings.Split(rounds, "; ") {
			for _, play := range strings.Split(round, ", ") {
				number, color, _ := strings.Cut(play, " ")
				amount := utils.ToInt(number)
				switch color {
				case "red":
					if amount > 12 {
						continue gameLoop
					}
				case "green":
					if amount > 13 {
						continue gameLoop
					}
				case "blue":
					if amount > 14 {
						continue gameLoop
					}
				}
			}
		}
		sum += utils.ToInt(strings.Split(game, " ")[1])
	}

	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	return sum
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	var sum int
	for _, line := range lines {
		minimum := make(map[string]int)
		_, rounds, _ := strings.Cut(line, ": ")
		for _, round := range strings.Split(rounds, "; ") {
			for _, play := range strings.Split(round, ", ") {
				number, color, _ := strings.Cut(play, " ")
				amount := utils.ToInt(number)
				if minimum[color] < amount {
					minimum[color] = amount
				}
			}
		}
		sum += minimum["red"] * minimum["green"] * minimum["blue"]
	}

	return sum
}
