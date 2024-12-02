package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"7 6 4 2 1",
	"1 2 7 8 9",
	"9 7 6 2 1",
	"1 3 2 4 5",
	"8 6 4 4 1",
	"1 3 6 7 9",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 2, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 4, part2(testInput))
}
