package main

import (
	_ "embed"
	"fmt"
	"strings"
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
	for _, line := range lines {
		println(line)
	}

	return -1
}

func part2(lines []string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return -1
}
