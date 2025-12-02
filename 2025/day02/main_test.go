package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func Test_part1(t *testing.T) {
	assert.Equal(t, 1227775554, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 4174379265, part2(testInput))
}
