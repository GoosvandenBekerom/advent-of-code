package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/data"
	"github.com/GoosvandenBekerom/advent-of-code/data/directions"
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

	grid := make(map[data.Vector]byte)

	for y, line := range lines {
		for x, c := range []byte(line) {
			grid[data.Vector{X: x, Y: y}] = c
		}
	}

	answer := 0
	for y, line := range lines {
		for x := range line {
			pos := data.Vector{X: x, Y: y}
			if grid[pos] != '@' {
				continue
			}
			rollsAround := 0
			for _, dir := range directions.All {
				if grid[pos.Add(dir.Vector)] == '@' {
					rollsAround++
					if rollsAround >= 4 {
						break
					}
				}
			}
			if rollsAround < 4 {
				fmt.Println(x, y)
				answer++
			}
		}
	}

	return answer
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	grid := make(map[data.Vector]byte)

	for y, line := range lines {
		for x, c := range []byte(line) {
			grid[data.Vector{X: x, Y: y}] = c
		}
	}

	answer := 0

	for {
		answerAtStart := answer
		for y, line := range lines {
			for x := range line {
				pos := data.Vector{X: x, Y: y}
				if grid[pos] != '@' {
					continue
				}
				rollsAround := 0
				for _, dir := range directions.All {
					if grid[pos.Add(dir.Vector)] == '@' {
						rollsAround++
						if rollsAround >= 4 {
							break
						}
					}
				}
				if rollsAround < 4 {
					grid[pos] = '.'
					answer++
				}
			}
		}
		if answer == answerAtStart {
			break
		}
	}

	return answer
}
