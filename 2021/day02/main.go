package main

import (
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

/*
forward X increases the horizontal position by X units.
down X increases the depth by X units.
up X decreases the depth by X units.
Note that since you're on a submarine, down and up affect your depth, and so they have the opposite result of what you might expect.

The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:

forward 5
down 5
forward 8
up 3
down 8
forward 2
Your horizontal position and depth both start at 0. The steps above would then modify them as follows:

forward 5 adds 5 to your horizontal position, a total of 5.
down 5 adds 5 to your depth, resulting in a value of 5.
forward 8 adds 8 to your horizontal position, a total of 13.
up 3 decreases your depth by 3, resulting in a value of 2.
down 8 adds 8 to your depth, resulting in a value of 10.
forward 2 adds 2 to your horizontal position, a total of 15.

After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)

Calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?
*/

// pattern matches a word followed by a space followed by a number group 1 = direction group 2 = velocity
var pattern = regexp.MustCompile("(\\w+) (\\d+)")

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var position int
	var depth int

	for _, line := range strings.Split(input, "\n") {
		matches := pattern.FindStringSubmatch(line)

		switch matches[1] {
		case "forward":
			position += toInt(matches[2])
			break
		case "up":
			depth -= toInt(matches[2])
			break
		case "down":
			depth += toInt(matches[2])
			break
		}
	}

	fmt.Printf("position: %d\n", position)
	fmt.Printf("depth: %d\n", depth)
	fmt.Printf("solution: %d\n", position*depth)
}

/*
In addition to horizontal position and depth, you'll also need to track a third value, aim, which also starts at 0.
The commands also mean something entirely different than you first thought:
	- down X increases your aim by X units.
	- up X decreases your aim by X units.
	- forward X does two things:
		- It increases your horizontal position by X units.
		- It increases your depth by your aim multiplied by X.
*/

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var position int
	var depth int
	var aim int

	for _, line := range strings.Split(input, "\n") {
		matches := pattern.FindStringSubmatch(line)

		switch matches[1] {
		case "forward":
			x := toInt(matches[2])
			position += x
			depth += aim * x
			break
		case "up":
			aim -= toInt(matches[2])
			break
		case "down":
			aim += toInt(matches[2])
			break
		}
	}

	fmt.Printf("position: %d\n", position)
	fmt.Printf("depth: %d\n", depth)
	fmt.Printf("solution: %d\n", position*depth)
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
