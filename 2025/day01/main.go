package main

import (
	_ "embed"
	"fmt"
	"strings"

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

	dialSize := 100

	pos := 50
	count := 0

	for _, line := range lines {
		step := utils.ToInt(line[1:]) % dialSize
		switch line[0] {
		case 'L':
			pos -= step
		case 'R':
			pos += step
		}
		if pos > 99 {
			pos = pos - dialSize
		}
		if pos < 0 {
			pos = dialSize + pos
		}
		if pos == 0 {
			count++
		}
		fmt.Println(line, step, pos, count)
	}

	return count
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	dialSize := 100

	pos := 50
	count := 0

	for _, line := range lines {
		startedAt := pos
		rawStep := utils.ToInt(line[1:])
		extra := rawStep / dialSize
		step := rawStep % dialSize
		switch line[0] {
		case 'L':
			pos -= step
		case 'R':
			pos += step
		}
		if pos > 99 {
			pos = pos - dialSize
			if pos != 0 {
				count++
			}
		}
		if pos < 0 {
			pos = dialSize + pos
			if pos != 0 && startedAt != 0 {
				count++
			}
		}
		if pos == 0 {
			count++
		}
		count += extra
		fmt.Println(line, step, pos, count, extra)
	}

	return count
}
