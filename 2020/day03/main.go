package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"log"
	"strings"
)

type grid [][]bool
type slope struct {
	right, down int
}

func main() {
	grid := input()
	trees1 := part1(grid)
	fmt.Printf("part 1: %d\n", trees1)
	trees2 := part2(grid)
	fmt.Printf("part 2: %d\n", trees2)
}

func part1(input grid) (count int) {
	return calculateTrees(input, slope{down: 1, right: 3})
}

func part2(input grid) (count int) {
	r1 := calculateTrees(input, slope{down: 1, right: 1})
	r3 := calculateTrees(input, slope{down: 1, right: 3})
	r5 := calculateTrees(input, slope{down: 1, right: 5})
	r7 := calculateTrees(input, slope{down: 1, right: 7})
	d2 := calculateTrees(input, slope{down: 2, right: 1})
	return r1 * r3 * r5 * r7 * d2
}

func calculateTrees(input grid, slope slope) (count int) {
	x := 0
	width := len(input[0])
	for y := 0; y < len(input); y += slope.down {
		if input[y][x] {
			count++
		}
		x = (x + slope.right) % width
	}
	return
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (output grid) {
	box := packr.New("day03", "./2020/day03")
	s, err := box.FindString("input")
	check(err)
	for row, line := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		output = append(output, make([]bool, len(line)))
		for col, c := range line {
			output[row][col] = c == '#'
		}
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
