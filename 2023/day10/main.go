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

var pipeDirections = map[rune][]data.Vector{
	'S': {{X: -1}, {X: 1}, {Y: -1}, {Y: 1}},
	'-': {{X: -1}, {X: 1}},
	'|': {{Y: -1}, {Y: 1}},
	'F': {{X: 1}, {Y: 1}},
	'7': {{X: -1}, {Y: 1}},
	'L': {{X: 1}, {Y: -1}},
	'J': {{X: -1}, {Y: -1}},
}

func canConnect(nextPipe rune, lastPos, nextPos data.Vector) bool {
	possibleSteps := pipeDirections[nextPipe]
	for _, step := range possibleSteps {
		if nextPos.Add(step) == lastPos {
			return true
		}
	}
	return false
}

func findPath(pipes map[data.Vector]rune, currentPath []data.Vector) (path []data.Vector) {
	lastPos := currentPath[len(currentPath)-1]
	last := pipes[lastPos]
	var beforeLastPos data.Vector
	if len(currentPath) > 1 {
		beforeLastPos = currentPath[len(currentPath)-2]
	}

	if last == 'S' && len(currentPath) != 1 {
		return currentPath
	}

	for _, direction := range pipeDirections[last] {
		nextPos := lastPos.Add(direction)
		if nextPos == beforeLastPos {
			continue
		}
		next, exists := pipes[nextPos]
		if !exists {
			continue
		}
		if canConnect(next, lastPos, nextPos) {
			return findPath(pipes, append(currentPath, nextPos))
		}
	}

	return nil
}

func part1(grid []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	var start data.Vector
	pipes := make(map[data.Vector]rune)
	for y, chars := range grid {
		for x, char := range chars {
			switch char {
			case '.':
				continue
			case 'S':
				start = data.Vector{X: x, Y: y}
				fallthrough
			default:
				pipes[data.Vector{X: x, Y: y}] = char
			}
		}
	}

	return len(findPath(pipes, []data.Vector{start})) / 2
}

func isEnclosed(pipePath map[data.Vector]struct{}, origin data.Vector, lowerX, lowerY, upperX, upperY int) bool {
	directions := []data.Vector{
		{-1, 0}, //left
		{0, -1}, //top
		{1, 0},  //right
		{0, 1},  //bottom
	}
	for _, stepInDirection := range directions {
		current := origin
		for {
			current = current.Add(stepInDirection)
			if _, isPipe := pipePath[current]; isPipe {
				break
			}
			if current.X < lowerX || current.Y < lowerY || current.X > upperX || current.Y > upperY {
				return false
			}
		}
	}
	return true
}

func part2(grid []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")
	var start data.Vector
	pipes := make(map[data.Vector]rune)
	for y, chars := range grid {
		for x, char := range chars {
			switch char {
			case '.':
				continue
			case 'S':
				start = data.Vector{X: x, Y: y}
				fallthrough
			default:
				pipes[data.Vector{X: x, Y: y}] = char
			}
		}
	}

	pipePath := make(map[data.Vector]struct{})
	lowerX, upperX, lowerY, upperY := len(grid[0]), 0, len(grid), 0
	for _, pipe := range findPath(pipes, []data.Vector{start}) {
		pipePath[pipe] = struct{}{}
		if pipe.X < lowerX {
			lowerX = pipe.X
		} else if pipe.X > upperX {
			upperX = pipe.X
		}
		if pipe.Y < lowerY {
			lowerY = pipe.Y
		} else if pipe.Y > upperY {
			upperY = pipe.Y
		}
	}

	var sum int
	for y, chars := range grid {
		for x := range chars {
			pos := data.Vector{X: x, Y: y}
			if _, isPipe := pipePath[pos]; isPipe {
				continue
			}
			if pos.X < lowerX || pos.Y < lowerY || pos.X > upperX || pos.Y > upperY {
				continue
			}
			if isEnclosed(pipePath, pos, lowerX, lowerY, upperX, upperY) {
				sum++
			}
		}
	}

	return sum
}

func sanitizeInput(grid []string) {
	var start data.Vector
	pipes := make(map[data.Vector]rune)
	for y, chars := range grid {
		for x, char := range chars {
			switch char {
			case '.':
				continue
			case 'S':
				start = data.Vector{X: x, Y: y}
				fallthrough
			default:
				pipes[data.Vector{X: x, Y: y}] = char
			}
		}
	}

	pipePath := make(map[data.Vector]struct{})
	for _, pipe := range findPath(pipes, []data.Vector{start}) {
		pipePath[pipe] = struct{}{}
	}

	for y, chars := range grid {
		for x, char := range chars {
			if _, isPipe := pipePath[data.Vector{X: x, Y: y}]; isPipe {
				print(string(char))
			} else {
				print(".")
			}
		}
		println()
	}
}
