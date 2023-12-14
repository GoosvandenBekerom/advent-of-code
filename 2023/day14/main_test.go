package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 136, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 64, part2(testInput))
}
