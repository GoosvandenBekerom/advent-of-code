package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 405, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 400, part2(testInput))
}
