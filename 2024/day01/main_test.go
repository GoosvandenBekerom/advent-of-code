package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 11, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 31, part2(testInput))
}
