package main

import (
	_ "embed"
	"fmt"
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

type vector struct{ x, y int }
type octopus struct {
	energy  int
	flashed bool
}

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	lines := strings.Split(input, "\n")
	grid := make([][]*octopus, len(lines))
	for y, row := range lines {
		grid[y] = make([]*octopus, len(row))
		for x, col := range row {
			grid[y][x] = &octopus{energy: int(col - 48)}
		}
	}

	var flashes int
	for i := 0; i < 100; i++ {
		for y, row := range grid {
			for x, octo := range row {
				octo.energy = (octo.energy + 1) % 10
				if octo.energy != 0 {
					continue
				}
				flash(vector{x, y}, grid)
			}
		}
		println("after step " + strconv.Itoa(i))
		for _, row := range grid {
			for _, octo := range row {
				if octo.flashed {
					octo.energy = 0
					octo.flashed = false
					flashes++
				}
				fmt.Printf("%d", octo.energy)
			}
			println()
		}
		println()
	}

	fmt.Printf("solution: %d\n", flashes)
}

func flash(p vector, grid [][]*octopus) {
	if grid[p.y][p.x].flashed {
		return
	}
	grid[p.y][p.x].flashed = true
	for _, nb := range neighbors(p.x, p.y) {
		octo := grid[nb.y][nb.x]
		octo.energy = (octo.energy + 1) % 10
		if octo.energy != 0 {
			continue
		}
		flash(nb, grid)
	}
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	lines := strings.Split(input, "\n")
	grid := make([][]*octopus, len(lines))
	for y, row := range lines {
		grid[y] = make([]*octopus, len(row))
		for x, col := range row {
			grid[y][x] = &octopus{energy: int(col - 48)}
		}
	}

	var flashes int
	var iteration int
	for {
		var flashesThisIteration int
		iteration++
		for y, row := range grid {
			for x, octo := range row {
				octo.energy = (octo.energy + 1) % 10
				if octo.energy != 0 {
					continue
				}
				flash(vector{x, y}, grid)
			}
		}
		for _, row := range grid {
			for _, octo := range row {
				if octo.flashed {
					octo.energy = 0
					octo.flashed = false
					flashes++
					flashesThisIteration++
				}
			}
		}
		if flashesThisIteration == 100 {
			break
		}
	}

	fmt.Printf("solution: %d\n", iteration)
}

// ----------------------------------------
// utils
// ----------------------------------------

func neighbors(x, y int) []vector {
	var nb []vector
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
			if ay > 9 || ay < 0 {
				break
			}
			if ax > 9 || ax < 0 {
				break
			}
			nb = append(nb, vector{
				x: ax,
				y: ay,
			})
			break
		}
	}
	return nb
}
