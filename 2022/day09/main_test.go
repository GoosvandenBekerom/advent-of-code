package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part1(t *testing.T) {
	testInput := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	assert.Equal(t, 13, part1(testInput))
}

func Test_part2(t *testing.T) {
	testInput := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	assert.Equal(t, 36, part2(testInput))
}
