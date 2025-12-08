package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/utils"
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

	type equation struct {
		nums []int
		add  bool
	}
	var equations []equation
	for _, line := range lines {
		fields := strings.Fields(line)
		if equations == nil {
			equations = make([]equation, len(fields))
		}
		for i, field := range fields {
			if field == "+" || field == "*" {
				equations[i].add = field == "+"
				continue
			}
			equations[i].nums = append(equations[i].nums, utils.ToInt(field))
		}
	}

	answer := 0
	for _, e := range equations {
		toAdd := 0
		if !e.add {
			toAdd = 1
		}
		for _, num := range e.nums {
			if e.add {
				toAdd += num
			} else {
				toAdd *= num
			}
		}
		answer += toAdd
	}

	return answer
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	nums := make([]string, len(lines[0]))
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines)-1; y++ {
			nums[x] += string(lines[y][x])
		}
	}

	operators := strings.Fields(lines[len(lines)-1])
	operatorI := 0
	toAdd := 0
	if operators[operatorI] == "*" {
		toAdd = 1
	}

	answer := 0
	for _, num := range nums {
		trimmed := strings.TrimSpace(num)
		if trimmed == "" {
			answer += toAdd
			operatorI++
			toAdd = 0
			if operators[operatorI] == "*" {
				toAdd = 1
			}
			continue
		}
		value := utils.ToInt(trimmed)
		if operators[operatorI] == "*" {
			toAdd *= value
		} else {
			toAdd += value
		}
	}
	answer += toAdd

	return answer
}
