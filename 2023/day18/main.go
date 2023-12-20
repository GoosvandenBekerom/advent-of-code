package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/GoosvandenBekerom/advent-of-code/data"
	"github.com/GoosvandenBekerom/advent-of-code/data/directions"
	"github.com/GoosvandenBekerom/advent-of-code/math"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type instruction struct {
	direction directions.Direction
	length    int
}

type lagoon struct {
	vertices  []data.Vector
	perimeter int
}

func toLagoon(instructions []instruction) lagoon {
	cursor := data.Vector{}
	result := lagoon{vertices: make([]data.Vector, 0)}
	for _, instr := range instructions {
		result.vertices = append(result.vertices, cursor)
		result.perimeter += instr.length
		cursor = data.Vector{
			X: cursor.X + instr.direction.X*instr.length,
			Y: cursor.Y + instr.direction.Y*instr.length,
		}
	}
	return result
}

// Area -> https://rosettacode.org/wiki/Shoelace_formula_for_polygonal_area
func (l lagoon) Area() int {
	area := 0
	for c := 0; c < len(l.vertices)-1; c++ {
		area += l.vertices[c].X*l.vertices[c+1].Y - l.vertices[c+1].X*l.vertices[c].Y
	}
	area += l.vertices[len(l.vertices)-1].X*l.vertices[0].Y - l.vertices[0].X*l.vertices[len(l.vertices)-1].Y
	area = math.Abs(area)

	return l.perimeter + (area-l.perimeter)/2 + 1
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	var instructions []instruction
	for _, line := range lines {
		parts := strings.Fields(line)
		var d directions.Direction
		switch parts[0][0] {
		case 'U':
			d = directions.Up
		case 'R':
			d = directions.Right
		case 'D':
			d = directions.Down
		case 'L':
			d = directions.Left
		}
		instructions = append(instructions, instruction{direction: d, length: utils.ToInt(parts[1])})
	}
	return toLagoon(instructions).Area()
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var instructions []instruction
	for _, line := range lines {
		parts := strings.Fields(line)
		length, err := strconv.ParseInt(parts[2][2:len(parts[2])-2], 16, strconv.IntSize)
		utils.Check(err)

		var d directions.Direction
		switch parts[2][7] {
		case '0':
			d = directions.Right
		case '1':
			d = directions.Down
		case '2':
			d = directions.Left
		case '3':
			d = directions.Up
		}

		instructions = append(instructions, instruction{direction: d, length: int(length)})
	}
	return toLagoon(instructions).Area()
}
