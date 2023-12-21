package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"...........",
	".....###.#.",
	".###.##..#.",
	"..#.#...#..",
	"....#.#....",
	".##..S####.",
	".##..#...#.",
	".......##..",
	".##.#.####.",
	".##..##.##.",
	"...........",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 16, part1(testInput, 6))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 0, part2(testInput))
}
