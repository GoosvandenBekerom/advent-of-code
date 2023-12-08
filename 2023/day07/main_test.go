package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = []string{
	"2345A 1",
	"Q2KJJ 13",
	"Q2Q2Q 19",
	"T3T3J 17",
	"T3Q33 11",
	"2345J 3",
	"J345A 2",
	"32T3K 5",
	"T55J5 29",
	"KK677 7",
	"KTJJT 34",
	"QQQJA 31",
	"JJJJJ 37",
	"JAAAA 43",
	"AAAAJ 59",
	"AAAAA 61",
	"2AAAA 23",
	"2JJJJ 53",
	"JJJJ2 41",
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 6592, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 6839, part2(testInput))
}
