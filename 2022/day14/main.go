package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/data"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
	"math"
	"strings"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func simulateSand(obstacles map[data.Vector]struct{}, highestY int, hasFloor bool) int {
	var units int
	var done bool
	for /* Sand generation loop */ {
		p := data.Vector{X: 500}
		if _, blocked := obstacles[p]; blocked {
			break
		}

		for {
			nextY := p.Y + 1
			if nextY == highestY+2 {
				break
			}
			if _, blocked := obstacles[data.Vector{X: p.X, Y: nextY}]; blocked {
				if _, blocked := obstacles[data.Vector{X: p.X - 1, Y: nextY}]; blocked {
					if _, blocked := obstacles[data.Vector{X: p.X + 1, Y: nextY}]; blocked {
						break
					} else {
						p.X++
					}
				} else {
					p.X--
				}
			}
			p.Y++
			if !hasFloor && p.Y > highestY {
				done = true
				break
			}
		}
		if done {
			break
		}
		obstacles[p] = struct{}{}
		units++
	}

	return units
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	obstacles, highestY := parseInput(lines)
	return simulateSand(obstacles, highestY, false)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	obstacles, highestY := parseInput(lines)
	return simulateSand(obstacles, highestY, true)
}

func parseInput(lines []string) (map[data.Vector]struct{}, int) {
	verticalSlice := make(map[data.Vector]struct{})
	var highestY int
	for _, line := range lines {
		var previous *data.Vector
		for _, raw := range strings.Split(line, " -> ") {
			x, y, _ := strings.Cut(raw, ",")
			coords := data.Vector{X: utils.ToInt(x), Y: utils.ToInt(y)}

			if coords.Y > highestY {
				highestY = coords.Y
			}

			if previous == nil {
				verticalSlice[coords] = struct{}{}
			} else {
				// add every position between previous and coords
				if previous.X == coords.X {
					var lowest, highest int
					if previous.Y < coords.Y {
						lowest = previous.Y
						highest = coords.Y
					} else {
						lowest = coords.Y
						highest = previous.Y
					}
					for y := lowest; y <= highest; y++ {
						verticalSlice[data.Vector{X: coords.X, Y: y}] = struct{}{}
					}
				} else {
					var lowest, highest int
					if previous.X < coords.X {
						lowest = previous.X
						highest = coords.X
					} else {
						lowest = coords.X
						highest = previous.X
					}
					for x := lowest; x <= highest; x++ {
						verticalSlice[data.Vector{X: x, Y: coords.Y}] = struct{}{}
					}
				}
			}

			previous = &coords
		}
	}
	return verticalSlice, highestY
}

func printSlice(rocks map[data.Vector]struct{}) {
	var minx, maxx, miny, maxy int
	minx = math.MaxInt
	for position := range rocks {
		if position.X < minx {
			minx = position.X
		}
		if position.X > maxx {
			maxx = position.X
		}
		if position.Y < miny {
			miny = position.Y
		}
		if position.Y > maxy {
			maxy = position.Y
		}
	}

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if _, ok := rocks[data.Vector{X: x, Y: y}]; ok {
				print("# ")
			} else {
				print(". ")
			}
		}
		println()
	}
}
