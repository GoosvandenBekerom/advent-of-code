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
	var sum int
	for y := miny; y <= maxy; y++ {
		var edgePositionsInRow []data.Vector
		for x := minx; x <= maxx; x++ {
			pos := data.Vector{X: x, Y: y}
			if edge.Has(pos) {
				edgePositionsInRow = append(edgePositionsInRow, pos)
			}
		}
		if len(edgePositionsInRow)%2 != 0 {
			// Filter out horizontal edges by taking first x position of edges as "start" and last as "end"
			var prev int
			var edgeStart *int
			var newEdgePositionsInRow []data.Vector
			for i := 0; i < len(edgePositionsInRow); i++ {
				if edgeStart != nil && (edgePositionsInRow[i].X-prev > 1 || i == len(edgePositionsInRow)-1) {
					newEdgePositionsInRow = append(newEdgePositionsInRow, data.Vector{X: *edgeStart, Y: y})
					newEdgePositionsInRow = append(newEdgePositionsInRow, data.Vector{X: edgePositionsInRow[i].X, Y: y})
					edgeStart = nil
					continue
				}

				prev = edgePositionsInRow[i].X
				if edgeStart == nil {
					edgeStart = &edgePositionsInRow[i].X
				}
			}
			edgePositionsInRow = newEdgePositionsInRow
		}
		for i := 0; i < len(edgePositionsInRow); i += 2 {
			sum += edgePositionsInRow[i+1].X - edgePositionsInRow[i].X + 1
		}
		//println("sum after row", y, sum)
	}
	return sum
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	instructions := parse(lines)
	path := walk(instructions)
	return calculateArea(path) // must be higher than 49582
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return -1
}
