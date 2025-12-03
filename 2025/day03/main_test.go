package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 357, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 3121910778619, part2(testInput))
}
