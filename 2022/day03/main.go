package main

import (
	"bytes"
	_ "embed"
	"strings"

	"fmt"
	//"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var sum int
	for _, line := range strings.Split(input, "\n") {
		left, right := split(line)

	outer:
		for _, l := range left {
			for _, r := range right {
				if l == r {
					sum += priority(uint8(l))
					break outer
				}
			}
		}
	}

	return sum
}

func part2(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var previous []byte
	var sum int
	for i, line := range strings.Split(input, "\n") {
		if i%3 == 0 {
			if len(previous) != 0 {
				// first time we get in this loop previous is not set yet
				sum += priority(previous[0])
			}

			// first of new group
			previous = []byte(line)
			continue
		}

		var common []byte

		for _, b := range []byte(line) {
			if bytes.IndexByte(previous, b) != -1 {
				common = append(common, b)
			}
		}

		previous = common
	}

	// make sure the last badge is added
	sum += priority(previous[0])
	return sum
}

func split(s string) (string, string) {
	cut := len(s) / 2
	return s[:cut], s[cut:]
}

func priority(char uint8) int {
	if char > 90 {
		return int(char) - 96
	}
	return int(char) - 38
}
