package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 288, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 71503, part2(testInput))
}
