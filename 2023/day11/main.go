package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/data"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type set map[int]struct{}

func (s set) add(i int) {
	s[i] = struct{}{}
}

func (s set) has(i int) bool {
	_, ok := s[i]
	return ok
}

func (s set) remove(i int) {
	delete(s, i)
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	return calculate(lines, 2)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	return calculate(lines, 1000000)
}

func calculate(lines []string, expansionSize int) int {
	emptyCols := make(set)
	for i := 0; i < len(lines[0]); i++ {
		emptyCols.add(i)
	}
	emptyRows := make(set)
	for rowI, line := range lines {
		rowEmpty := true
		for colI, char := range line {
			if char == '#' {
				rowEmpty = false
				emptyCols.remove(colI)
			}
		}
		if rowEmpty {
			emptyRows.add(rowI)
		}
	}

	galaxies := make(map[data.Vector]int)
	y := 0
	galaxyIndex := 1
	for row, line := range lines {
		x := 0
		for col, char := range line {
			if char == '#' {
				galaxies[data.Vector{X: x, Y: y}] = galaxyIndex
				galaxyIndex++
			}
			if emptyCols.has(col) {
				x += expansionSize
			} else {
				x++
			}
		}
		if emptyRows.has(row) {
			y += expansionSize
		} else {
			y++
		}
	}

	indexedGalaxies := make([]data.Vector, 0, len(galaxies))
	for galaxy := range galaxies {
		indexedGalaxies = append(indexedGalaxies, galaxy)
	}

	uniquePairs := make(map[data.Vector]struct{})
	for i := 0; i < len(indexedGalaxies); i++ {
		for j := len(indexedGalaxies) - 1; j >= 0; j-- {
			if i >= j {
				continue
			}
			uniquePairs[data.Vector{X: min(i, j), Y: max(i, j)}] = struct{}{}
		}
	}

	var sum int
	for pair := range uniquePairs {
		galaxy1, galaxy2 := indexedGalaxies[pair.X], indexedGalaxies[pair.Y]
		sum += galaxy1.ManhattanDistance(galaxy2)
	}

	return sum
}
