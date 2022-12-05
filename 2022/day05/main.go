package main

import (
	_ "embed"
	"github.com/GoosvandenBekerom/advent-of-code/datastructures"
	"strconv"
	"strings"

	"fmt"
	"log"
)

//go:embed input
var input string

func main() {
	stacks, moves := parseInput()

	message := part1(stacks, moves)
	fmt.Printf("message: %s\n", message)

	stacks, moves = parseInput()
	message = part2(stacks, moves)
	fmt.Printf("message: %s\n", message)
}

// ----------------------------------------
// solution
// ----------------------------------------

type move struct {
	amount, from, to int
}

func parseInput() (stacks map[int]*datastructures.Stack[string], moves []move) {
	rawStacks, rawMoves, _ := strings.Cut(input, "\n\n")

	lines := strings.Split(rawStacks, "\n")

	stacks = make(map[int]*datastructures.Stack[string])
	for _, char := range strings.ReplaceAll(lines[len(lines)-1], " ", "") {
		var stack datastructures.Stack[string]
		stacks[toInt(string(char))] = &stack
	}

	for lineNum := len(lines) - 2; lineNum >= 0; lineNum-- {
		for i := 0; i < len(stacks); i++ {
			from := i * 4
			to := from + 3
			if to > len(lines[lineNum]) {
				break
			}

			letter := lines[lineNum][from+1 : to-1]

			if strings.TrimSpace(letter) == "" {
				continue
			}

			stacks[i+1].Push(letter)
		}
	}

	for _, rawMove := range strings.Split(rawMoves, "\n") {
		split := strings.Split(rawMove, " ")
		moves = append(moves, move{
			amount: toInt(split[1]),
			from:   toInt(split[3]),
			to:     toInt(split[5]),
		})
	}

	return stacks, moves
}

func part1(stacks map[int]*datastructures.Stack[string], moves []move) string {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	for _, m := range moves {
		for i := 0; i < m.amount; i++ {
			value, ok := stacks[m.from].Pop()
			if !ok {
				panic("empty stack during moves")
			}

			stacks[m.to].Push(value)
		}
	}

	return calculateMessage(stacks)
}

func part2(stacks map[int]*datastructures.Stack[string], moves []move) string {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	for i, m := range moves {
		boxes, ok := stacks[m.from].Take(m.amount)
		if !ok {
			panic("taking to much from stack " + strconv.Itoa(i))
		}

		for _, box := range boxes {
			stacks[m.to].Push(box)
		}
	}

	return calculateMessage(stacks)
}

func calculateMessage(stacks map[int]*datastructures.Stack[string]) string {
	var message string
	for i := 1; i < len(stacks)+1; i++ {
		stack := stacks[i]
		letter, ok := stack.Pop()
		if !ok {
			panic("empty stack during calculating message")
		}

		message += letter
	}

	return message
}

// ----------------------------------------
// utils
// ----------------------------------------

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
