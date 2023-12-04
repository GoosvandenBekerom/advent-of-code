package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = []string{
	"467..114..",
	"...*......",
	"..35...633",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 4361, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 467835, part2(testInput))
}
