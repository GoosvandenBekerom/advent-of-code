package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	println(part1())
	println(part2())
}

/*
rules:
If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
*/
func part1() int {
	instructions := input()
	compass := []direction{NORTH, EAST, SOUTH, WEST}
	facing := 1
	pos := vector{0, 0}
	for _, instruction := range instructions {
		printInstruction(instruction)
		var dir direction
		if instruction.direction == FORWARD {
			dir = compass[facing]
		} else {
			dir = instruction.direction
		}
		switch dir {
		case NORTH:
			pos.y += instruction.value
		case SOUTH:
			pos.y -= instruction.value
		case EAST:
			pos.x += instruction.value
		case WEST:
			pos.x -= instruction.value
		case LEFT:
			rotation := instruction.value / 90
			facing = (facing - rotation + 4) % 4
		case RIGHT:
			rotation := instruction.value / 90
			facing = (facing + rotation) % 4
		}
		fmt.Printf("position: %v\n", pos)
	}
	return int(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
}

func part2() int {
	instructions := input()
	pos := vector{0, 0}
	waypoint := vector{10, 1}
	for _, instruction := range instructions {
		printInstruction(instruction)
		switch instruction.direction {
		case NORTH:
			waypoint.y += instruction.value
		case SOUTH:
			waypoint.y -= instruction.value
		case EAST:
			waypoint.x += instruction.value
		case WEST:
			waypoint.x -= instruction.value
		case LEFT:
			rotation := instruction.value / 90
			for i := 0; i < rotation; i++ {
				waypoint = rotate(waypoint, false)
			}
		case RIGHT:
			rotation := instruction.value / 90
			for i := 0; i < rotation; i++ {
				waypoint = rotate(waypoint, true)
			}
		case FORWARD:
			for i := 0; i < instruction.value; i++ {
				pos.x += waypoint.x
				pos.y += waypoint.y
			}
		}
		fmt.Printf("position of waypoint: %v\n", waypoint)
		fmt.Printf("position of ship: %v\n", pos)
	}
	return int(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y)))
}

func rotate(point vector, clockwise bool) (result vector) {
	if clockwise {
		result.x = point.y
		result.y = -point.x
	} else {
		result.x = -point.y
		result.y = point.x
	}
	return
}

func printInstruction(instruction instruction) {
	var format string
	switch instruction.direction {
	case NORTH:
		format = "Move north by %d"
	case SOUTH:
		format = "Move south by %d"
	case EAST:
		format = "Move east by %d"
	case WEST:
		format = "Move west by %d"
	case LEFT:
		format = "Turn left by %d degrees"
	case RIGHT:
		format = "Turn right by %d degrees"
	case FORWARD:
		format = "Move forward by %d"
	}
	fmt.Printf(format+"\n", instruction.value)
}

// ----------------------------------------
// domain
// ----------------------------------------

type direction int

const (
	NORTH   direction = iota
	SOUTH   direction = iota
	EAST    direction = iota
	WEST    direction = iota
	LEFT    direction = iota
	RIGHT   direction = iota
	FORWARD direction = iota
)

type instruction struct {
	direction direction
	value     int
}

type vector struct {
	x, y int
}

// ----------------------------------------
// utils
// ----------------------------------------

func input() (instructions []instruction) {
	box := packr.New("day12", "./2020/day12")
	s, err := box.FindString("input")
	check(err)
	for _, raw := range strings.Split(strings.TrimSuffix(s, "\n"), "\n") {
		instructions = append(instructions, convert(raw))
	}
	return
}

func convert(raw string) instruction {
	var dir direction
	switch raw[0] {
	case 'N':
		dir = NORTH
	case 'S':
		dir = SOUTH
	case 'E':
		dir = EAST
	case 'W':
		dir = WEST
	case 'L':
		dir = LEFT
	case 'R':
		dir = RIGHT
	case 'F':
		dir = FORWARD
	default:
		panic("unknown character received for direction: " + string(raw[0]))
	}
	val, err := strconv.Atoi(raw[1:])
	check(err)
	return instruction{
		direction: dir,
		value:     val,
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
