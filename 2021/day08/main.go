package main

import (
	_ "embed"
	"sort"
	"strings"

	"fmt"
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
	// for part 1 only consider numbers with unique number of segments (1,4,7,8)
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	const (
		one   = 2
		four  = 4
		seven = 3
		eight = 7
	)

	var count int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "|")
		digits := strings.Split(strings.Trim(parts[1], " "), " ")

		for _, s := range digits {
			switch len(s) {
			case one:
				fallthrough
			case four:
				fallthrough
			case seven:
				fallthrough
			case eight:
				count++
				break
			}
		}
	}

	fmt.Printf("solution: %d\n", count)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var solution int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "|")
		segments := strings.Split(strings.Trim(parts[0], " "), " ")
		digits := strings.Split(strings.Trim(parts[1], " "), " ")

		// find all connections
		connections := make(map[int]string)
		for len(connections) != 10 {
			for _, seg := range segments {
				switch len(seg) {
				// unique lengths (1,4,7,8)
				case 2:
					connections[1] = seg
					break
				case 4:
					connections[4] = seg
					break
				case 3:
					connections[7] = seg
					break
				case 7:
					connections[8] = seg
					break
				// length of 5 (2, 3 or 5)
				case 5:
					// 2, 3, 5 should all intersect 2 chars with with 1
					seg1, found := connections[1]
					if !found {
						break
					}
					// if 2 chars intersect it should be 3
					if len(segment(seg1).intersect(segment(seg))) == 2 {
						connections[3] = seg
						break
					}
					// 2 and 5 should also intersect with 4
					seg4, ok := connections[4]
					if !ok {
						break
					}
					// if 2 chars intersect it should be 2
					if len(segment(seg4).intersect(segment(seg))) == 2 {
						connections[2] = seg
						break
					}
					// else it should be 5
					connections[5] = seg
					break
				// length of 6 (0, 6 or 9)
				case 6:
					// 0, 6, 9 should all intersect with 4
					seg4, found := connections[4]
					if !found {
						break
					}
					// if 4 chars intersect it should be 9
					if len(segment(seg4).intersect(segment(seg))) == 4 {
						connections[9] = seg
						break
					}
					// 0 and 6 should also intersect with 1
					seg1, found := connections[1]
					if !found {
						break
					}
					// if 2 chars intersect it should be 0
					if len(segment(seg1).intersect(segment(seg))) == 2 {
						connections[0] = seg
						break
					}
					// else it should be 6
					connections[6] = seg
					break
				}
			}
		}

		solutions := make(map[string]int)
		for n, seg := range connections {
			solutions[sortChars(seg)] = n
		}
		var total int
		for _, seg := range digits {
			total *= 10
			total += solutions[sortChars(seg)]
		}
		solution += total
	}

	fmt.Printf("solution: %d\n", solution)

}

// ----------------------------------------
// utils
// ----------------------------------------

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }

func sortChars(s string) string {
	var r runes
	for _, rv := range s {
		r = append(r, rv)
	}
	sort.Sort(r)
	return string(r)
}

type segment string

func (s segment) toMap() map[rune]bool {
	result := make(map[rune]bool)
	for _, c := range s {
		result[c] = true
	}
	return result
}

func (s segment) diff(other segment) string {
	var result string
	found := make(map[rune]bool)
	tgt := other.toMap()
	for c := range s.toMap() {
		if found[c] {
			continue
		}
		if tgt[c] {
			continue
		}
		result += string(c)
		found[c] = true
	}
	return result
}

func (s segment) intersect(other segment) string {
	if len(other) > len(s) {
		return other.intersect(s)
	}
	var result string
	found := make(map[rune]bool)
	tgt := other.toMap()
	for c := range s.toMap() {
		if found[c] {
			continue
		}
		if tgt[c] {
			result += string(c)
			found[c] = true
		}
	}
	return result
}
