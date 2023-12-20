package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput1 = []string{
	"broadcaster -> a, b, c",
	"%a -> b",
	"%b -> c",
	"%c -> inv",
	"&inv -> a",
}

var testInput2 = []string{
	"broadcaster -> a",
	"%a -> inv, con",
	"&inv -> b",
	"%b -> con",
	"&con -> output",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 32000000, part1(testInput1))
	assert.Equal(t, 11687500, part1(testInput2))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 0, part2(testInput1))
	assert.Equal(t, 0, part2(testInput2))
}
