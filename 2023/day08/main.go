package main

import (
	"container/ring"
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/math"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type node struct {
	value       string
	left, right string
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	directions := ring.New(len(lines[0]))
	for _, dir := range lines[0] {
		directions.Value = dir
		directions = directions.Next()
	}

	nodes := make(map[string]node)

	for _, line := range lines[2:] {
		nodes[line[:3]] = node{
			value: line[:3],
			left:  line[7:10],
			right: line[12:15],
		}
	}

	var steps int
	current := nodes["AAA"]
	for {
		steps++
		if directions.Value.(rune) == 'L' {
			current = nodes[current.left]
		} else {
			current = nodes[current.right]
		}
		if current.value == "ZZZ" {
			break
		}
		directions = directions.Next()
	}

	return steps
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	directions := ring.New(len(lines[0]))
	for _, dir := range lines[0] {
		directions.Value = dir
		directions = directions.Next()
	}

	nodes := make(map[string]node)
	var currentPositions []node

	for _, line := range lines[2:] {
		val := line[:3]
		nodes[val] = node{
			value: val,
			left:  line[7:10],
			right: line[12:15],
		}
		if val[2] == 'A' {
			currentPositions = append(currentPositions, nodes[val])
		}
	}

	steps := 0
	done := 0
	doneSteps := make([]int, len(currentPositions))
	for {
		steps++
		for i, current := range currentPositions {
			if doneSteps[i] != 0 {
				continue
			}
			if directions.Value.(rune) == 'L' {
				currentPositions[i] = nodes[current.left]
			} else {
				currentPositions[i] = nodes[current.right]
			}
			if currentPositions[i].value[2] == 'Z' {
				doneSteps[i] = steps
				done++
			}
		}

		if done == len(currentPositions) {
			break
		}

		directions = directions.Next()
	}

	return math.LeastCommonMultiple(doneSteps[0], doneSteps[1], doneSteps[2:]...)
}
