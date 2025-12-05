package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 3, part1(parse(testInput)))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 14, part2(parse(testInput)))
}
