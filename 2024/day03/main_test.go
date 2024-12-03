package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	assert.Equal(t, 161, part1([]string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 48, part2([]string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}))
}
