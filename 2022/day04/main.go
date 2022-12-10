package main

import (
	_ "embed"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"strconv"

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

	var sumContaining int
	for _, line := range strings.Split(input, "\n") {
		left, right, _ := strings.Cut(line, ",")

		range1 := parseRange(left)
		range2 := parseRange(right)

		if (range1.lower >= range2.lower && range1.upper <= range2.upper) ||
			(range2.lower >= range1.lower && range2.upper <= range1.upper) {
			sumContaining++
		}
	}

	return sumContaining
}

func part2(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var sumOverlapping int
	for _, line := range strings.Split(input, "\n") {
		left, right, _ := strings.Cut(line, ",")

		range1 := parseRange(left)
		range2 := parseRange(right)

		if (range1.lower >= range2.lower && range1.lower <= range2.upper) ||
			(range1.upper >= range2.lower && range1.upper <= range2.upper) {
			sumOverlapping++
			continue
		}

		if (range2.lower >= range1.lower && range2.lower <= range1.upper) ||
			(range2.upper >= range1.lower && range2.upper <= range1.upper) {
			sumOverlapping++
			continue
		}
	}

	return sumOverlapping
}

type Range struct {
	lower, upper int
}

func parseRange(s string) Range {
	left, right, _ := strings.Cut(s, "-")
	var r Range

	lower, err := strconv.Atoi(left)
	utils.Check(err)
	r.lower = lower

	upper, err := strconv.Atoi(right)
	utils.Check(err)
	r.upper = upper

	return r
}
