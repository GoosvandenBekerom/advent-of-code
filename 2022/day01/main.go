package main

import (
	_ "embed"
	"sort"
	"strconv"

	"fmt"
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

	var currentElf, largestElf int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			if currentElf > largestElf {
				largestElf = currentElf
			}
			currentElf = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		currentElf += calories
	}

	if currentElf > largestElf {
		largestElf = currentElf
	}

	return largestElf
}

func part2(input string) int {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	var currentElf int
	var elfs []int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			elfs = append(elfs, currentElf)
			currentElf = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		currentElf += calories
	}

	sort.Ints(elfs)

	var sum int
	for i := len(elfs) - 1; i > len(elfs)-4; i-- {
		sum += elfs[i]
	}

	return sum
}
