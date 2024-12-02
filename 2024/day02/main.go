package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	gmath "github.com/GoosvandenBekerom/advent-of-code/math"
	"github.com/GoosvandenBekerom/advent-of-code/utils"
)

//go:embed input
var input string

func main() {
	lines := strings.Split(input, "\n")
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var safe int

	for _, line := range lines {
		if processLine(line, false) == 0 {
			safe++
		}
	}

	return safe
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var verdicts []int

	for _, line := range lines {
		verdicts = append(verdicts, processLine(line, false))
	}

	var safe int
	for i, unsafeSteps := range verdicts {
		if unsafeSteps <= 1 {
			safe++
		}
		if unsafeSteps == 4 {
			if processLine(lines[i], true) == 1 {
				safe++
			}
		}
	}

	return safe
}

func processLine(line string, flipDirection bool) int {
	var prev *int
	var direction bool
	var unsafeSteps int
	for i, raw := range strings.Split(line, " ") {
		v := utils.ToInt(raw)
		if prev == nil {
			prev = &v
			continue
		}
		if i == 1 {
			direction = math.Signbit(float64(*prev - v))
			if flipDirection {
				direction = !direction
			}
		}
		distance := gmath.Abs(*prev - v)
		if distance == 0 || distance > 3 || math.Signbit(float64(*prev-v)) != direction {
			unsafeSteps++
		}
		prev = &v
	}
	return unsafeSteps
}
