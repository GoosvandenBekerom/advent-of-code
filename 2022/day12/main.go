package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/data"
	"math"
	"strings"

	"github.com/beefsack/go-astar"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type WorldMap map[data.Vector]*tile

type tile struct {
	char     uint8
	position data.Vector
	world    WorldMap
	visited  map[data.Vector]int
}

func (t tile) elevation() uint8 {
	switch t.char {
	case 'S':
		return 'a'
	case 'E':
		return 'z'
	default:
		return t.char
	}
}

func (t tile) PathNeighbors() []astar.Pather {
	currentElevation := t.elevation()
	var neighbors []astar.Pather
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		pos := data.Vector{
			X: t.position.X + offset[0],
			Y: t.position.Y + offset[1],
		}
		if n := t.world[pos]; n != nil {
			dest := n.elevation()
			if dest <= currentElevation || dest-currentElevation == 1 {
				// only add neighbor if it has elevation difference of 0 or 1
				neighbors = append(neighbors, n)
			}
		}
	}
	return neighbors
}

func (t tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t tile) PathEstimatedCost(to astar.Pather) float64 {
	destination := to.(*tile)
	absX := math.Abs(float64(destination.position.X - t.position.X))
	absY := math.Abs(float64(destination.position.Y - t.position.Y))
	return absX + absY
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	world := make(WorldMap)
	var start, end *tile
	for y, line := range lines {
		for x, char := range []uint8(line) {
			pos := data.Vector{X: x, Y: y}
			t := &tile{char: char, position: pos, world: world}
			world[pos] = t

			if char == 'S' {
				start = t
			} else if char == 'E' {
				end = t
			}
		}
	}
	path, _, _ := astar.Path(start, end)
	return len(path) - 1
}

func printPath(path []astar.Pather, columns int, rows int) {
	world := make(WorldMap)
	for _, item := range path {
		t := item.(*tile)
		world[t.position] = t
	}

	for y := 0; y < columns; y++ {
		for x := 0; x < rows; x++ {
			if _, ok := world[data.Vector{X: x, Y: y}]; ok {
				print(". ")
			} else {
				print("# ")
			}
		}
		println()
	}
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	world := make(WorldMap)
	var starts []*tile
	var end *tile
	for y, line := range lines {
		for x, char := range []uint8(line) {
			pos := data.Vector{X: x, Y: y}
			t := &tile{char: char, position: pos, world: world}
			world[pos] = t

			if t.elevation() == 'a' {
				starts = append(starts, t)
			} else if char == 'E' {
				end = t
			}
		}
	}

	var shortest int
	for _, start := range starts {
		path, _, _ := astar.Path(start, end)
		length := len(path) - 1
		if length != -1 && (shortest == 0 || length < shortest) {
			shortest = length
		}
	}

	return shortest
}
