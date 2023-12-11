package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/data"
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
	return calculate(lines, 2)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	return calculate(lines, 1000000)
}

func calculate(lines []string, expansionSize int) int {
	emptyCols := data.NewSet[int]()
	for i := 0; i < len(lines[0]); i++ {
		emptyCols.Add(i)
	}
	emptyRows := data.NewSet[int]()
	for rowI, line := range lines {
		rowEmpty := true
		for colI, char := range line {
			if char == '#' {
				rowEmpty = false
				emptyCols.Remove(colI)
			}
		}
		if rowEmpty {
			emptyRows.Add(rowI)
		}
	}

	galaxies := data.NewSet[data.Vector]()
	y := 0
	for row, line := range lines {
		x := 0
		for col, char := range line {
			if char == '#' {
				galaxies.Add(data.Vector{X: x, Y: y})
			}
			if emptyCols.Has(col) {
				x += expansionSize
			} else {
				x++
			}
		}
		if emptyRows.Has(row) {
			y += expansionSize
		} else {
			y++
		}
	}

	indexedGalaxies := make([]data.Vector, 0, galaxies.Len())
	for _, galaxy := range galaxies.Values() {
		indexedGalaxies = append(indexedGalaxies, galaxy)
	}

	uniquePairs := data.NewSet[data.Vector]()
	for i := 0; i < len(indexedGalaxies); i++ {
		for j := len(indexedGalaxies) - 1; j >= 0; j-- {
			if i >= j {
				continue
			}
			uniquePairs.Add(data.Vector{X: min(i, j), Y: max(i, j)})
		}
	}

	var sum int
	for _, pair := range uniquePairs.Values() {
		galaxy1, galaxy2 := indexedGalaxies[pair.X], indexedGalaxies[pair.Y]
		sum += galaxy1.ManhattanDistance(galaxy2)
	}

	return sum
}
