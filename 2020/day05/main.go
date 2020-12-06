package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

const (
	rows = 127
	cols = 7
)

func main() {
	passes := input()
	fmt.Printf("part 1: %d\n", part1(passes))
	fmt.Printf("part 2: %d\n", part2(passes))
}

func part1(passes []string) int {
	seats := parseAndSortDesc(passes)
	return seats[0].ID()
}

func part2(passes []string) int {
	seats := parseAndSortDesc(passes)
	prev := seats[0].ID() + 1
	for _, seat := range seats {
		if prev-seat.ID() == 2 {
			return prev - 1
		}
		prev = seat.ID()
	}
	return -1
}

func parseAndSortDesc(passes []string) []seat {
	seats := make([]seat, len(passes))
	for i, pass := range passes {
		seats[i] = parse(pass)
	}
	sort.Sort(byIDdesc(seats))
	return seats
}

func parse(pass string) seat {
	low, high := 0, rows
	left, right := 0, cols
	for _, c := range pass[:7] {
		switch c {
		case 'F':
			high = low + ((high - low) / 2)
		case 'B':
			low = low + ((high - low) / 2) + 1
		}
	}
	for _, c := range pass[len(pass)-3:] {
		switch c {
		case 'L':
			right = left + ((right - left) / 2)
		case 'R':
			left = left + ((right - left) / 2) + 1
		}
	}
	return seat{pass: pass, row: low, column: left}
}

// ----------------------------------------
// domain
// ----------------------------------------

type seat struct {
	pass        string
	row, column int
	_id         int
}

func (s seat) ID() int {
	if s._id == 0 {
		s._id = (s.row * 8) + s.column
	}
	return s._id
}

func (s seat) String() string {
	return fmt.Sprintf("%s - row: %d, column: %d, ID: %d", s.pass, s.row, s.column, s.ID())
}

type byIDdesc []seat

func (a byIDdesc) Len() int           { return len(a) }
func (a byIDdesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIDdesc) Less(i, j int) bool { return a[i].ID() > a[j].ID() }

// ----------------------------------------
// utils
// ----------------------------------------

func input() []string {
	box := packr.New("day05", "./2020/day05")
	s, err := box.FindString("input")
	check(err)
	return strings.Split(strings.TrimSuffix(s, "\n"), "\n")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
