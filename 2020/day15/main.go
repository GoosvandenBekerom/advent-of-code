package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	println(part1())
	println(part2())
}

type entry struct {
	last, before int
}

func part1() int {
	return solve(2020)
}

func part2() int {
	return solve(30000000)
}

func solve(limit int) int {
	// history is a map with a number as its key, and its last occurrence as its value
	history := make(map[int]entry)
	i := 0
	prev := 0
	for n := range input() {
		history[n] = entry{last: i, before: -1}
		i++
		prev = n
	}
	for ; i < limit; i++ {
		current := 0
		if pPos, pExists := history[prev]; pExists {
			if pPos.before >= 0 {
				current = pPos.last - pPos.before
			}
			if old, cExists := history[current]; cExists {
				history[current] = entry{last: i, before: old.last}
			} else {
				history[current] = entry{last: i, before: -1}
			}
		}
		prev = current
	}
	return prev
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (nums chan int) {
	box := packr.New("day15", "./2020/day15")
	s, err := box.FindString("input")
	check(err)
	nums = make(chan int, 1)
	go func() {
		defer close(nums)
		for _, raw := range strings.Split(strings.TrimSuffix(s, "\n"), ",") {
			num, err := strconv.Atoi(raw)
			check(err)
			nums <- num
		}
	}()
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
