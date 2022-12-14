package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 31, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 29, part2(testInput))
}
