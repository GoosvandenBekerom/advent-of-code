package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var score int
	for _, line := range strings.Split(input, "\n") {
		opponent := line[0]
		you := line[2]
		score += rockPaperScissors(opponent, you) + scores[you]
	}

	return score
}

func part2(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	points := map[uint8]int{
		'X': 0,
		'Y': 3,
		'Z': 6,
	}

	var score int

	for _, line := range strings.Split(input, "\n") {
		opponent := line[0]
		outcome := line[2]

		you := calculateOwnMove(opponent, outcome)
		score += scores[you] + points[outcome]
	}

	return score
}

var scores = map[uint8]int{
	'A': 1, // ROCK
	'X': 1, // ROCK

	'B': 2, // PAPER
	'Y': 2, // PAPER

	'C': 3, // SCISSORS
	'Z': 3, // SCISSORS
}

func rockPaperScissors(a, b uint8) int {
	if scores[a] == scores[b] {
		return 3
	}

	switch a {
	case 'A':
		if b == 'Y' {
			return 6
		}
		return 0
	case 'B':
		if b == 'X' {
			return 0
		}
		return 6
	case 'C':
		if b == 'X' {
			return 6
		}
		return 0
	default:
		panic("unknown input")
	}
}

func calculateOwnMove(opponent, outcome uint8) uint8 {
	switch outcome {
	case 'Y': // draw
		return opponent
	case 'X': // loss
		switch opponent {
		case 'A':
			return 'Z'
		case 'B':
			return 'X'
		case 'C':
			return 'Y'
		}
	case 'Z': // win
		switch opponent {
		case 'A':
			return 'Y'
		case 'B':
			return 'Z'
		case 'C':
			return 'X'
		}
	}
	panic("unknown input")
}
