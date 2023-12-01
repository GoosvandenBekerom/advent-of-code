package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput1 = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 142, part1(testInput1))
}

var testInput2 = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 281, part2(testInput2))
}
