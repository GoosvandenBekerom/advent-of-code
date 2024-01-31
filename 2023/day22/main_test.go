package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"1,0,1~1,2,1",
	"0,0,2~2,0,2",
	"0,2,3~2,2,3",
	"0,0,4~0,2,4",
	"2,0,5~2,2,5",
	"0,1,6~2,1,6",
	"1,1,8~1,1,9",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 0, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 0, part2(testInput))
}
