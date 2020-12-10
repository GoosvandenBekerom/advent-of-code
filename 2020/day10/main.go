package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	fmt.Printf("part 1: %d\n", part1())
	fmt.Printf("part 2: %d\n", part2())
}

func part1() int {
	adapters := input()
	sort.Ints(adapters)
	ones, twos, threes, prev := 0, 0, 0, 0
	threes++ // add own device
	for _, adapter := range adapters {
		if adapter == prev+1 {
			ones++
		} else if adapter == prev+2 {
			twos++
		} else if adapter == prev+3 {
			threes++
		} else {
			panic("no adapter found with difference of 1, 2 or 3")
		}
		prev = adapter
	}
	fmt.Printf("ones %d, two's %d, three's %d,\n", ones, twos, threes)
	return ones * threes
}

func part2() int {
	adapters := input()
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3) // add own device
	fmt.Printf("%v\n", adapters)
	return findCombinations(adapters, 0, make(map[int]int))
}

func findCombinations(adapters []int, start int, cache map[int]int) int {
	cached, cacheHit := cache[start]
	if cacheHit {
		return cached
	}
	count := 0
	for _, adapter := range adapters {
		diff := adapter - start
		if diff <= 3 && diff > 0 {
			count += findCombinations(adapters, adapter, cache)
		}
	}
	// If we are at the end then we have reached a valid adaptor chain
	if start == adapters[len(adapters)-1] {
		count++
	}
	cache[start] = count
	return count
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (adapters []int) {
	box := packr.New("day10", "./2020/day10")
	s, err := box.FindString("input")
	check(err)
	for _, in := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		num, err := strconv.Atoi(in)
		check(err)
		adapters = append(adapters, num)
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
