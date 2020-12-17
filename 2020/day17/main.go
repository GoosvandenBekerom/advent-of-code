package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	println(part1())
	println(part2())
}

/*
During a cycle, all cubes simultaneously change their state according to the following rules:

If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
*/
func part1() int {
	space := input()
	for i := 0; i < 6; i++ {
		space.step()
		fmt.Printf("\nafter %d cycles:\n", i)
		space.print()
	}
	return space.countActive()
}

func part2() int {
	space := input()
	space.Wenabled = true
	for i := 0; i < 6; i++ {
		space.step()
	}
	return space.countActive()
}

type position struct {
	x, y, z, w int
}

type dimension struct {
	cubes map[position]bool
	minX  int
	maxX  int
	minY  int
	maxY  int
	minZ  int
	maxZ  int

	Wenabled bool
	minW     int
	maxW     int
}

func (d *dimension) set(x, y, z, w int, active bool) {
	d.cubes[position{x, y, z, w}] = active
}

func (d *dimension) isActive(x, y, z, w int) bool {
	if active, ok := d.cubes[position{x, y, z, w}]; ok {
		return active
	}
	return false
}

func (d *dimension) countNeighbors(x, y, z, w int) (count int) {
	check := func(dx, dy, dz, dw int) {
		if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
			return
		}
		if d.isActive(x+dx, y+dy, z+dz, w+dw) {
			count++
		}
	}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if d.Wenabled {
					for dw := -1; dw <= 1; dw++ {
						check(dx, dy, dz, dw)
					}
					continue
				}
				check(dx, dy, dz, 0)
			}
		}
	}
	return
}

func (d *dimension) step() {
	work := &dimension{cubes: map[position]bool{}}

	stepCube := func(x, y, z, w int) {
		active := d.isActive(x, y, z, w)
		neighborCount := d.countNeighbors(x, y, z, w)

		switch {
		case active && (neighborCount == 2 || neighborCount == 3):
			work.set(x, y, z, w, true)
		case !active && neighborCount == 3:
			work.set(x, y, z, w, true)
		default:
			work.set(x, y, z, w, false)
		}
	}

	for x := d.minX - 1; x <= d.maxX+1; x++ {
		for y := d.minY - 1; y <= d.maxY+1; y++ {
			for z := d.minZ - 1; z <= d.maxZ+1; z++ {
				if d.Wenabled {
					for w := d.minW - 1; w <= d.maxW+1; w++ {
						stepCube(x, y, z, w)
					}
					continue
				}
				stepCube(x, y, z, 0)
			}
		}
	}

	d.cubes = work.cubes
	d.minX--
	d.minY--
	d.minZ--
	d.minW--
	d.maxX++
	d.maxY++
	d.maxZ++
	d.maxW++
}

func (d *dimension) countActive() (count int) {
	for _, active := range d.cubes {
		if active {
			count++
		}
	}
	return
}

func (d *dimension) print() {
	printCube := func(x, y, z, w int) {
		if d.isActive(x, y, z, w) {
			print("#")
		} else {
			print(".")
		}
	}
	for z := d.minZ; z <= d.maxZ; z++ {
		fmt.Printf("\nz = %d\n", z)
		for y := d.minY; y <= d.maxY; y++ {
			for x := d.minX; x <= d.maxX; x++ {
				if d.Wenabled {
					for w := d.minW - 1; w <= d.maxW+1; w++ {
						printCube(x, y, z, w)
					}
					continue
				}
				printCube(x, y, z, 0)
			}
			println()
		}
	}
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (space *dimension) {
	box := packr.New("day17", "./2020/day17")
	s, err := box.FindString("input")
	check(err)
	space = &dimension{cubes: map[position]bool{}}
	for y, line := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		for x, char := range line {
			space.set(x, y, 0, 0, char == '#')
			space.maxX = x
		}
		space.maxY = y
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
