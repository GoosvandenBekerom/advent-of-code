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
		if y != len(g)-1 {
			builder.WriteByte('\n')
		}
	}
	return builder.String()
}

const (
	empty              byte = '.'
	mirrorUp           byte = '/'
	mirrorDown         byte = '\\'
	splitterVertical   byte = '|'
	splitterHorizontal byte = '-'
)

type beam struct {
	path []data.Vector
	dir  data.Vector
	done bool
}

func (b beam) key() string {
	return fmt.Sprintf("dir:x%dy%d,path:%v", b.dir.X, b.dir.Y, b.path)
}

func calculateEnergy(grid [][]byte, origins []beam) int {
	var highest int
	for _, o := range origins {
		origin := o
		beams := []*beam{
			&origin,
		}
		originCache := data.NewSet[string]()

		startNewBeam := func(origin, direction data.Vector) {
			b := &beam{
				path: []data.Vector{origin},
				dir:  direction,
			}
			key := b.key()
			if originCache.Has(key) {
				return
			}
			originCache.Add(key)
			beams = append(beams, b)
		}

		for {
			for _, b := range beams {
				if b.done {
					continue
				}
				lastPos := b.path[len(b.path)-1]
				nextPos := lastPos.Add(b.dir)
				// check bounds
				if nextPos.X < 0 || nextPos.X >= len(grid[0]) || nextPos.Y < 0 || nextPos.Y >= len(grid) {
					b.done = true
					continue
				}
				// move beam to next position
				b.path = append(b.path, nextPos)
				// see if beam needs to change direction or split
				switch grid[nextPos.Y][nextPos.X] {
				case mirrorUp:
					// swap dir x and y negative
					b.dir = data.Vector{X: -b.dir.Y, Y: -b.dir.X}
				case mirrorDown:
					// swap dir x and y
					b.dir = data.Vector{X: b.dir.Y, Y: b.dir.X}
				case splitterVertical:
					if b.dir.X != 0 {
						b.done = true
						// create new upward/downward beams from next position
						// might be a bug if amount of beams per pos becomes important in part 2
						startNewBeam(nextPos, data.Vector{X: 0, Y: 1})
						startNewBeam(nextPos, data.Vector{X: 0, Y: -1})
					}
				case splitterHorizontal:
					if b.dir.Y != 0 {
						b.done = true
						// create new rightward/leftward beams from next position
						// might be a bug if amount of beams per pos becomes important in part 2
						startNewBeam(nextPos, data.Vector{X: 1, Y: 0})
						startNewBeam(nextPos, data.Vector{X: -1, Y: 0})
					}
				}
			}

			finished := true
			for _, b := range beams {
				if !b.done {
					finished = false
				}
			}
			if finished {
				break
			}
		}

		energized := data.NewSet[data.Vector]()
		for _, b := range beams {
			for _, pos := range b.path {
				energized.Add(pos)
			}
		}

		if energized.Len() > highest {
			highest = energized.Len()
		}
	}

	return highest
}

func printEnergy(lx, ly int, energized data.Set[data.Vector]) {
	for y := 0; y < ly; y++ {
		for x := 0; x < lx; x++ {
			if energized.Has(data.Vector{X: x, Y: y}) {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	grid := ByteGrid(lines)
	return calculateEnergy(grid, []beam{
		{path: []data.Vector{{X: 0, Y: 0}}, dir: data.Vector{X: 1}},
	})
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	grid := ByteGrid(lines)
	var origins []beam
	// top row
	for x := range grid[0] {
		origins = append(origins, beam{path: []data.Vector{{X: x}}, dir: data.Vector{Y: 1}})
	}
	// bottom row
	bottomY := len(grid) - 1
	for x := range grid[bottomY] {
		origins = append(origins, beam{path: []data.Vector{{X: x, Y: bottomY}}, dir: data.Vector{Y: -1}})
	}
	// left column
	for y := 0; y < len(grid); y++ {
		origins = append(origins, beam{path: []data.Vector{{Y: y}}, dir: data.Vector{X: 1}})
	}
	// right column
	for y := 0; y < len(grid); y++ {
		origins = append(origins, beam{path: []data.Vector{{X: len(grid[0]) - 1, Y: y}}, dir: data.Vector{X: -1}})
	}
	return calculateEnergy(grid, origins)
}
