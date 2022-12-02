package main

import (
	_ "embed"
	"strconv"

	"fmt"
	"log"
	"strings"
)

//go:embed input
var input string

/*
A = 1 ROCK
B = 2 PAPER
C = 3 SCISSORS
LOSS = 0
DRAW = 3
WIN = 6
*/

func main() {
	part1()
	part2()
}

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var score int

	for _, line := range strings.Split(input, "\n") {
		opponent := line[0]
		you := line[2]
		score += rockPaperScissors(opponent, you) + scores[you]
	}

	fmt.Printf("total score: %d\n", score)
}

func part2() {
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

	fmt.Printf("total score: %d\n", score)
}

// ----------------------------------------
// utils
// ----------------------------------------

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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	check(err)
	return v
}
