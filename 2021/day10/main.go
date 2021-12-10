package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input
var input string

func main() {
	part1()
	part2()
}

/*
If a chunk opens with (, it must close with ).
If a chunk opens with [, it must close with ].
If a chunk opens with {, it must close with }.
If a chunk opens with <, it must close with >.
*/

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	openings := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	var wrong []rune
	for _, line := range strings.Split(input, "\n") {
		stack := newRuneStack()
		for _, char := range line {
			_, isOpening := openings[char]
			if isOpening {
				stack.push(char)
				continue
			}

			if stack.empty() || openings[stack.pop()] != char {
				wrong = append(wrong, char)
				break
			}
		}
	}

	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	var solution int
	for _, w := range wrong {
		solution += points[w]
	}
	fmt.Printf("solution: %d\n", solution)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	openings := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	var fixes [][]rune
	for _, line := range strings.Split(input, "\n") {
		stack := newRuneStack()
		var corrupt bool
		for _, char := range line {
			_, isOpening := openings[char]
			if isOpening {
				stack.push(char)
				continue
			}

			if len(stack.s) == 0 || openings[stack.pop()] != char {
				corrupt = true
				break
			}
		}

		if corrupt {
			continue
		}

		var fix []rune
		for !stack.empty() {
			if closing, isOpening := openings[stack.pop()]; isOpening {
				fix = append(fix, closing)
			}
		}
		fixes = append(fixes, fix)
	}

	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	scores := make([]int, 0, len(fixes))
	for _, fix := range fixes {
		score := 0
		for _, char := range fix {
			score *= 5
			score += points[char]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	solution := scores[len(scores)/2]
	fmt.Printf("solution: %d\n", solution)
}

// ----------------------------------------
// utils
// ----------------------------------------

type runeStack struct {
	s []rune
}

func newRuneStack() *runeStack {
	return &runeStack{
		s: make([]rune, 0),
	}
}

func (rs *runeStack) empty() bool {
	return len(rs.s) == 0
}

func (rs *runeStack) push(r rune) {
	rs.s = append(rs.s, r)
}

func (rs *runeStack) pop() rune {
	n := len(rs.s) - 1
	r := rs.s[n]
	rs.s = rs.s[:n]
	return r
}
