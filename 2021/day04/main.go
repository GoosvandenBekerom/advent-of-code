package main

import (
	_ "embed"
	"regexp"

	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
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

type box struct {
	num    int
	marked bool
}

type board struct {
	unmarked map[int]struct{}
	rows     [5][5]*box
	cols     [5][5]*box

	isDone bool
}

func (b board) print() {
	for _, row := range b.rows {
		for _, b := range row {
			if b.marked {
				fmt.Printf("X\t")
			} else {
				fmt.Printf("%d\t", b.num)
			}
		}
		println()
	}
}

func (b board) isWinner() bool {
	if len(b.unmarked) > 20 {
		return false
	}
	// check rows
	for _, row := range b.rows {
		bingo := true
		for _, b := range row {
			if !b.marked {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}
	// check cols
	for _, col := range b.cols {
		bingo := true
		for _, b := range col {
			if !b.marked {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}
	return false
}

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	parts := strings.Split(normalizeNewlines(input), "\n\n")

	// prepare the game
	var numbers []int
	for _, s := range strings.Split(parts[0], ",") {
		numbers = append(numbers, toInt(s))
	}

	fmt.Printf("%#v\n\n", numbers)

	pattern := regexp.MustCompile("\\d+")

	var boards []board
	for i := 1; i < len(parts); i++ {
		b := board{
			unmarked: make(map[int]struct{}),
		}
		for i, s := range pattern.FindAllString(parts[i], -1) {
			num := toInt(s)
			x := int(i / 5)
			y := i % 5
			b.unmarked[num] = struct{}{}
			b.rows[x][y] = &box{num: num}
			b.cols[y][x] = &box{num: num}
		}
		boards = append(boards, b)
	}

	var last int
	var winner *board = nil

	// play the game, start by pulling a number
	for index, num := range numbers {
		last = num
		// process number on all boards
		for _, b := range boards {
			if _, found := b.unmarked[num]; !found {
				continue
			}
			for i := 0; i < 25; i++ {
				x := int(i / 5)
				y := i % 5
				if b.rows[x][y].num == num {
					b.rows[x][y].marked = true
					b.cols[y][x].marked = true
					break
				}
			}
			delete(b.unmarked, num)
		}

		// cant have bingo before 5 numbers have been processed
		if index < 5 {
			continue
		}

		// check if a board has won bingo
		for _, b := range boards {
			if b.isWinner() {
				winner = &b
				break
			}
		}

		if winner != nil {
			break
		}
	}

	var sum int
	for num := range winner.unmarked {
		sum += num
	}

	fmt.Printf("solution: %v\n", sum*last)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	parts := strings.Split(normalizeNewlines(input), "\n\n")

	// prepare the game
	var numbers []int
	for _, s := range strings.Split(parts[0], ",") {
		numbers = append(numbers, toInt(s))
	}

	fmt.Printf("%#v\n\n", numbers)

	pattern := regexp.MustCompile("\\d+")

	var boards []board
	for i := 1; i < len(parts); i++ {
		b := board{
			unmarked: make(map[int]struct{}),
		}
		for i, s := range pattern.FindAllString(parts[i], -1) {
			num := toInt(s)
			x := int(i / 5)
			y := i % 5
			b.unmarked[num] = struct{}{}
			b.rows[x][y] = &box{num: num}
			b.cols[y][x] = &box{num: num}
		}
		boards = append(boards, b)
	}

	leftover := make(map[int]struct{})
	for i := range boards {
		leftover[i] = struct{}{}
	}

	var last int
	var loser *board = nil

	// play the game, start by pulling a number
	for index, num := range numbers {
		last = num
		// process number on all boards
		for _, b := range boards {
			if b.isDone {
				continue
			}
			if _, found := b.unmarked[num]; !found {
				continue
			}
			for i := 0; i < 25; i++ {
				x := int(i / 5)
				y := i % 5
				if b.rows[x][y].num == num {
					b.rows[x][y].marked = true
					b.cols[y][x].marked = true
					break
				}
			}
			delete(b.unmarked, num)
		}

		// cant have bingo before 5 numbers have been processed
		if index < 5 {
			continue
		}

		// check if a board has won bingo
		for i, b := range boards {
			if b.isDone {
				continue
			}
			if b.isWinner() {
				b.isDone = true
				if len(leftover) == 1 {
					for loserIndex := range leftover {
						loser = &boards[loserIndex]
						break
					}
				}
				delete(leftover, i)
			}
		}

		if len(leftover) == 0 {
			break
		}
	}

	var sum int
	for num := range loser.unmarked {
		sum += num
	}

	println()
	loser.print()
	fmt.Printf("last: %v\n", last)
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("solution: %v\n", sum*last)
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

func normalizeNewlines(s string) string {
	d := []byte(s)
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return string(d)
}
