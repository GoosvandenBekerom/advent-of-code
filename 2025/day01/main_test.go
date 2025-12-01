package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 3, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 6, part2(testInput))
}
