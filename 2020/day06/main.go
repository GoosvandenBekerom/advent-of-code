package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	questions := input()
	fmt.Printf("part 1: %d\n", part1(questions))
	fmt.Printf("part 2: %d\n", part2(questions))
}

func part1(groups []group) (sum int) {
	for _, group := range groups {
		set := make(map[rune]bool)
		for _, c := range strings.Join(group, "") {
			set[c] = true
		}
		sum += len(set)
	}
	return
}

func part2(groups []group) (sum int) {
	for _, group := range groups {
		set := make(map[rune]int)
		for _, c := range strings.Join(group, "") {
			set[c]++
		}
		for _, n := range set {
			if n == len(group) {
				sum++
			}
		}
	}
	return
}

// ----------------------------------------
// utils
// ----------------------------------------

type group []string

func input() (groups []group) {
	box := packr.New("day06", "./2020/day06")
	s, err := box.FindString("input")
	check(err)
	for _, in := range strings.Split(strings.TrimSuffix(s, "\n"), "\n\n") {
		groups = append(groups, strings.Split(in, "\n"))
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
