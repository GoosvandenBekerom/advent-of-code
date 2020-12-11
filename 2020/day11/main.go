package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	//fmt.Printf("part 1: %d\n", part1())
	fmt.Printf("part 2: %d\n", part2())
}

/*
rules:
If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
*/
func part1() int {
	grid := input()
	for {
		cp := cp(grid)
		changed := false
		for y, row := range cp {
			for x, spot := range row {
				if spot == FLOOR {
					continue
				}
				adjacent := findAdjacent(x, y, cp)
				noneTaken := true
				adjTaken := 0
				for _, adj := range adjacent {
					if adj == TAKEN {
						noneTaken = false
						adjTaken++
					}
				}
				if noneTaken && spot == EMPTY {
					grid[y][x] = TAKEN
					changed = true
				}
				if adjTaken >= 4 && spot == TAKEN {
					grid[y][x] = EMPTY
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}
	return countOccupied(grid)
}

func countOccupied(grid [][]spot) (count int) {
	for _, row := range grid {
		for _, spot := range row {
			if spot == TAKEN {
				count++
			}
		}
	}
	return
}

func printGrid(grid [][]spot) {
	for _, row := range grid {
		for _, spot := range row {
			c := '_'
			switch spot {
			case EMPTY:
				c = 'L'
			case TAKEN:
				c = '#'
			case FLOOR:
				c = '.'
			}
			fmt.Printf("%s ", string(c))
		}
		println()
	}
	println()
}

func part2() int {
	grid := input()
	for {
		cp := cp(grid)
		changed := false
		for y, row := range cp {
			for x, spot := range row {
				if spot == FLOOR {
					continue
				}
				adjacent := findAdjacent(x, y, cp)
				noneTaken := true
				adjTaken := 0
				for _, adj := range adjacent {
					if adj == TAKEN {
						noneTaken = false
						adjTaken++
					}
				}
				if noneTaken && spot == EMPTY {
					grid[y][x] = TAKEN
					changed = true
				}
				if adjTaken >= 5 && spot == TAKEN {
					grid[y][x] = EMPTY
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}
	return countOccupied(grid)
}

func findAdjacent(x, y int, grid [][]spot) (adj []spot) {
	dir := []struct{ x, y int }{
		{-1, 0},  //left
		{-1, -1}, //topleft
		{0, -1},  //top
		{1, -1},  //topright
		{1, 0},   //right
		{1, 1},   //bottomright
		{0, 1},   //bottom
		{-1, 1},  //bottomleft
	}
	for _, d := range dir {
		ax := x + d.x
		ay := y + d.y
		for {
			if ay >= len(grid) || ay < 0 {
				break
			}
			if ax >= len(grid[0]) || ax < 0 {
				break
			}
			next := grid[ay][ax]
			if next == FLOOR {
				ax += d.x
				ay += d.y
				continue
			}
			adj = append(adj, next)
			break
		}
	}
	return
}

// ----------------------------------------
// domain
// ----------------------------------------

type spot int

const (
	FLOOR spot = iota
	EMPTY spot = iota
	TAKEN spot = iota
)

func spotFrom(c rune) spot {
	switch c {
	case '.':
		return FLOOR
	case 'L':
		return EMPTY
	case '#':
		return TAKEN
	default:
		panic("got unknown character: " + string(c))
	}
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (grid [][]spot) {
	box := packr.New("day11", "./2020/day11")
	s, err := box.FindString("input")
	check(err)
	for y, line := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		grid = append(grid, make([]spot, len(line)))
		for x, char := range line {
			grid[y][x] = spotFrom(char)
		}
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func cp(grid [][]spot) (out [][]spot) {
	out = make([][]spot, len(grid))
	for y, vy := range grid {
		out[y] = make([]spot, len(vy))
		for x, vx := range vy {
			out[y][x] = vx
		}
	}
	return
}
