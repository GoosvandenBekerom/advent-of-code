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

	fmt.Printf("num of ranges fully contained: %d\n", sumContaining)
}

func part2() {
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

	fmt.Printf("num of ranges overlapping: %d\n", sumOverlapping)
}

type Range struct {
	lower, upper int
}

func parseRange(s string) Range {
	left, right, _ := strings.Cut(s, "-")
	var r Range

	lower, err := strconv.Atoi(left)
	check(err)
	r.lower = lower

	upper, err := strconv.Atoi(right)
	check(err)
	r.upper = upper

	return r
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
