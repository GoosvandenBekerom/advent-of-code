package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 4277556, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 3263827, part2(testInput))
}
