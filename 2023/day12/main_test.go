package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 21, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 525152, part2(testInput))
}
