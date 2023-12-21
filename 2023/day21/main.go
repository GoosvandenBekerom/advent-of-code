package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/data/directions"

	"github.com/GoosvandenBekerom/advent-of-code/data"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines, 64))
	fmt.Println(part2(lines, 26501365))
}

type Grid [][]byte

func parse(lines []string) (grid Grid, start data.Vector) {
	grid = make([][]byte, len(lines))
	for y, line := range lines {
		grid[y] = []byte(line)
		if start.X != 0 {
			continue
		}
		for x, b := range grid[y] {
			if b == 'S' {
				start = data.Vector{X: x, Y: y}
				break
			}
		}
	}
	grid[start.Y][start.X] = '.'
	return grid, start
}

func solve(grid Grid, start data.Vector) map[data.Vector]int {
	visited := make(map[data.Vector]int)
	maxX, maxY := len(grid[0])-1, len(grid)-1

	type n struct {
		data.Vector
		distance int
	}

	q := data.NewQueue[n]()
	q.Enqueue(n{start, 0})
	for !q.Empty() {
		v := q.Dequeue()
		if _, done := visited[v.Vector]; done {
			continue
		}
		visited[v.Vector] = v.distance

		for _, d := range directions.AllOrthogonal {
			next := v.Add(d.Vector)
			if next.X < 0 || next.X > maxX || next.Y < 0 || next.Y > maxY {
				continue
			}
			if _, done := visited[next]; done || grid[next.Y][next.X] == '#' {
				continue
			}
			q.Enqueue(n{next, v.distance + 1})
		}
	}
	return visited
}

func part1(lines []string, steps int) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	var count int
	for _, distance := range solve(parse(lines)) {
		if distance <= steps && distance%2 == 0 {
			count++
		}
	}
	return count
}

func part2(lines []string, steps int) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	visited := solve(parse(lines))
	size := len(lines)
	half := size / 2
	var evenCorners, oddCorners, evenTotal, oddTotal int
	for _, distance := range visited {
		if distance > half {
			if distance%2 == 0 {
				evenCorners++
			} else {
				oddCorners++
			}
		}
		if distance%2 == 0 {
			evenTotal++
		} else {
			oddTotal++
		}
	}
	n := steps / size
	even := n * n
	odd := (n + 1) * (n + 1)

	return even*evenTotal + odd*oddTotal - ((n + 1) * oddCorners) + (n * evenCorners)
}
