package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	`.|...\....`,
	`|.-.\.....`,
	`.....|-...`,
	`........|.`,
	`..........`,
	`.........\`,
	`..../.\\..`,
	`.-.-/..|..`,
	`.|....-|.\`,
	`..//.|....`,
}

func Test_part1(t *testing.T) {
	assert.Equal(t, 46, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 51, part2(testInput))
}
