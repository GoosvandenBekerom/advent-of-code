package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testInput() []monkey {
	return []monkey{
		{
			items: []int{79, 98},
			calculateWorryLevel: func(i int) int {
				return i * 19
			},
			divider: 23,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 2
				}
				return 3
			},
		},
		{
			items: []int{54, 65, 75, 74},
			calculateWorryLevel: func(i int) int {
				return i + 6
			},
			divider: 19,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 2
				}
				return 0
			},
		},
		{
			items: []int{79, 60, 97},
			calculateWorryLevel: func(i int) int {
				return i * i
			},
			divider: 13,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 1
				}
				return 3
			},
		},
		{
			items: []int{74},
			calculateWorryLevel: func(i int) int {
				return i + 3
			},
			divider: 17,
			findMonkeyToThrowTo: func(new int, divider int) int {
				if new%divider == 0 {
					return 0
				}
				return 1
			},
		},
	}
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 10605, part1(testInput(), true))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 2713310158, part2(testInput(), false))
}
