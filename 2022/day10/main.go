package main

import (
	_ "embed"
	"strconv"

	"fmt"
	"log"
	"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type computer struct {
	cycle int
	x     int
}

func (cpu *computer) calculateSignalStrength() int {
	return cpu.cycle * cpu.x
}

func (cpu *computer) cycleOverlapsSprite(lineWidth int) bool {
	cursor := cpu.cycle%lineWidth - 1
	return cursor == cpu.x-1 || cursor == cpu.x || cursor == cpu.x+1
}

type operation struct {
	cycles  int
	execute func(string, *computer)
}

var operations = map[string]operation{
	"noop": {1, nil},
	"addx": {2, func(arg string, cpu *computer) {
		cpu.x += toInt(arg)
	}},
}

type instruction struct {
	operation operation
	arg       string
}

func (i instruction) execute(cpu *computer) {
	if i.operation.execute != nil {
		i.operation.execute(i.arg, cpu)
	}
}

func parseInstructions(lines string) []instruction {
	var instructions []instruction
	for _, line := range strings.Split(lines, "\n") {
		in, arg, _ := strings.Cut(line, " ")

		op, found := operations[in]
		if !found {
			panic("unknown instruction: " + in)
		}

		instructions = append(instructions, instruction{
			operation: op,
			arg:       arg,
		})
	}
	return instructions
}

func part1(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	cpu := computer{x: 1}
	nextSignificant := 20
	var sum int

	for _, instruction := range parseInstructions(input) {
		for i := 0; i < instruction.operation.cycles; i++ {
			cpu.cycle++
			if cpu.cycle == nextSignificant {
				sum += cpu.calculateSignalStrength()
				nextSignificant += 40
			}
		}

		instruction.execute(&cpu)
	}

	return sum
}

func part2(input string) string {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	cpu := computer{x: 1}
	lineWidth := 40
	var drawing string

	for _, instruction := range parseInstructions(input) {
		for i := 0; i < instruction.operation.cycles; i++ {
			cpu.cycle++
			if cpu.cycleOverlapsSprite(lineWidth) {
				drawing += "#"
			} else {
				drawing += "."
			}

			if cpu.cycle%lineWidth == 0 {
				drawing += "\n"
			}
		}

		instruction.execute(&cpu)
	}

	return drawing
}

// ----------------------------------------
// utils
// ----------------------------------------

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	check(err)
	return v
}
