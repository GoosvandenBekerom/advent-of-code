package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"2413432311323",
	"3215453535623",
	"3255245654254",
	"3446585845452",
	"4546657867536",
	"1438598798454",
	"4457876987766",
	"3637877979653",
	"4654967986887",
	"4564679986453",
	"1224686865563",
	"2546548887735",
	"4322674655533",
}

var testInput2 = []string{
	"111111111111",
	"999999999991",
	"999999999991",
	"999999999991",
	"999999999991",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 102, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 94, part2(testInput))
	assert.Equal(t, 71, part2(testInput2))
}
