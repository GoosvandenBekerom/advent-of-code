package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/math"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"golang.org/x/exp/slices"
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
	var left, right []int
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left = append(left, utils.ToInt(parts[0]))
		right = append(right, utils.ToInt(parts[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	var distance int
	for i := range left {
		distance += math.Abs(left[i] - right[i])
	}

	return distance
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var left []int
	count := make(map[int]int)

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left = append(left, utils.ToInt(parts[0]))
		count[utils.ToInt(parts[1])]++
	}

	var answer int
	for _, v := range left {
		answer += v * count[v]
	}

	return answer
}
