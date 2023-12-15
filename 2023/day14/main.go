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
	done := true
	for {
		done = true
		for y := 0; y < len(g); y++ {
			for x := 0; x < len(g[y]); x++ {
				if g[y][x] == 'O' {
					if y+dy < 0 || y+dy >= len(g) || x+dx < 0 || x+dx >= len(g[y]) {
						continue
					}

					if g[y+dy][x+dx] == '.' {
						g[y+dy][x+dx] = 'O'
						g[y][x] = '.'
						done = false
					}
				}
			}
		}

		if done {
			break
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
	grid := ByteGrid(lines)
	seen := map[string]int{grid.String(): 0}
	indexes := map[int]string{0: grid.String()}

	var i int
	for {
		i++
		grid.tilt(-1, 0) // North
		grid.tilt(0, -1) // West
		grid.tilt(1, 0)  // South
		grid.tilt(0, 1)  // East

		key := grid.String()
		firstTimeSeen, seenBefore := seen[key]
		if seenBefore {
			var iterationsSinceSeen = i - firstTimeSeen
			var doneIndex = firstTimeSeen + ((1000000000 - firstTimeSeen) % iterationsSinceSeen)
			return ByteGrid(strings.Split(strings.TrimSpace(indexes[doneIndex]), "\n")).weight()
		}
		seen[key] = i
		indexes[i] = key
	}
}
