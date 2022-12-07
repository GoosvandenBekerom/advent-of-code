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

// ----------------------------------------
// solution
// ----------------------------------------

func part1(input string) string {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 1:")
	for _, line := range strings.Split(input, "\n") {
		println(line)
	}

	return ""
}

func part2(input string) string {
	fmt.Println("\n___________________________________________")
	fmt.Println("part 2:")

	return ""
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
