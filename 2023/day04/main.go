package main

import (
	_ "embed"
	"fmt"
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
	for _, line := range lines {
		_, allNumbers, _ := strings.Cut(line, ": ")
		winningNumbers, ownNumbers, _ := strings.Cut(allNumbers, " | ")
		winning := make(map[string]struct{})
		for _, n := range strings.Split(winningNumbers, " ") {
			if n == "" {
				continue
			}
			winning[n] = struct{}{}
		}
		var score int
		for _, n := range strings.Split(ownNumbers, " ") {
			if _, ok := winning[n]; ok {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		sum += score
	}

	return sum
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	amountOfWinningNumbers := make([]int, len(lines))
	for card, line := range lines {
		_, allNumbers, _ := strings.Cut(line, ": ")
		winningNumbers, ownNumbers, _ := strings.Cut(allNumbers, " | ")
		winning := make(map[string]struct{})
		for _, n := range strings.Split(winningNumbers, " ") {
			if n == "" {
				continue
			}
			winning[n] = struct{}{}
		}
		for _, n := range strings.Split(ownNumbers, " ") {
			if _, ok := winning[n]; ok {
				amountOfWinningNumbers[card]++
			}
		}
	}

	copies := make(map[int]int)
	for card, amount := range amountOfWinningNumbers {
		copies[card] += 1
		for i := 0; i < copies[card]; i++ {
			for j := 1; j <= amount; j++ {
				copies[card+j] += 1
			}
		}
	}

	var sum int
	for _, amount := range copies {
		sum += amount
	}

	return sum
}
