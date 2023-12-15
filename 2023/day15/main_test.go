package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

func Test_part1(t *testing.T) {
	assert.Equal(t, 52, part1("HASH"))
	assert.Equal(t, 1320, part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 145, part2(testInput))
}
