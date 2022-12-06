package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part1(t *testing.T) {
	assert.Equal(t, 7, part1("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 5, part1("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 6, part1("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 10, part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 11, part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 19, part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	assert.Equal(t, 23, part2("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	assert.Equal(t, 23, part2("nppdvjthqldpwncqszvftbrmjlhg"))
	assert.Equal(t, 29, part2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	assert.Equal(t, 26, part2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
