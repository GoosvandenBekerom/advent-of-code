package main

import (
	_ "embed"
	"fmt"
	"regexp"
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

var part1Regex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func part1(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	var total int

	for _, line := range lines {
		for _, match := range part1Regex.FindAllStringSubmatch(line, -1) {
			total += utils.ToInt(match[1]) * utils.ToInt(match[2])
		}
	}

	return total
}

var part2Regex = regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var total int
	enabled := true

	for _, line := range lines {
		for _, match := range part2Regex.FindAllStringSubmatch(line, -1) {
			switch match[0] {
			case "don't()":
				enabled = false
			case "do()":
				enabled = true
			default:
				if enabled {
					total += utils.ToInt(match[2]) * utils.ToInt(match[3])
				}
			}
		}
	}

	return total
}
