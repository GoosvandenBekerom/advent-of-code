package main

import (
	"bytes"
	_ "embed"
	"strconv"
	"strings"

	"fmt"
	"log"
	//"strings"
)

//go:embed input
var input string

func main() {
	part1()
	part2()
}

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
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

	fmt.Printf("sum of priorities: %d\n", sum)
}

func part2() {
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

	fmt.Printf("sum of badge priorities: %d\n", sum)
}

// for insights:
// left, right := split("some string")
// fmt.Printf("left:\t%s (%d)\nright:\t%s (%d)\n", left, len(left), right, len(right))
func split(s string) (string, string) {
	cut := len(s) / 2
	return s[:cut], s[cut:]
}

// for insights:
//
//	 fmt.Printf("a = %d\t calculated = %d\n", 'a', priority('a'))
//		fmt.Printf("z = %d\t calculated = %d\n", 'z', priority('z'))
//		fmt.Printf("A = %d\t calculated = %d\n", 'A', priority('A'))
//		fmt.Printf("Z = %d\t calculated = %d\n", 'Z', priority('Z'))
func priority(char uint8) int {
	if char > 90 {
		return int(char) - 96
	}
	return int(char) - 38
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	check(err)
	return v
}
