package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 24, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 93, part2(testInput))
}
