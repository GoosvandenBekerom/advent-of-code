package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func Test_part1(t *testing.T) {
	assert.Equal(t, int64(95437), part1(testInput))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, int64(24933642), part2(testInput))
}
