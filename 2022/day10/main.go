package main

import (
	_ "embed"
	"fmt"
	"github.com/GoosvandenBekerom/advent-of-code/2022/computer"
	"strings"
)

//go:embed input
var input string

func main() {
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	cpu := computer.CPU{X: 1}
	nextSignificant := 20
	var sum int

	for _, instruction := range computer.ParseInstructions(strings.Split(input, "\n")) {
		for i := 0; i < instruction.OP.CycleCost; i++ {
			cpu.Tick()
			if cpu.Cycle == nextSignificant {
				sum += cpu.CalculateSignalStrength()
				nextSignificant += 40
			}
		}
		instruction.Execute(&cpu)
	}

	return sum
}

func part2(input string) string {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	cpu := computer.CPU{X: 1}
	lineWidth := 40
	var drawing string

	for _, instruction := range computer.ParseInstructions(strings.Split(input, "\n")) {
		for i := 0; i < instruction.OP.CycleCost; i++ {
			cpu.Tick()
			// TODO: there is still a bug where sometimes the last position of a line
			//  gets a false positive. But not required to fix for this particular assignment.
			cursor := cpu.Cycle%lineWidth - 1
			if cursor == cpu.X-1 || cursor == cpu.X || cursor == cpu.X+1 {
				drawing += "#"
			} else {
				drawing += "."
			}

			if cpu.Cycle%lineWidth == 0 {
				drawing += "\n"
			}
		}
		instruction.Execute(&cpu)
	}

	return drawing
}
