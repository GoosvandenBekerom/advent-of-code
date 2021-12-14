package main

import (
	"bytes"
	_ "embed"
	"regexp"
	"strconv"

	"fmt"
	"log"
	"strings"
)

//go:embed input
var input string

func main() {
	part1()
	part2()
}

// ----------------------------------------
// solution
// ----------------------------------------

type vector struct{ x, y int }
type fold struct {
	axle uint8
	line int
}

func processInput() (map[vector]struct{}, []fold, int, int) {
	parts := strings.Split(normalizeNewlines(input), "\n\n")
	rawCoords := strings.Split(parts[0], "\n")
	positions := make(map[vector]struct{})
	var maxX, maxY int

	for _, raw := range rawCoords {
		coords := strings.Split(raw, ",")
		x := toInt(coords[0])
		y := toInt(coords[1])
		positions[vector{x: x, y: y}] = struct{}{}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	rawFolds := strings.Split(parts[1], "\n")
	pattern := regexp.MustCompile("(x|y)=(\\d+)")
	folds := make([]fold, len(rawFolds))

	for i, raw := range rawFolds {
		matches := pattern.FindStringSubmatch(raw)
		folds[i] = fold{
			axle: matches[1][0],
			line: toInt(matches[2]),
		}
	}
	return positions, folds, maxX, maxY
}

func printPositions(positions map[vector]struct{}, maxX, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, exists := positions[vector{x, y}]; exists {
				print("#")
				continue
			}
			print(".")
		}
		println()
	}
}

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	positions, folds, maxX, maxY := processInput()

	f := folds[0]

	var points map[vector]struct{}
	if f.axle == 'y' {
		points = copyPartial(positions, maxX+1, f.line)
	} else {
		points = copyPartial(positions, f.line, maxY+1)
	}

	if f.axle == 'y' {
		for x := 0; x <= maxX; x++ {
			for i := f.line + 1; i < f.line*2+1; i++ {
				if _, ok := positions[vector{x: x, y: i}]; !ok {
					continue
				}
				y := f.line - (i - f.line)
				points[vector{x: x, y: y}] = struct{}{}
			}
		}
	} else {
		for y := 0; y <= maxY; y++ {
			for i := f.line + 1; i < f.line*2+1; i++ {
				if _, ok := positions[vector{x: i, y: y}]; !ok {
					continue
				}
				x := f.line - (i - f.line)
				points[vector{x: x, y: y}] = struct{}{}
			}
		}
	}

	var sum int
	for pos := range points {
		if f.axle == 'y' {
			if pos.y < f.line {
				sum++
			}
		} else {
			if pos.x < f.line {
				sum++
			}
		}
	}

	println(sum)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	positions, folds, maxX, maxY := processInput()

	for _, f := range folds {
		var points map[vector]struct{}
		if f.axle == 'y' {
			points = copyPartial(positions, maxX+1, f.line)
		} else {
			points = copyPartial(positions, f.line, maxY+1)
		}

		if f.axle == 'y' {
			for x := 0; x <= maxX; x++ {
				for i := f.line + 1; i < f.line*2+1; i++ {
					if _, ok := positions[vector{x: x, y: i}]; !ok {
						continue
					}
					y := f.line - (i - f.line)
					points[vector{x: x, y: y}] = struct{}{}
				}
			}
		} else {
			for y := 0; y <= maxY; y++ {
				for i := f.line + 1; i < f.line*2+1; i++ {
					if _, ok := positions[vector{x: i, y: y}]; !ok {
						continue
					}
					x := f.line - (i - f.line)
					points[vector{x: x, y: y}] = struct{}{}
				}
			}
		}

		positions = points
	}

	// fmt.Printf("%#v\n", positions)
	printPositions(positions, 40, 6)
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

func normalizeNewlines(s string) string {
	d := []byte(s)
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return string(d)
}

func copyPartial(in map[vector]struct{}, xlimit, ylimit int) map[vector]struct{} {
	out := make(map[vector]struct{})
	for k, v := range in {
		if k.x < xlimit && k.y < ylimit {
			out[k] = v
		}
	}
	return out
}
