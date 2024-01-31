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
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

type vector3 struct {
	x, y, z int
}

func (v vector3) add(other vector3) vector3 {
	v.x += other.x
	v.y += other.y
	v.z += other.z
	return v
}

type brick struct {
	start, end vector3
}

func parse(lines []string) (bricks []brick) {
	for _, line := range lines {
		left, right, _ := strings.Cut(line, "~")
		l, r := strings.Split(left, ","), strings.Split(right, ",")
		bricks = append(bricks, brick{
			start: vector3{x: utils.ToInt(l[0]), y: utils.ToInt(l[1]), z: utils.ToInt(l[2])},
			end:   vector3{x: utils.ToInt(r[0]), y: utils.ToInt(r[1]), z: utils.ToInt(r[2])},
		})
	}
	return bricks
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	bricks := parse(lines)
	space := data.NewSet[vector3]()
	for _, b := range bricks {
		dir := vector3{x: sign(b.end.x - b.start.x), y: sign(b.end.y - b.start.y), z: sign(b.end.z - b.start.z)}
		current := b.start
		space.Add(current)
		for current != b.end {
			current = current.add(dir)
			space.Add(current)
		}
	}

	fmt.Printf("%v\n", space)

	return -1
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return -1
}

func sign(num int) int {
	if num == 0 {
		return 0
	}
	return int(math.Copysign(1, float64(num)))
}
