package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/data"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type instruction struct {
	direction data.Vector
	steps     int
}

func parseInstruction(line string) instruction {
	dir, steps, _ := strings.Cut(line, " ")
	var direction data.Vector
	switch dir {
	case "U":
		direction = data.Vector{Y: -1}
		break
	case "R":
		direction = data.Vector{X: 1}
		break
	case "D":
		direction = data.Vector{Y: 1}
		break
	case "L":
		direction = data.Vector{X: -1}
		break
	}
	return instruction{
		direction: direction,
		steps:     utils.ToInt(steps),
	}
}

func part1(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var H data.Vector
	var T data.Vector

	visited := make(map[data.Vector]struct{})
	visited[T] = struct{}{}

	for _, line := range strings.Split(input, "\n") {
		move := parseInstruction(line)
		for i := 0; i < move.steps; i++ {
			H = H.Add(move.direction)

			if math.Abs(float64(T.X-H.X)) > 1 || math.Abs(float64(T.Y-H.Y)) > 1 {
				T = data.Vector{
					X: H.X - move.direction.X,
					Y: H.Y - move.direction.Y,
				}
			}

			visited[T] = struct{}{}
		}
	}

	return len(visited)
}

func part2(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var H data.Vector
	var knots [9]data.Vector

	visited := make(map[data.Vector]struct{})
	visited[knots[len(knots)-1]] = struct{}{}

	for _, line := range strings.Split(input, "\n") {
		move := parseInstruction(line)
		for i := 0; i < move.steps; i++ {
			H = H.Add(move.direction)

			previous := H
			for i, knot := range knots {

				distanceX := math.Abs(float64(knot.X - previous.X))
				distanceY := math.Abs(float64(knot.Y - previous.Y))
				if distanceX > 1 || distanceY > 1 {
					if distanceX > 0 {
						if knot.X < previous.X {
							knots[i].X++
						} else {
							knots[i].X--
						}
					}
					if distanceY > 0 {
						if knot.Y < previous.Y {
							knots[i].Y++
						} else {
							knots[i].Y--
						}
					}
				}

				previous = knots[i]
			}

			visited[knots[len(knots)-1]] = struct{}{}
		}
	}

	return len(visited)
}

func printGrid(visited map[data.Vector]struct{}) {
	var minx, maxx, miny, maxy int
	for position := range visited {
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
			if _, ok := visited[data.Vector{X: x, Y: y}]; ok {
				print("# ")
			} else {
				print(". ")
			}
		}
		println()
	}
}
