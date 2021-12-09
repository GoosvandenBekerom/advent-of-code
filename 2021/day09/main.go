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

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))

	for y, row := range lines {
		grid[y] = make([]int, len(row))
		for x, col := range row {
			grid[y][x] = int(col - 48) // 48 == decimal for ascii "0"
		}
	}

	var solution int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			height := grid[y][x]
			if height == 9 {
				continue
			}

			vals := make([]int, 0, 5)
			vals = append(vals, height)
			vals = append(vals, findNeighbors(x, y, grid)...)

			sort.Ints(vals)

			if height == vals[0] {
				solution += height + 1
			}
		}
	}

	fmt.Printf("solution: %d\n", solution)
}

type vector struct {
	x, y int
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	lines := strings.Split(input, "\n")
	grid := make(map[vector]int)

	for y, row := range lines {
		for x, col := range row {
			grid[vector{x, y}] = int(col - 48) // 48 == decimal for ascii "0"
		}
	}

	directions := []vector{
		{-1, 0}, //left
		{0, -1}, //top
		{1, 0},  //right
		{0, 1},  //bottom
	}

	var basins []int
	for pos, height := range grid {
		var neighbors, lower int
		for _, d := range directions {
			v := pos
			v.x += d.x
			v.y += d.y
			h, ok := grid[v]
			if ok {
				neighbors++
				if height < h {
					lower++
				}
			}
		}
		if neighbors == lower {
			basins = append(basins, findBasin(grid, pos))
		}
	}

	sort.Ints(basins)
	a := basins[len(basins)-3]
	b := basins[len(basins)-2]
	c := basins[len(basins)-1]
	fmt.Printf("solution: %d\n", a*b*c)
}

func findBasin(grid map[vector]int, from vector) int {
	seen := make(map[vector]bool)
	stack := []vector{from}
	directions := []vector{
		{-1, 0}, //left
		{0, -1}, //top
		{1, 0},  //right
		{0, 1},  //bottom
	}
	var count int
	for len(stack) > 0 {
		head := stack[0]
		stack = stack[1:]
		for _, d := range directions {
			v := head
			v.x += d.x
			v.y += d.y

			if seen[v] {
				continue
			}
			seen[v] = true
			h, ok := grid[v]
			if !ok {
				continue
			}
			// 9's should be ignored in basins
			if h < 9 {
				count++
				stack = append(stack, v)
			}
		}
	}
	return count
}

// ----------------------------------------
// utils
// ----------------------------------------

func findNeighbors(x, y int, grid [][]int) (adj []int) {
	dir := []struct{ x, y int }{
		{-1, 0}, //left
		{0, -1}, //top
		{1, 0},  //right
		{0, 1},  //bottom
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
			adj = append(adj, grid[ay][ax])
			break
		}
	}
	return
}
