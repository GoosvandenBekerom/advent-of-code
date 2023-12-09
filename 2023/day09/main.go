package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/data"
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
	for _, line := range lines {
		sum += nextInSequence(utils.Map(strings.Split(line, " "), func(item string) int {
			return utils.ToInt(item)
		}))
	}

	return sum
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	var sum int
	for _, line := range lines {
		sequence := utils.Map(strings.Split(line, " "), func(item string) int {
			return utils.ToInt(item)
		})
		utils.ReverseSlice(sequence)
		sum += nextInSequence(sequence)
	}

	return sum
}

func nextInSequence(sequence []int) int {
	sequences := new(data.Stack[[]int])
	sequences.Push(sequence)
	for {
		current, _ := sequences.Pop()
		diff := make([]int, 0, len(current)-1)
		var prev int
		first := true
		for _, num := range current {
			if !first {
				diff = append(diff, num-prev)
			}
			prev = num
			first = false
		}

		sequences.Push(current)

		if utils.All(diff, func(i int) bool { return i == 0 }) {
			break
		}

		sequences.Push(diff)
	}

	var last int
	for !sequences.IsEmpty() {
		v, _ := sequences.Pop()
		last = v[len(v)-1] + last
	}
	return last
}
