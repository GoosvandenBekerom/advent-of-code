package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/data/directions"

	"github.com/GoosvandenBekerom/advent-of-code/utils"

	"github.com/GoosvandenBekerom/advent-of-code/data"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type instruction struct {
	direction byte
	length    int
	colour    string
}

func parse(lines []string) (instructions []instruction) {
	for _, line := range lines {
		parts := strings.Fields(line)
		instructions = append(instructions, instruction{
			direction: parts[0][0],
			length:    utils.ToInt(parts[1]),
			colour:    parts[2][1 : len(parts[2])-1],
		})
	}
	return instructions
}

func walk(instructions []instruction) data.Set[data.Vector] {
	path := data.NewSet[data.Vector]()
	cursor := data.Vector{}
	path.Add(cursor)
	for _, instr := range instructions {
		var d directions.Direction
		switch instr.direction {
		case 'U':
			d = directions.Up
		case 'R':
			d = directions.Right
		case 'D':
			d = directions.Down
		case 'L':
			d = directions.Left
		}
		for i := 0; i < instr.length; i++ {
			cursor = cursor.Add(d.Vector)
			path.Add(cursor)
		}
	}
	return path
}

func calculateArea(edge data.Set[data.Vector]) int {
	minx, maxx, miny, maxy := math.MaxInt, 0, math.MaxInt, 0
	for _, p := range edge.Values() {
		if p.X < minx {
			minx = p.X
		}
		if p.X > maxx {
			maxx = p.X
		}
		if p.Y < miny {
			miny = p.Y
		}
		if p.Y > maxy {
			maxy = p.Y
		}
	}
	var sb strings.Builder
	var sum int
	for y := miny; y <= maxy; y++ {
		inArea := false
		lastOriginal := '.'
		horizontalLine := false
		if y == 2 {
			print()
		}
		for x := minx; x <= maxx+5; x++ { // TODO: outer horizontal edges never get to inArea == false (remove +5 after fix)
			char := '.'
			if edge.Has(data.Vector{X: x, Y: y}) {
				if lastOriginal == '.' {
					inArea = !inArea
				} else {
					horizontalLine = true
				}
				char = '#'
				sum++
				lastOriginal = '#'
			} else {
				if horizontalLine {
					// x - 1 was still a horizontal line
					//
					inArea = false
					horizontalLine = false
				}
				lastOriginal = '.'
				if inArea {
					char = '#'
					sum++
				}
			}
			sb.WriteRune(char)
		}
		sb.WriteRune('\n')
	}
	println(sb.String())
	return sum
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	instructions := parse(lines)
	path := walk(instructions)
	return calculateArea(path)
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return -1
}
