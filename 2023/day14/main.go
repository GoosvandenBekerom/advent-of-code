package main

import (
	"bytes"
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

type Grid [][]byte

func ByteGrid(lines []string) (g Grid) {
	for _, line := range lines {
		g = append(g, []byte(line))
	}
	return g
}

func (g Grid) String() string {
	var builder strings.Builder
	for y := 0; y < len(g); y++ {
		builder.Write(g[y])
		builder.WriteByte('\n')
	}
	return builder.String()
}

func (g Grid) tilt(dy, dx int) {
	for x := 0; x < len(g[0]); x++ {
		inEmptySpace := false
		lastEmpty := 0
		for y := 0; y < len(g); y++ {
			switch g[y][x] {
			case '#':
				inEmptySpace = false
			case 'O':
				if !inEmptySpace {
					break
				}
				g[lastEmpty][x] = 'O'
				g[y][x] = '.'
				lastEmpty++
			case '.':
				if !inEmptySpace {
					lastEmpty = y
					inEmptySpace = true
				}
			}
		}
	}
}

func (g Grid) weight() int {
	var weight int
	multiplier := len(g)
	for _, row := range g {
		weight += multiplier * bytes.Count(row, []byte("O"))
		multiplier--
	}
	return weight
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	grid := ByteGrid(lines)
	grid.tilt(-1, 0)
	return grid.weight()
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return -1
}
