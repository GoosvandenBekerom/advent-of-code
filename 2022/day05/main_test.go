package main

import (
	"github.com/GoosvandenBekerom/advent-of-code/datastructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part1(t *testing.T) {
	stacks, moves := map[int]*datastructures.Stack[string]{
		1: {"Z", "N"},
		2: {"M", "C", "D"},
		3: {"P"},
	}, []move{
		{amount: 1, from: 2, to: 1},
		{amount: 3, from: 1, to: 3},
		{amount: 2, from: 2, to: 1},
		{amount: 1, from: 1, to: 2},
	}

	assert.Equal(t, "CMZ", part1(stacks, moves))
}

func Test_part2(t *testing.T) {
	stacks, moves := map[int]*datastructures.Stack[string]{
		1: {"Z", "N"},
		2: {"M", "C", "D"},
		3: {"P"},
	}, []move{
		{amount: 1, from: 2, to: 1},
		{amount: 3, from: 1, to: 3},
		{amount: 2, from: 2, to: 1},
		{amount: 1, from: 1, to: 2},
	}

	assert.Equal(t, "MCD", part2(stacks, moves))
}
