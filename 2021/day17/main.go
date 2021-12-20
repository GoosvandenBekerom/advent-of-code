package main

import (
	"strconv"

	"fmt"
	"log"
)

func main() {
	part1()
	part2()
}

// ----------------------------------------
// solution
// ----------------------------------------

type vector struct{ x, y int }

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	// example: target area: x=20..30, y=-10..-5
	//min := vector{x: 20, y: -10}
	//max := vector{x: 30, y: -5}
	// input: 	target area: x=139..187, y=-148..-89
	min := vector{x: 139, y: -148}
	max := vector{x: 187, y: -89}

	solution := 0
	dy := min.y
	if dy < 0 {
		dy = -dy
	}
	dy *= 10

	for x := 1; x <= max.x; x++ {
		for y := -dy; y < dy; y++ {
			velocity := vector{x: x, y: y}

			maxY := 0
			p := vector{}
			for p.x+velocity.x <= max.x && p.y+velocity.y >= min.y {
				p.x += velocity.x
				p.y += velocity.y
				velocity.y -= 1
				if velocity.x < 0 {
					velocity.x += 1
				} else if velocity.x > 0 {
					velocity.x -= 1
				}
				if p.y > maxY {
					maxY = p.y
				}
			}
			if p.x >= min.x && p.x <= max.x && p.y >= min.y && p.y <= max.y {
				if maxY > solution {
					solution = maxY
				}
			}
		}
	}

	println(solution)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	// example: target area: x=20..30, y=-10..-5
	//min := vector{x: 20, y: -10}
	//max := vector{x: 30, y: -5}
	// input: 	target area: x=139..187, y=-148..-89
	min := vector{x: 139, y: -148}
	max := vector{x: 187, y: -89}

	unique := make(map[vector]bool)
	dy := min.y
	if dy < 0 {
		dy = -dy
	}
	dy *= 10

	for x := 1; x <= max.x; x++ {
		for y := -dy; y < dy; y++ {
			vel := vector{x: x, y: y}

			p := vector{}
			for p.x+vel.x <= max.x && p.y+vel.y >= min.y {
				p.x += vel.x
				p.y += vel.y
				vel.y -= 1
				if vel.x < 0 {
					vel.x += 1
				} else if vel.x > 0 {
					vel.x -= 1
				}
			}
			if p.x >= min.x && p.x <= max.x && p.y >= min.y && p.y <= max.y {
				unique[vector{x: x, y: y}] = true
			}
		}
	}

	println(len(unique))
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	check(err)
	return v
}
