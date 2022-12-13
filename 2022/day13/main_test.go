package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = [][2]string{
	{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
	},
	{
		"[[1],[2,3,4]]",
		"[[1],4]",
	},
	{
		"[9]",
		"[[8,7,6]]",
	},
	{
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
	},
	{
		"[7,7,7,7]",
		"[7,7,7]",
	},
	{
		"[]",
		"[3]",
	},
	{
		"[[[]]]",
		"[[]]",
	},
	{
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	},
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 13, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 140, part2(testInput))
}
