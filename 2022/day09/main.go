package main

import (
	_ "embed"
	"github.com/GoosvandenBekerom/advent-of-code/datastructures"
	"math"
	"strconv"

	"fmt"
	"log"
	"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

// ----------------------------------------
// solution
// ----------------------------------------

type instruction struct {
	direction datastructures.Vector
	steps     int
}

func parseInstruction(line string) instruction {
	dir, steps, _ := strings.Cut(line, " ")
	var direction datastructures.Vector
	switch dir {
	case "U":
		direction = datastructures.Vector{Y: -1}
		break
	case "R":
		direction = datastructures.Vector{X: 1}
		break
	case "D":
		direction = datastructures.Vector{Y: 1}
		break
	case "L":
		direction = datastructures.Vector{X: -1}
		break
	}
	return instruction{
		direction: direction,
		steps:     toInt(steps),
	}
}

func part1(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var H datastructures.Vector
	var T datastructures.Vector

	visited := make(map[datastructures.Vector]struct{})
	visited[T] = struct{}{}

	for _, line := range strings.Split(input, "\n") {
		move := parseInstruction(line)
		for i := 0; i < move.steps; i++ {
			H = H.Add(move.direction)

			if math.Abs(float64(T.X-H.X)) > 1 || math.Abs(float64(T.Y-H.Y)) > 1 {
				T = datastructures.Vector{
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

	var H datastructures.Vector
	var knots [9]datastructures.Vector

	visited := make(map[datastructures.Vector]struct{})
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

func printGrid(visited map[datastructures.Vector]struct{}) {
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
			if _, ok := visited[datastructures.Vector{X: x, Y: y}]; ok {
				print("# ")
			} else {
				print(". ")
			}
		}
		println()
	}
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	check(err)
	return v
}
