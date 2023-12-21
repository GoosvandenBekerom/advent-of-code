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
	fmt.Println(part1(lines, 64))
	fmt.Println(part2(lines))
}

type Grid [][]byte

func (g Grid) print(unique map[data.Vector]byte) {
	var sb strings.Builder
	for y, bytes := range g {
		for x, b := range bytes {
			if _, exists := unique[data.Vector{X: x, Y: y}]; exists {
				sb.WriteRune('O')
			} else {
				sb.WriteRune(rune(b))
			}
		}
		sb.WriteRune('\n')
	}
	println(sb.String())
}

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
	return grid, start
}

func solve(grid Grid, start data.Vector, steps int) int {
	visited := make(map[data.Vector]byte)
	maxX, maxY := len(grid[0])-1, len(grid)-1

	type n struct {
		data.Vector
		steps int
	}

	q := data.NewQueue[n]()
	q.Enqueue(n{Vector: start})

	for !q.Empty() {
		v := q.Dequeue()
		if v.steps > steps {
			continue
		}
		if _, alreadyVisited := visited[v.Vector]; alreadyVisited {
			continue
		}

		char := grid[v.Y][v.X]
		visited[v.Vector] = char

		if char == '#' {
			continue
		}

		q.Enqueue(n{Vector: data.Vector{X: min(v.X+1, maxX), Y: v.Y}, steps: v.steps + 1})
		q.Enqueue(n{Vector: data.Vector{X: max(v.X-1, 0), Y: v.Y}, steps: v.steps + 1})
		q.Enqueue(n{Vector: data.Vector{X: v.X, Y: min(v.Y+1, maxY)}, steps: v.steps + 1})
		q.Enqueue(n{Vector: data.Vector{X: v.X, Y: max(v.Y-1, 0)}, steps: v.steps + 1})
	}

	//grid.print(visited)
	count := 1 // include start
	for p, char := range visited {
		if (p.X+p.Y)%2 != (start.X+start.Y)%2 {
			continue
		}
		if char == '.' {
			count++
		}
	}
	return count
}

func part1(lines []string, steps int) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	grid, start := parse(lines)
	return solve(grid, start, steps)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return -1
}
