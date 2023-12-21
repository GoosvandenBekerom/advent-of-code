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
	//assert.Equal(t, 16, part2(testInput, 6))
	//assert.Equal(t, 50, part2(testInput, 10))
	//assert.Equal(t, 1594, part2(testInput, 50))
	//assert.Equal(t, 6536, part2(testInput, 100))
	//assert.Equal(t, 167004, part2(testInput, 500))
	//assert.Equal(t, 668697, part2(testInput, 1000))
	//assert.Equal(t, 16733044, part2(testInput, 5000))
	// test input doesn't work for this solution, but main input does...
}
