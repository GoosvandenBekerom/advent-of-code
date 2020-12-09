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
	cypher := input()
	return findInvalidNumber(cypher, 25)
}

func part2() int {
	cypher := input()
	num := findInvalidNumber(cypher, 25)
	sum := findContiguousSum(cypher, num)
	sort.Ints(sum)
	return sum[0] + sum[len(sum)-1]
}

func valid(want int, ints []int) bool {
	for i, v1 := range ints {
		for j, v2 := range ints {
			if i == j {
				continue
			}
			if v1+v2 == want {
				return true
			}
		}
	}
	return false
}

func findInvalidNumber(cypher []int, preamble int) int {
	for i := preamble; i < len(cypher)-1; i++ {
		if !valid(cypher[i], cypher[i-preamble:i]) {
			return cypher[i]
		}
	}
	panic("no invalid number found in cypher")
}

func findContiguousSum(cypher []int, want int) []int {
	for i := 0; i < len(cypher); i++ {
		sum := cypher[i]
		for j := i + 1; j < len(cypher); j++ {
			sum += cypher[j]
			if sum > want {
				break
			}
			if sum == want {
				return cypher[i : j+1]
			}
		}
	}
	panic("no contiguous sum found in cypher")
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (cypher []int) {
	box := packr.New("day09", "./2020/day09")
	s, err := box.FindString("input")
	check(err)
	for _, in := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		num, err := strconv.Atoi(in)
		check(err)
		cypher = append(cypher, num)
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
