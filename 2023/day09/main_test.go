package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 114, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 2, part2(testInput))
}
