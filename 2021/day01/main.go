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
	part1()
	part2()
}

// ----------------------------------------
// solution
// ----------------------------------------

func part1() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")

	count := -1
	var previous int
	for _, line := range strings.Split(input, "\n") {
		current := toInt(line)
		if current > previous {
			count++
		}
		previous = current
	}
	fmt.Printf("increased %d times\n", count)
}

func part2() {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	lines := strings.Split(input, "\n")
	measurements := make([]int, len(lines))

	for i, line := range lines {
		measurements[i] = toInt(line)
	}

	count := -1
	var previous int
	for i := 1; i < len(measurements)-1; i++ {
		current := measurements[i-1] + measurements[i] + measurements[i+1]
		if current > previous {
			count++
		}
		previous = current
	}

	fmt.Printf("increased %d times\n", count)
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
