package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/data"
	"github.com/GoosvandenBekerom/advent-of-code/data/directions"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type state struct {
	cost             int
	position         data.Vector
	direction        directions.Direction
	stepsInDirection int
}

func (s state) LessThan(other state) bool {
	return s.cost < other.cost
}

type distKey struct {
	position         data.Vector
	direction        directions.Direction
	stepsInDirection int
}

// Dijkstra
func findShortestPath(grid [][]int, from, to data.Vector, ultra bool) int {
	dist := make(map[distKey]int)
	h := new(data.Heap[state])

	dist[distKey{position: from, direction: directions.Right, stepsInDirection: 0}] = 0
	dist[distKey{position: from, direction: directions.Down, stepsInDirection: 0}] = 0
	heap.Push(h, state{cost: 0, position: from, direction: directions.Right, stepsInDirection: 0})

	for {
		s := heap.Pop(h).(state)
		if s.position == to {
			if !ultra || s.stepsInDirection >= 4 {
				return s.cost
			}
		}

		k := distKey{position: s.position, direction: s.direction, stepsInDirection: s.stepsInDirection}
		if v, exists := dist[k]; exists && s.cost > v {
			continue
		}

		for _, dir := range directions.AllOrthogonal {
			if dir == s.direction.Opposite() {
				continue
			}
			pos := s.position.Add(dir.Vector)
			if pos.X < 0 || pos.Y < 0 || pos.X >= len(grid[0]) || pos.Y >= len(grid) {
				continue
			}

			steps := 1
			if dir == s.direction {
				steps = s.stepsInDirection + 1
			}

			next := state{cost: s.cost + grid[pos.Y][pos.X], position: pos, direction: dir, stepsInDirection: steps}
			nextKey := distKey{position: pos, direction: dir, stepsInDirection: steps}

			oldCost, exists := dist[nextKey]
			if ultra {
				if (dir == s.direction || s.stepsInDirection >= 4) && steps <= 10 && (!exists || next.cost < oldCost) {
					heap.Push(h, next)
					dist[nextKey] = next.cost
				}
			} else {
				if steps <= 3 && (!exists || next.cost < oldCost) {
					heap.Push(h, next)
					dist[nextKey] = next.cost
				}
			}
		}
	}
}

func parseGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, val := range line {
			grid[y][x] = utils.ToInt(string(val))
		}
	}
	return grid
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	grid := parseGrid(lines)
	return findShortestPath(grid, data.Vector{}, data.Vector{X: len(grid[0]) - 1, Y: len(grid) - 1}, false)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	grid := parseGrid(lines)
	return findShortestPath(grid, data.Vector{}, data.Vector{X: len(grid[0]) - 1, Y: len(grid) - 1}, true)
}
