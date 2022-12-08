package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testGrid = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 21, part1(testGrid))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 8, part2(testGrid))
}
