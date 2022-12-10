package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/data"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	stacks, moves := parseInput()
	fmt.Println(part1(stacks, moves))

	stacks, moves = parseInput()
	fmt.Println(part2(stacks, moves))
}

type move struct {
	amount, from, to int
}

func parseInput() (stacks map[int]*data.Stack[string], moves []move) {
	rawStacks, rawMoves, _ := strings.Cut(input, "\n\n")

	lines := strings.Split(rawStacks, "\n")

	stacks = make(map[int]*data.Stack[string])
	for _, char := range strings.ReplaceAll(lines[len(lines)-1], " ", "") {
		var stack data.Stack[string]
		stacks[utils.ToInt(string(char))] = &stack
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
			amount: utils.ToInt(split[1]),
			from:   utils.ToInt(split[3]),
			to:     utils.ToInt(split[5]),
		})
	}

	return stacks, moves
}

func part1(stacks map[int]*data.Stack[string], moves []move) string {
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

func part2(stacks map[int]*data.Stack[string], moves []move) string {
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

func calculateMessage(stacks map[int]*data.Stack[string]) string {
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
