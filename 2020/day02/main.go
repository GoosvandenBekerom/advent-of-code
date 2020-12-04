package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	entries := input()
	valid1 := part1(entries)
	fmt.Printf("part 1: %d\n", valid1)
	valid2 := part2(entries)
	fmt.Printf("part 2: %d\n", valid2)
}

type entry struct {
	left, right int
	char        rune
	pass        []rune
}

func part1(entries []entry) (valid int) {
	for _, entry := range entries {
		count := 0
		for _, c := range entry.pass {
			if c == entry.char {
				count++
			}
		}
		if count >= entry.left && count <= entry.right {
			valid++
		}
	}
	return
}

func part2(entries []entry) (valid int) {
	for _, entry := range entries {
		l := entry.pass[entry.left-1]
		r := entry.pass[entry.right-1]
		if l == r {
			continue
		}
		if l == entry.char || r == entry.char {
			valid++
		}
	}
	return
}

// ----------------------------------------
// utils
// ----------------------------------------

type rangeStr string

func (s rangeStr) parse() (int, int) {
	nums := strings.Split(string(s), "-")
	left, err := strconv.Atoi(nums[0])
	check(err)
	right, err := strconv.Atoi(nums[1])
	check(err)
	return left, right
}

func input() (entries []entry) {
	box := packr.New("day02", "./2020/day02")
	s, err := box.FindString("input")
	check(err)
	for _, line := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		parts := strings.Split(line, " ")
		l, h := rangeStr(parts[0]).parse()
		e := entry{
			left:  l,
			right: h,
			char:  []rune(parts[1])[0],
			pass:  []rune(parts[2]),
		}
		entries = append(entries, e)
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
