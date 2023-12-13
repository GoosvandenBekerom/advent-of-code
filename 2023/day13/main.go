package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	var sum int
	for _, grid := range toGrid(lines) {
		if col := grid.FindVerticalMirror(); col > 0 {
			sum += col
		} else if row := grid.FindHorizontalMirror(); row > 0 {
			sum += 100 * row
		}
	}

	return sum
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	var sum int
	for _, grid := range toGrid(lines) {
		for _, smudged := range grid.Smudges() {
			smudged.OriginalColumn = grid.FindVerticalMirror()
			if col := smudged.FindVerticalMirror(); col > 0 {
				sum += col
				break
			}

			smudged.OriginalRow = grid.FindHorizontalMirror()
			if row := smudged.FindHorizontalMirror(); row > 0 {
				sum += 100 * row
				break
			}
		}
	}

	return sum
}

type Grid struct {
	Positions      [][]byte
	OriginalRow    int
	OriginalColumn int
}

func toGrid(lines []string) []Grid {
	result := make([]Grid, 0)
	grid := Grid{
		Positions: make([][]byte, 0),
	}
	for _, line := range lines {
		if len(line) == 0 {
			result = append(result, grid)
			grid = Grid{
				Positions: make([][]byte, 0),
			}
			continue
		}

		grid.Positions = append(grid.Positions, []byte(line))
	}
	return append(result, grid)
}

func (g Grid) isMirrorColumn(a, b int) bool {
	if a < 0 || b >= len(g.Positions[0]) {
		return true
	}

	for y := 0; y < len(g.Positions); y++ {
		if g.Positions[y][a] != g.Positions[y][b] {
			return false
		}
	}

	return g.isMirrorColumn(a-1, b+1)
}

func (g Grid) isMirrorRow(a, b int) bool {
	if a < 0 || b >= len(g.Positions) {
		return true
	}

	for x := 0; x < len(g.Positions[0]); x++ {
		if g.Positions[a][x] != g.Positions[b][x] {
			return false
		}
	}

	return g.isMirrorRow(a-1, b+1)
}

func (g Grid) FindVerticalMirror() int {
	for x := 0; x < len(g.Positions[0])-1; x++ {
		if x+1 == g.OriginalColumn {
			// skip original
			continue
		}
		if g.isMirrorColumn(x, x+1) {
			return x + 1
		}
	}
	return 0
}

func (g Grid) FindHorizontalMirror() int {
	for y := 0; y < len(g.Positions)-1; y++ {
		if y+1 == g.OriginalRow {
			// skip original
			continue
		}
		if g.isMirrorRow(y, y+1) {
			return y + 1
		}
	}
	return 0
}

func (g Grid) Smudges() []Grid {
	smudgedGrids := make([]Grid, 0, len(g.Positions[0])*len(g.Positions))

	for y := 0; y < len(g.Positions); y++ {
		for x := 0; x < len(g.Positions[0]); x++ {
			smudged := g.DeepCopy()
			if g.Positions[y][x] == '#' {
				smudged.Positions[y][x] = '.'
			} else {
				smudged.Positions[y][x] = '#'
			}
			smudgedGrids = append(smudgedGrids, smudged)
		}
	}

	return smudgedGrids
}

func (g Grid) DeepCopy() Grid {
	result := Grid{
		Positions: make([][]byte, len(g.Positions)),
	}
	for y := 0; y < len(g.Positions); y++ {
		result.Positions[y] = make([]byte, len(g.Positions[y]))
		copy(result.Positions[y], g.Positions[y])
	}
	return result
}
