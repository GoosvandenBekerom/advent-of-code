package main

import (
	_ "embed"
	"regexp"
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

// group 1: x1, group 2: x2, group 3: y1, group 4: y2
var pattern = regexp.MustCompile("(\\d+),(\\d+) -> (\\d+),(\\d+)")

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	grid := make(map[string]int)

	for _, line := range strings.Split(input, "\n") {
		matches := pattern.FindStringSubmatch(line)
		x1r := toInt(matches[1])
		y1r := toInt(matches[2])
		x2r := toInt(matches[3])
		y2r := toInt(matches[4])

		x1 := min(x1r, x2r)
		x2 := max(x1r, x2r)
		y1 := min(y1r, y2r)
		y2 := max(y1r, y2r)

		if x1 == x2 {
			// horizontal line
			for y := y1; y <= y2; y++ {
				key := fmt.Sprintf("%d-%d", x1, y)
				grid[key]++
			}
		} else if y1 == y2 {
			// vertical line
			for x := x1; x <= x2; x++ {
				key := fmt.Sprintf("%d-%d", x, y1)
				grid[key]++
			}
		}
	}

	var solution int
	for _, overlaps := range grid {
		if overlaps > 1 {
			solution++
		}
	}
	fmt.Printf("solution: %d\n", solution)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	grid := make(map[string]int)

	for _, line := range strings.Split(input, "\n") {
		matches := pattern.FindStringSubmatch(line)
		x1r := toInt(matches[1])
		y1r := toInt(matches[2])
		x2r := toInt(matches[3])
		y2r := toInt(matches[4])

		x1 := min(x1r, x2r)
		x2 := max(x1r, x2r)
		y1 := min(y1r, y2r)
		y2 := max(y1r, y2r)

		if x1 == x2 {
			// horizontal line
			for y := y1; y <= y2; y++ {
				key := fmt.Sprintf("%d-%d", x1, y)
				grid[key]++
			}
			continue
		}

		if y1 == y2 {
			// vertical line
			for x := x1; x <= x2; x++ {
				key := fmt.Sprintf("%d-%d", x, y1)
				grid[key]++
			}
			continue
		}

		if x1r == x1 {
			y1 = y1r
			y2 = y2r
		} else {
			y1 = y2r
			y2 = y1r
		}

		// diagonal line
		var step int
		if y1 > y2 {
			step = -1
		} else {
			step = 1
		}

		y := y1
		for x := x1; x <= x2; x++ {
			key := fmt.Sprintf("%d-%d", x, y)
			grid[key]++
			y += step
		}
	}

	var solution int
	for _, overlaps := range grid {
		if overlaps > 1 {
			solution++
		}
	}
	fmt.Printf("solution: %d\n", solution)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
